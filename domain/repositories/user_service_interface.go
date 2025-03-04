package repositories

import "library-mngmt/domain/entities"

type UserRepositoriesInterface interface {
	CreateUser(user entities.User) error
	GetUserByID(id uint) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	GetUserList() ([]entities.User, error)
	UpdateUser(user entities.User) error
	DeleteUser(id uint) error
	UpdatePassword(user entities.User) error
	GetBorrowedHistoryByUserID(id uint) ([]entities.BorrowRecord, error)
	UpdateRefreshKey(user entities.UserToken) error
}
