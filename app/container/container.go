package container

import (
	"database/sql"
	"library-mngmt/domain/repositories"

	"github.com/redis/go-redis/v9"
)

type Container struct {
	Adapters     Adapters
	Repositories Repositories
}

type Repositories struct {
	UserRepository  repositories.UserRepositoriesInterface
	BookRepository  repositories.BookRepositoriesInterface
	AuthRepository  repositories.AuthRepositoriesInterface
	AdminRepository repositories.AdminRepositoriesInterface
}

type Adapters struct {
	Database *sql.DB
	Redis    *redis.Client
}
