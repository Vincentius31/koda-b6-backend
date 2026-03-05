package di

import (
	"koda-b6-backend/internal/handlers"
	"koda-b6-backend/internal/repository"
	"koda-b6-backend/internal/service"

	"github.com/jackc/pgx/v5"
)

type Container struct {
	db *pgx.Conn

	//Users
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handlers.UserHandler

	//Roles
	roleRepo    *repository.RoleRepository
	roleService *service.RoleService
	roleHandler *handlers.RoleHandler
}

func NewContainer(db *pgx.Conn) *Container {
	container := &Container{
		db: db,
	}

	container.initDependencies()

	return container
}

func (c *Container) initDependencies() {
	//Users
	c.userRepo = repository.NewUserRepository(c.db)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)

	//Roles
	c.roleRepo = repository.NewRoleRepository(c.db)
	c.roleService = service.NewRoleService(c.roleRepo)
	c.roleHandler = handlers.NewRoleHandler(c.roleService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) RoleHandler() *handlers.RoleHandler {
	return c.roleHandler
}
