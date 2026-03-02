package routes

import (
	"koda-b6-backend/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

// SetupRoutes menerima parameter conn agar bisa digunakan oleh handler
func SetupRoutes(r *gin.Engine, conn *pgx.Conn) {
	
	// --- User Routes ---
	
	r.POST("/register", func(ctx *gin.Context) {
		handlers.RegisterHandler(ctx, conn)
	})

	r.POST("/login", func(ctx *gin.Context) {
		handlers.LoginHandler(ctx, conn)
	})

	r.GET("/users", func(ctx *gin.Context) {
		handlers.GetAllUsersHandler(ctx, conn)
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		handlers.GetUserByIdHandler(ctx, conn)
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		handlers.UpdateUserHandler(ctx, conn)
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		handlers.DeleteUserHandler(ctx, conn)
	})
}