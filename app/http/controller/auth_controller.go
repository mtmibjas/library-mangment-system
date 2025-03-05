package controller

import (
	"errors"
	"library-mngmt/app/config"
	"library-mngmt/app/container"
	"library-mngmt/app/http/request"
	"library-mngmt/app/http/response"
	"library-mngmt/domain/usecases"
	"library-mngmt/pkg"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Config       *config.Config
	Adapters     *container.Container
	AuthService  *usecases.AuthService
	UserServices *usecases.UserService
}

func NewAuthController(cfg *config.Config, ctr *container.Container) *AuthController {
	return &AuthController{
		Config:       cfg,
		Adapters:     ctr,
		AuthService:  usecases.NewAuthService(ctr),
		UserServices: usecases.NewUserService(ctr),
	}
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body request.CreateLoginRequest true "Login"
// @Success 201 {object} response.LoginResponse "Successful login"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/login [post]
func (ac *AuthController) Login(c echo.Context) error {
	req := request.CreateLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	res, err := ac.AuthService.Login(req, ac.Config.Service.JWTSecret, ac.Config.Service.EncrytionKey)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": res,
	})
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Security BearerAuth
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "Logout successful"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/logout [post]
func (ac *AuthController) Logout(c echo.Context) error {
	userID := c.Get("user_id").(float64)
	if err := ac.AuthService.Logout(uint(userID)); err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"message": "Logout successful",
	})
}

// RefreshToken godoc
// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Security BearerAuth
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param body body request.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} response.RefreshTokenResponse "Successful refresh token"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/refresh-token [post]
func (ac *AuthController) RefreshToken(c echo.Context) error {
	req := request.RefreshTokenRequest{}
	if err := c.Bind(&req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return response.Error(c, http.StatusBadRequest, err)
	}
	res, err := ac.AuthService.RefreshToken(req, ac.Config.Service.JWTSecret)
	if err != nil {
		return response.Error(c, http.StatusInternalServerError, err)
	}
	return response.Send(c, http.StatusOK, map[string]any{
		"data": res,
	})
}

func (ac *AuthController) ValidateAPIKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-KEY")
		if apiKey == "" {
			return response.Error(c, http.StatusUnauthorized, errors.New("missing API Key"))
		}
		useID := c.Get("user_id").(float64)
		user, err := ac.UserServices.GetUser(uint(useID))
		if err != nil {
			return response.Error(c, http.StatusInternalServerError, err)
		}
		key, err := pkg.Decrypt(user.APIKey, ac.Config.Service.EncrytionKey)
		if err != nil {
			return response.Error(c, http.StatusInternalServerError, err)
		}
		if key != apiKey {
			return response.Error(c, http.StatusUnauthorized, errors.New("invalid API Key"))
		}

		return next(c)
	}
}

func (ac *AuthController) ValidateRolePermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Get("role_id").(float64)
			if !ac.AuthService.ValidateRolePermission(uint(role), permission) {
				return response.Error(c, http.StatusForbidden, errors.New("forbidden"))
			}
			return next(c)
		}
	}
}
