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
}