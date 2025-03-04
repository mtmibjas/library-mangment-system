package container

import (
	"database/sql"
	"library-mngmt/domain/repositories"
)

type Container struct {
	Adapters     Adapters
	Repositories Repositories
}

type Repositories struct {
	DataRepository  repositories.DataRepositoriesInterface
	UserRepository  repositories.UserRepositoriesInterface
	BookRepository  repositories.BookRepositoriesInterface
	AuthRepository  repositories.AuthRepositoriesInterface
	AdminRepository repositories.AdminRepositoriesInterface
}

type Adapters struct {
	Database *sql.DB
}
