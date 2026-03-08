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
}
