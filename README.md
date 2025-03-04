# library-mangment-system

This is a backend api written in Go using the Echo framework that give needed api for library-mangment-system

## Prerequisites

- Go 1.21.6
- Git

## Getting Started
- Setup PostgresDB
- Do the migration

### Helpful Commands
- `make run`: Starts the web analyzer.
- `make test`: Executes the test suite.
- `make swag`: initate swagger document.
- `make admin-run`: insert seed
- `make db-mig-create`: create migration sql
- `make db-seed-create`: create seed sql
- `make db-mig cmd=up`: up migration sql
- `make db-mig cmd=down`: down migration sql
- `make db-seed cmd=up`: up seed sql
- `make db-seed cmd=down`: down seed sql

### Clone the Repository

```bash
git https://github.com/mtmibjas/library-mangment-system.git
cd library-mangment-system.
```

### Instruction to setup the project
1. Clone the project
2. Setup the PostgresDB
3. Add config details .env.local and config/*.yaml
4. Run `make db-mig cmd=up`
5. Run `make db-seed cmd=up`
6. Run `make run`