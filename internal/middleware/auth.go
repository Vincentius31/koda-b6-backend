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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Authorization header is required",
				Data:    nil,
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Invalid authorization format. Please use 'Bearer <token>'",
				Data:    nil,
			})
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("APP_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})

		// Jika token expired atau signature salah
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Invalid or expired token",
				Data:    nil,
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if userIDFloat, exists := claims["user_id"].(float64); exists {
				ctx.Set("user_id", int(userIDFloat))
			}
			
			ctx.Set("email", claims["email"])
			ctx.Next() 
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{
				Success: false,
				Message: "Unauthorized. Please login.",
				Data:    nil,
			})
		}
	}
}