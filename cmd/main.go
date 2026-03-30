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
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}

	r := gin.Default()
	r.Use(corsMiddleware())

	if _, err := os.Stat("uploads/users"); os.IsNotExist(err) {
		os.MkdirAll("uploads/users", 0755)
	}
	r.Static("/uploads", "./uploads")

	config, err := pgxpool.ParseConfig("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to process database configuration: %v\n", err)
		os.Exit(1)
	}

	config.MaxConns = 20                    
	config.MinConns = 5                       
	config.MaxConnIdleTime = 30 * time.Minute 

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create database connection pool: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "The database is not responding: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to the database")

	routes.SetupRoutes(r, pool)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	fmt.Printf("The server runs on the port %s...\n", port)
	r.Run(fmt.Sprintf(":%s", port))
}