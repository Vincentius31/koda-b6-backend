package di

import (
	"koda-b6-backend/internal/handlers"
	"koda-b6-backend/internal/repository"
	"koda-b6-backend/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	pool *pgxpool.Pool

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

	//transaction
	transactionRepo    *repository.TransactionRepository
	transactionService *service.TransactionService
	transactionHandler *handlers.TransactionHandler

	//transaction_product
	transactionProductRepo    *repository.TransactionProductRepository
	transactionProductService *service.TransactionProductService
	transactionProductHandler *handlers.TransactionProductHandler

	//review
	reviewRepo    *repository.ReviewRepository
	reviewService *service.ReviewService
	reviewHandler *handlers.ReviewHandler

	//forgot_password
	forgotPasswordRepo    *repository.ForgotPasswordRepository
	forgotPasswordService *service.ForgotPasswordService
	authHandler           *handlers.AuthHandler

	//landing_page
	landingService *service.LandingService
	landingHandler *handlers.LandingHandler

	//Product_page
	productPageService *service.ProductPageService
	productPageHandler *handlers.ProductPageHandler

	//Prdouct_detail
	productDetailService *service.DetailProductService
	productDetailHandler *handlers.DetailProductHandler
}

func NewContainer(pool *pgxpool.Pool) *Container {
	container := &Container{
		pool: pool,
	}

	container.initDependencies()

	return container
}

func (c *Container) initDependencies() {
	// Ganti semua c.db menjadi c.pool
	
	//Users
	c.userRepo = repository.NewUserRepository(c.pool)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)

	//Roles
	c.roleRepo = repository.NewRoleRepository(c.pool)
	c.roleService = service.NewRoleService(c.roleRepo)
	c.roleHandler = handlers.NewRoleHandler(c.roleService)

	//Category
	c.categoryRepo = repository.NewCategoryRepository(c.pool)
	c.categoryService = service.NewCategoryService(c.categoryRepo)
	c.categoryHandler = handlers.NewCategoryHandler(c.categoryService)

	//Product
	c.productRepo = repository.NewProductRepository(c.pool)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handlers.NewProductHandler(c.productService)

	//Product_category
	c.productCategoryRepo = repository.NewProductCategoryRepository(c.pool)
	c.productCategoryService = service.NewProductCategoryService(c.productCategoryRepo)
	c.productCategoryHandler = handlers.NewProductCategoryHandler(c.productCategoryService)

	//Product_image
	c.productImageRepo = repository.NewProductImageRepository(c.pool)
	c.productImageService = service.NewProductImageService(c.productImageRepo)
	c.productImageHandler = handlers.NewProductImageHandler(c.productImageService)

	//Product_variant
	c.productVariantRepo = repository.NewProductVariantRepository(c.pool)
	c.productVariantService = service.NewProductVariantService(c.productVariantRepo)
	c.productVariantHandler = handlers.NewProductVariantHandler(c.productVariantService)

	//Product_size
	c.productSizeRepo = repository.NewProductSizeRepository(c.pool)
	c.productSizeService = service.NewProductSizeService(c.productSizeRepo)
	c.productSizeHandler = handlers.NewProductSizeHandler(c.productSizeService)

	//discount
	c.discountRepo = repository.NewDiscountRepository(c.pool)
	c.discountService = service.NewDiscountService(c.discountRepo)
	c.discountHandler = handlers.NewDiscountHandler(c.discountService)

	//cart
	c.cartRepo = repository.NewCartRepository(c.pool)
	c.cartService = service.NewCartService(c.cartRepo)
	c.cartHandler = handlers.NewCartHandler(c.cartService)

	//transaction
	c.transactionRepo = repository.NewTransactionRepository(c.pool)
	c.transactionService = service.NewTransactionService(c.transactionRepo)
	c.transactionHandler = handlers.NewTransactionHandler(c.transactionService)

	//transaction_product
	c.transactionProductRepo = repository.NewTransactionProductRepository(c.pool)
	c.transactionProductService = service.NewTransactionProductService(c.transactionProductRepo)
	c.transactionProductHandler = handlers.NewTransactionProductHandler(c.transactionProductService)

	//review
	c.reviewRepo = repository.NewReviewRepository(c.pool)
	c.reviewService = service.NewReviewService(c.reviewRepo)
	c.reviewHandler = handlers.NewReviewHandler(c.reviewService)

	//forgot_password
	c.forgotPasswordRepo = repository.NewForgotPasswordRepository(c.pool)
	c.forgotPasswordService = service.NewForgotPasswordService(c.userRepo, c.forgotPasswordRepo)
	c.authHandler = handlers.NewAuthHandler(c.userService, c.forgotPasswordService)

	//landing_page
	c.landingService = service.NewLandingService(c.productRepo, c.reviewRepo)
	c.landingHandler = handlers.NewLandingHandler(c.landingService)

	//Products_page
	c.productPageService = service.NewProductPageService(c.productRepo, c.discountRepo)
	c.productPageHandler = handlers.NewProductPageHandler(c.productPageService)

	//Product_detail
	c.productDetailService = service.NewDetailProductService(c.productRepo)
	c.productDetailHandler = handlers.NewDetailProductHandler(c.productDetailService)
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

func (c *Container) TransactionHandler() *handlers.TransactionHandler {
	return c.transactionHandler
}

func (c *Container) TransactionProductHandler() *handlers.TransactionProductHandler {
	return c.transactionProductHandler
}

func (c *Container) ReviewHandler() *handlers.ReviewHandler {
	return c.reviewHandler
}

func (c *Container) AuthHandler() *handlers.AuthHandler {
	return c.authHandler
}

func (c *Container) LandingHandler() *handlers.LandingHandler {
	return c.landingHandler
}

func (c *Container) ProductPageHandler() *handlers.ProductPageHandler {
	return c.productPageHandler
}

func (c *Container) ProductDetailHandler() *handlers.DetailProductHandler {
	return c.productDetailHandler
}