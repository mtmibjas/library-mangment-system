package resolver

import (
	"library-mngmt/app/container"
	"library-mngmt/repositories"
)

func (r *Resolver) resolveRepositories() {
	r.Repositories = container.Repositories{
		UserRepository:  repositories.NewUserRepository(r.Config, &r.Adapters),
		BookRepository:  repositories.NewBookRepository(r.Config, &r.Adapters),
		AuthRepository:  repositories.NewAuthRepository(r.Config, &r.Adapters),
		AdminRepository: repositories.NewAdminRepository(r.Config, &r.Adapters),
	}
}
