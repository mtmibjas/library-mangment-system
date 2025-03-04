package resolver

import (
	"fmt"
	"library-mngmt/app/config"
	"library-mngmt/app/container"
)

type Resolver struct {
	Config       *config.Config
	Adapters     container.Adapters
	Repositories container.Repositories
}

func NewAdapter(cfg *config.Config) *Resolver {
	return &Resolver{
		Config: cfg,
	}
}

func (r *Resolver) Resolve() *container.Container {
	r.resolveAdapters()
	r.resolveRepositories()

	return &container.Container{
		Adapters:     r.Adapters,
		Repositories: r.Repositories,
	}
}

func (r *Resolver) resolveAdapters() {
	db, err := resolveDBAdapter(r.Config.DB)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	r.Adapters = container.Adapters{
		Database: db,
	}
}
