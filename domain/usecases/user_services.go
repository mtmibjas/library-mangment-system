package usecases

import (
	"library-mngmt/app/http/request"
	"library-mngmt/domain/entities"
	"library-mngmt/domain/globals"
	"library-mngmt/pkg"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(user request.CreateUserRequest, key string) error {

	text, err := globals.GenerateAPIKey()
	if err != nil {
		return err
	}
	apiKey, err := pkg.Encrypt(text, key)
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userModel := entities.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
		RoleID:   user.RoleID,
		APIKey:   apiKey,
	}
	if err := s.UserRepository.CreateUser(userModel); err != nil {

	}
	return nil
}

func (s *UserService) GetUser(id uint) (entities.User, error) {
	return s.UserRepository.GetUserByID(id)
}
func (s *UserService) GetUserByEmail(email string) (entities.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}
func (s *UserService) GetUserList() ([]entities.User, error) {
	return s.UserRepository.GetUserList()
}
func (s *UserService) UpdateUser(req request.UpdateUserRequest) error {
	user := entities.User{
		ID:     req.ID,
		Name:   req.Name,
		Email:  req.Email,
		RoleID: req.RoleID,
	}

	if err := s.UserRepository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}
func (s *UserService) UpdatePassword(req request.UpdatePasswordRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := entities.User{
		ID:       req.ID,
		Password: string(hashedPassword),
	}
	if err := s.UserRepository.UpdatePassword(user); err != nil {
		return err
	}
	return nil
}
func (s *UserService) DeleteUser(userID uint) error {
	if err := s.UserRepository.DeleteUser(userID); err != nil {
		return err
	}
	return nil
}
func (s *UserService) GetBorrowedHistoryByUserID(id uint) ([]entities.BorrowRecord, error) {
	return s.UserRepository.GetBorrowedHistoryByUserID(id)
}
