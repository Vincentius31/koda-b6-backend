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
	reviewHandler := container.ReviewHandler()
	authHandler := container.AuthHandler()
	landingHandler := container.LandingHandler()
	productPageHandler := container.ProductPageHandler()
	productDetailHandler := container.ProductDetailHandler()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/forgot-password", authHandler.RequestForgotPassword)
		authRoutes.PATCH("/forgot-password", authHandler.ResetPassword)
	}

	landingRoutes := r.Group("/landing")
	{
		landingRoutes.GET("/recommended-products", landingHandler.GetRecommendedProducts)
		landingRoutes.GET("/reviews", landingHandler.GetLatestReviews)
	}

	productPageRoutes := r.Group("/products")
	{
		productPageRoutes.GET("", productPageHandler.GetCatalog)
		productPageRoutes.GET("/promos", productPageHandler.GetPromos)
	}

	productDetailRoutes := r.Group("/detail-product")
	{
		productDetailRoutes.GET("/:id", productDetailHandler.GetDetail)
	}
	userRoutes := r.Group("/")
	userRoutes.Use(middleware.AuthMiddleware())
	{
		userRoutes.GET("/profile", userHandler.GetProfile)
		userRoutes.PATCH("/profile", userHandler.UpdateProfile)
		userRoutes.POST("/profile/upload", userHandler.UploadProfile)
	}

	adminRoutes := r.Group("/admin")
	{
		userRoutes := adminRoutes.Group("/users")
		userRoutes.Use(middleware.AuthMiddleware())
		{
			userRoutes.GET("", userHandler.GetAll)
			userRoutes.GET("/:id", userHandler.GetByID)
			userRoutes.PATCH("/:id/upload", userHandler.UploadProfile)
			userRoutes.PATCH("/:id", userHandler.Update)
			userRoutes.DELETE("/:id", userHandler.Delete)
		}
	
		roleRoutes := adminRoutes.Group("/roles")
		{
			roleRoutes.GET("", roleHandler.GetAll)
			roleRoutes.GET("/:id", roleHandler.GetByID)
			roleRoutes.POST("", roleHandler.Create)
			roleRoutes.PATCH("/:id", roleHandler.Update)
			roleRoutes.DELETE("/:id", roleHandler.Delete)
		}
	
		categoryRoutes := adminRoutes.Group("/categories")
		{
			categoryRoutes.GET("", categoryHandler.GetAll)
			categoryRoutes.GET("/:id", categoryHandler.GetByID)
			categoryRoutes.POST("", categoryHandler.Create)
			categoryRoutes.PATCH("/:id", categoryHandler.Update)
			categoryRoutes.DELETE("/:id", categoryHandler.Delete)
		}
	
		productRoutes := adminRoutes.Group("/product")
		{
			productRoutes.GET("", productHandler.GetAll)
			productRoutes.GET("/:id", productHandler.GetByID)
			productRoutes.POST("", productHandler.Create)
			productRoutes.PATCH("/:id", productHandler.Update)
			productRoutes.DELETE("/:id", productHandler.Delete)
		}
	
		productCategoryRoutes := adminRoutes.Group("/productcategory")
		{
			productCategoryRoutes.GET("", productCategoryHandler.GetAll)
			productCategoryRoutes.GET("/:id", productCategoryHandler.GetByID)
			productCategoryRoutes.POST("", productCategoryHandler.Create)
			productCategoryRoutes.PATCH("/:id", productCategoryHandler.Update)
			productCategoryRoutes.DELETE("/:id", productCategoryHandler.Delete)
		}
	
		productImageRoutes := adminRoutes.Group("/productimage")
		{
			productImageRoutes.GET("", productImageHandler.GetAll)
			productImageRoutes.GET("/:id", productImageHandler.GetByID)
			productImageRoutes.POST("", productImageHandler.Create)
			productImageRoutes.PATCH("/:id", productImageHandler.Update)
			productImageRoutes.DELETE("/:id", productImageHandler.Delete)
		}
	
		productVariantRoutes := adminRoutes.Group("/productvariant")
		{
			productVariantRoutes.GET("", productVariantHandler.GetAll)
			productVariantRoutes.GET("/:id", productVariantHandler.GetByID)
			productVariantRoutes.POST("", productVariantHandler.Create)
			productVariantRoutes.PATCH("/:id", productVariantHandler.Update)
			productVariantRoutes.DELETE("/:id", productVariantHandler.Delete)
		}
	
		productSizeRoutes := adminRoutes.Group("/productsize")
		{
			productSizeRoutes.GET("", productSizeHandler.GetAll)
			productSizeRoutes.GET("/:id", productSizeHandler.GetByID)
			productSizeRoutes.POST("", productSizeHandler.Create)
			productSizeRoutes.PATCH("/:id", productSizeHandler.Update)
			productSizeRoutes.DELETE("/:id", productSizeHandler.Delete)
		}
	
		discountRoutes := adminRoutes.Group("/discount")
		{
			discountRoutes.GET("", discountHandler.GetAll)
			discountRoutes.GET("/:id", discountHandler.GetByID)
			discountRoutes.POST("", discountHandler.Create)
			discountRoutes.PATCH("/:id", discountHandler.Update)
			discountRoutes.DELETE("/:id", discountHandler.Delete)
		}
	
		cartRoutes := adminRoutes.Group("/cart")
		{
			cartRoutes.GET("", cartHandler.GetAll)
			cartRoutes.GET("/:id", cartHandler.GetByID)
			cartRoutes.POST("", cartHandler.Create)
			cartRoutes.PATCH("/:id", cartHandler.Update)
			cartRoutes.DELETE("/:id", cartHandler.Delete)
		}
	
		transactionRoutes := adminRoutes.Group("/transaction")
		{
			transactionRoutes.GET("", transactionHandler.GetAll)
			transactionRoutes.GET("/:id", transactionHandler.GetByID)
			transactionRoutes.POST("", transactionHandler.Create)
			transactionRoutes.PATCH("/:id", transactionHandler.Update)
			transactionRoutes.DELETE("/:id", transactionHandler.Delete)
		}
	
		transactionProductRoutes := adminRoutes.Group("/transactionproduct")
		{
			transactionProductRoutes.GET("", transactionProductHandler.GetAll)
			transactionProductRoutes.GET("/:id", transactionProductHandler.GetByID)
			transactionProductRoutes.POST("", transactionProductHandler.Create)
			transactionProductRoutes.PATCH("/:id", transactionProductHandler.Update)
			transactionProductRoutes.DELETE("/:id", transactionProductHandler.Delete)
		}
	
		reviewRoutes := adminRoutes.Group("/review")
		{
			reviewRoutes.GET("", reviewHandler.GetAll)
			reviewRoutes.GET("/:id", reviewHandler.GetByID)
			reviewRoutes.POST("", reviewHandler.Create)
			reviewRoutes.PATCH("/:id", reviewHandler.Update)
			reviewRoutes.DELETE("/:id", reviewHandler.Delete)
		}
	}

}
