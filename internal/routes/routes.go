package routes

import (
	"koda-b6-backend/internal/di"
	"koda-b6-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r *gin.Engine, conn *pgx.Conn) {
	container := di.NewContainer(conn)

	userHandler := container.UserHandler()
	roleHandler := container.RoleHandler()
	categoryHandler := container.CategoryHandler()
	productHandler := container.ProductHandler()
	productCategoryHandler := container.ProductCategoryHandler()
	productImageHandler := container.ProductImageHandler()
	productVariantHandler := container.ProductVariantHandler()
	productSizeHandler := container.ProductSizeHandler()
	discountHandler := container.DiscountHandler()
	cartHandler := container.CartHandler()
	transactionHandler := container.TransactionHandler()
	transactionProductHandler := container.TransactionProductHandler()

	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Create)

	userRoutes := r.Group("/users")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("", userHandler.GetAll)
		userRoutes.GET("/:id", userHandler.GetByID)
		userRoutes.PATCH("/:id/upload", userHandler.UploadProfile)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}

	roleRoutes := r.Group("/roles")
	{
		roleRoutes.GET("", roleHandler.GetAll)
		roleRoutes.GET("/:id", roleHandler.GetByID)
		roleRoutes.POST("", roleHandler.Create)
		roleRoutes.PUT("/:id", roleHandler.Update)
		roleRoutes.DELETE("/:id", roleHandler.Delete)
	}

	categoryRoutes := r.Group("/categories")
	{
		categoryRoutes.GET("", categoryHandler.GetAll)
		categoryRoutes.GET("/:id", categoryHandler.GetByID)
		categoryRoutes.POST("", categoryHandler.Create)
		categoryRoutes.PUT("/:id", categoryHandler.Update)
		categoryRoutes.DELETE("/:id", categoryHandler.Delete)
	}

	productRoutes := r.Group("/product")
	{
		productRoutes.GET("", productHandler.GetAll)
		productRoutes.GET("/:id", productHandler.GetByID)
		productRoutes.POST("", productHandler.Create)
		productRoutes.PUT("/:id", productHandler.Update)
		productRoutes.DELETE("/:id", productHandler.Delete)
	}

	productCategoryRoutes := r.Group("/productcategory")
	{
		productCategoryRoutes.GET("", productCategoryHandler.GetAll)
		productCategoryRoutes.GET("/:id", productCategoryHandler.GetByID)
		productCategoryRoutes.POST("", productCategoryHandler.Create)
		productCategoryRoutes.PUT("/:id", productCategoryHandler.Update)
		productCategoryRoutes.DELETE("/:id", productCategoryHandler.Delete)
	}

	productImageRoutes := r.Group("/productimage")
	{
		productImageRoutes.GET("", productImageHandler.GetAll)
		productImageRoutes.GET("/:id", productImageHandler.GetByID)
		productImageRoutes.POST("", productImageHandler.Create)
		productImageRoutes.PUT("/:id", productImageHandler.Update)
		productImageRoutes.DELETE("/:id", productImageHandler.Delete)
	}

	productVariantRoutes := r.Group("/productvariant")
	{
		productVariantRoutes.GET("", productVariantHandler.GetAll)
		productVariantRoutes.GET("/:id", productVariantHandler.GetByID)
		productVariantRoutes.POST("", productVariantHandler.Create)
		productVariantRoutes.PUT("/:id", productVariantHandler.Update)
		productVariantRoutes.DELETE("/:id", productVariantHandler.Delete)
	}

	productSizeRoutes := r.Group("/productsize")
	{
		productSizeRoutes.GET("", productSizeHandler.GetAll)
		productSizeRoutes.GET("/:id", productSizeHandler.GetByID)
		productSizeRoutes.POST("", productSizeHandler.Create)
		productSizeRoutes.PUT("/:id", productSizeHandler.Update)
		productSizeRoutes.DELETE("/:id", productSizeHandler.Delete)
	}

	discountRoutes := r.Group("/discount")
	{
		discountRoutes.GET("", discountHandler.GetAll)
		discountRoutes.GET("/:id", discountHandler.GetByID)
		discountRoutes.POST("", discountHandler.Create)
		discountRoutes.PUT("/:id", discountHandler.Update)
		discountRoutes.DELETE("/:id", discountHandler.Delete)
	}

	cartRoutes := r.Group("/cart")
	{
		cartRoutes.GET("", cartHandler.GetAll)
		cartRoutes.GET("/:id", cartHandler.GetByID)
		cartRoutes.POST("", cartHandler.Create)
		cartRoutes.PUT("/:id", cartHandler.Update)
		cartRoutes.DELETE("/:id", cartHandler.Delete)
	}

	transactionRoutes := r.Group("/transaction")
	{
		transactionRoutes.GET("", transactionHandler.GetAll)
		transactionRoutes.GET("/:id", transactionHandler.GetByID)
		transactionRoutes.POST("", transactionHandler.Create)
		transactionRoutes.PUT("/:id", transactionHandler.Update)
		transactionRoutes.DELETE("/:id", transactionHandler.Delete)
	}

	transactionProductRoutes := r.Group("/transactionproduct")
	{
		transactionProductRoutes.GET("", transactionProductHandler.GetAll)
		transactionProductRoutes.GET("/:id", transactionProductHandler.GetByID)
		transactionProductRoutes.POST("", transactionProductHandler.Create)
		transactionProductRoutes.PUT("/:id", transactionProductHandler.Update)
		transactionProductRoutes.DELETE("/:id", transactionProductHandler.Delete)
	}
}
