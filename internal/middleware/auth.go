package middleware

import (
	"control-financiero/internal/auth"
	"control-financiero/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := auth.ValidateJWT(tokenString, cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		// Establecer información del usuario en el contexto
		c.Set("userId", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("userRol", claims.Rol)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rol, exists := c.Get("userRol")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "No autorizado"})
			c.Abort()
			return
		}

		if rol != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: se requieren permisos de administrador"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) (primitive.ObjectID, error) {
	userID, exists := c.Get("userId")
	if !exists {
		return primitive.NilObjectID, nil
	}
	return userID.(primitive.ObjectID), nil
}
