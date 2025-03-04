package repositories

import (
	"database/sql"
	"library-mngmt/app/config"
	"library-mngmt/app/container"
)

type AdminRepository struct {
	Database *sql.DB
}
type UserRepository struct {
	Database *sql.DB
}
type BookRepository struct {
	Database *sql.DB
}
type AuthRepository struct {
	Database *sql.DB
}

func NewUserRepository(c *config.Config, a *container.Adapters) *UserRepository {
	return &UserRepository{
		Database: a.Database,
	}
}
func NewBookRepository(c *config.Config, a *container.Adapters) *BookRepository {
	return &BookRepository{
		Database: a.Database,
	}
}
func NewAuthRepository(c *config.Config, a *container.Adapters) *AuthRepository {
	return &AuthRepository{
		Database: a.Database,
	}
}
func NewAdminRepository(c *config.Config, a *container.Adapters) *AdminRepository {
	return &AdminRepository{
		Database: a.Database,
	}
}
