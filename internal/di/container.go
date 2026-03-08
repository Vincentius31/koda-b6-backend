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

	//Category
	categoryRepo    *repository.CategoryRepository
	categoryService *service.CategoryService
	categoryHandler *handlers.CategoryHandler

	//Product
	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handlers.ProductHandler

	//Product_category
	productCategoryRepo    *repository.ProductCategoryRepository
	productCategoryService *service.ProductCategoryService
	productCategoryHandler *handlers.ProductCategoryHandler
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

	//Category
	c.categoryRepo = repository.NewCategoryRepository(c.db)
	c.categoryService = service.NewCategoryService(c.categoryRepo)
	c.categoryHandler = handlers.NewCategoryHandler(c.categoryService)

	//Product
	c.productRepo = repository.NewProductRepository(c.db)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)

	//Product_category
	c.productCategoryRepo = repository.NewProductCategoryRepository(c.db)
	c.productCategoryService = service.NewProductCategoryService(c.productCategoryRepo)
	c.productCategoryHandler = handlers.NewProductCategoryHandler(c.productCategoryService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}

func (c *Container) RoleHandler() *handlers.RoleHandler {
	return c.roleHandler
}

func (c *Container) CategoryHandler() *handlers.CategoryHandler {
	return c.categoryHandler
}

func (c *Container) ProductHandler() *handlers.ProductHandler {
	return c.productHandler
}

func (c *Container) ProductCategoryHandler() *handlers.ProductCategoryHandler {
	return c.productCategoryHandler
}
