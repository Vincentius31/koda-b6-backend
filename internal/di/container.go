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

	//Product_images
	productImageRepo    *repository.ProductImageRepository
	productImageService *service.ProductImageService
	productImageHandler *handlers.ProductImageHandler

	//Product_variant
	productVariantRepo    *repository.ProductVariantRepository
	productVariantService *service.ProductVariantService
	productVariantHandler *handlers.ProductVariantHandler

	//Product_size
	productSizeRepo    *repository.ProductSizeRepository
	productSizeService *service.ProductSizeService
	productSizeHandler *handlers.ProductSizeHandler

	//discount
	discountRepo    *repository.DiscountRepository
	discountService *service.DiscountService
	discountHandler *handlers.DiscountHandler

	//cart
	cartRepo    *repository.CartRepository
	cartService *service.CartService
	cartHandler *handlers.CartHandler
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

	//Product_image
	c.productImageRepo = repository.NewProductImageRepository(c.db)
	c.productImageService = service.NewProductImageService(c.productImageRepo)
	c.productImageHandler = handlers.NewProductImageHandler(c.productImageService)

	//Product_variant
	c.productVariantRepo = repository.NewProductVariantRepository(c.db)
	c.productVariantService = service.NewProductVariantService(c.productVariantRepo)
	c.productVariantHandler = handlers.NewProductVariantHandler(c.productVariantService)

	//Product_size
	c.productSizeRepo = repository.NewProductSizeRepository(c.db)
	c.productSizeService = service.NewProductSizeService(c.productSizeRepo)
	c.productSizeHandler = handlers.NewProductSizeHandler(c.productSizeService)

	//discount
	c.discountRepo = repository.NewDiscountRepository(c.db)
	c.discountService = service.NewDiscountService(c.discountRepo)
	c.discountHandler = handlers.NewDiscountHandler(c.discountService)

	//cart
	c.cartRepo = repository.NewCartRepository(c.db)
	c.cartService = service.NewCartService(c.cartRepo)
	c.cartHandler = handlers.NewCartHandler(c.cartService)
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

func (c *Container) ProductImageHandler() *handlers.ProductImageHandler {
	return c.productImageHandler
}

func (c *Container) ProductVariantHandler() *handlers.ProductVariantHandler {
	return c.productVariantHandler
}

func (c *Container) ProductSizeHandler() *handlers.ProductSizeHandler {
	return c.productSizeHandler
}

func (c *Container) DiscountHandler() *handlers.DiscountHandler {
	return c.discountHandler
}

func (c *Container) CartHandler() *handlers.CartHandler {
	return c.cartHandler
}
