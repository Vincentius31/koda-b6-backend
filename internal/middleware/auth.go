package middleware

import (
	"fmt"
	"koda-b6-backend/internal/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Authorization header is required",
			})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Invalid authorization format. Please use 'Bearer <token>'",
			})
			ctx.Abort()
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("APP_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Invalid or expired token",
			})
			ctx.Abort()
			return 
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			ctx.Set("user_id", claims["user_id"])
			ctx.Set("email", claims["email"])
		}
		ctx.Next()
	}
}
