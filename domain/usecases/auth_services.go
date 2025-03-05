package usecases

import (
	"errors"
	"library-mngmt/app/http/request"
	"library-mngmt/app/http/response"
	"library-mngmt/domain/entities"
	"library-mngmt/pkg"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthService) Login(req request.CreateLoginRequest, jwtSecret, encrytionKey string) (res response.LoginResponse, err error) {

	user, err := a.UserRepository.GetUserByEmail(req.Email)
	if user.ID == 0 || err != nil {
		return res, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return res, errors.New("invalid password")
	}
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.RoleID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return res, err
	}
	refreshTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return res, err
	}
	if err := a.UserRepository.UpdateRefreshKey(entities.UserToken{
		ID:           user.ID,
		RefreshToken: refreshTokenString,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
	}); err != nil {
		return res, err
	}
	apiKey, err := pkg.Decrypt(user.APIKey, encrytionKey)
	if err != nil {
		return res, err
	}
	return response.LoginResponse{
		UserID:       user.ID,
		APIKey:       apiKey,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (a *AuthService) RefreshToken(req request.RefreshTokenRequest, jwtSecret string) (res response.RefreshTokenResponse, err error) {
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return res, err
	}
	userID := uint(claims["user_id"].(float64))
	user, err := a.UserRepository.GetUserByID(userID)
	if user.ID == 0 || err != nil {
		return res, errors.New("user not found")
	}
	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		return res, errors.New("token expired")
	}
	accessTokenClaims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.RoleID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return res, err
	}
	return response.RefreshTokenResponse{
		AccessToken: accessTokenString,
	}, nil
}

func (a *AuthService) Logout(userID uint) error {

	if err := a.UserRepository.UpdateRefreshKey(entities.UserToken{
		ID:           userID,
		RefreshToken: "",
		ExpiresAt:    time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (a *AuthService) ValidateRolePermission(roleID uint, permission string) bool {
	permissions, err := a.AdminRepository.GetPermissionsRoleID(roleID)
	if err != nil {
		return false
	}
	for _, p := range permissions {
		if p.Action == permission {
			return true
		}
	}

	return false
}
