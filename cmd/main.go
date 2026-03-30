package main

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/routes"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Origin")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}

func main() {
	godotenv.Load()

	r := gin.Default()
	r.Use(corsMiddleware())

	if _, err := os.Stat("uploads/users"); os.IsNotExist(err) {
		os.MkdirAll("uploads/users", 0755)
	}
	r.Static("/uploads", "./uploads")

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		fmt.Println("DATABASE_URL is not set in .env")
		return
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		os.Exit(1)
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Println("Database connection test failed")
		return
	}
	fmt.Println("Successfully connected to the database with connection pool!")

	routes.SetupRoutes(r, pool)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	r.Run(fmt.Sprintf(":%s", port))
}