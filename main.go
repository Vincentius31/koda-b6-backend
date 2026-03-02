package main

import (
	"koda-b6-backend/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8888")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.Data(http.StatusOK, "", []byte(""))
		} else {
			ctx.Next()
		}

	}
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	godotenv.Load()

	r.Run("localhost:8888")
}
