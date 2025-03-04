package usecases

import (
	"library-mngmt/app/container"
	"library-mngmt/domain/repositories"
)

type AdminService struct {
	AdminRepository repositories.AdminRepositoriesInterface
}
type UserService struct {
	UserRepository repositories.UserRepositoriesInterface
}
type BookService struct {
	BookRepository repositories.BookRepositoriesInterface
}
type AuthService struct {
	AuthRepository repositories.AuthRepositoriesInterface
	UserRepository repositories.UserRepositoriesInterface
}

func NewUserService(ctr *container.Container) *UserService {
	return &UserService{
		UserRepository: ctr.Repositories.UserRepository,
	}
}
func NewBookService(ctr *container.Container) *BookService {
	return &BookService{
		BookRepository: ctr.Repositories.BookRepository,
	}
}
func NewAuthService(ctr *container.Container) *AuthService {
	return &AuthService{
		AuthRepository: ctr.Repositories.AuthRepository,
		UserRepository: ctr.Repositories.UserRepository,
	}
}
func NewAdminService(ctr *container.Container) *AdminService {
	return &AdminService{
		AdminRepository: ctr.Repositories.AdminRepository,
	}
}
