package main

import (
	"context"
	"fmt"
	"koda-b6-backend/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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
	godotenv.Load()

	connConfig,err := pgx.ParseConfig("")

	if err != nil {
		fmt.Println("Failed to parse config")
		return
	}

	conn, err := pgx.Connect(context.Background(), connConfig.ConnString())
	if err != nil {
		fmt.Println("Connection Failed")
		return
	}
	defer conn.Close(context.Background())

	r := gin.Default()
	r.Use(corsMiddleware())

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
}
