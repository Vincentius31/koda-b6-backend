package routes

import (
	"koda-b6-backend/internal/di"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r *gin.Engine, conn *pgx.Conn) {
	container := di.NewContainer(conn)
	
	userHandler := container.UserHandler()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("", userHandler.GetAll)
		userRoutes.GET("/:id", userHandler.GetByID)
		userRoutes.POST("", userHandler.Create)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}
}