package middleware

import (
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{Success: false, Message: "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{Success: false, Message: "Invalid format"})
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("APP_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.WebResponse{Success: false, Message: "Sesi habis, silakan login ulang"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if val, exists := claims["user_id"]; exists {
				if floatVal, ok := val.(float64); ok {
					ctx.Set("user_id", int(floatVal))
				}
			}
			ctx.Set("email", claims["email"])
			ctx.Next()
		}
	}
}