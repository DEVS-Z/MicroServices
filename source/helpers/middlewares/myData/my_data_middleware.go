package my_data_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func InjectUserIDAsParamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsAny, exists := c.Get("tokenData")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token claims found"})
			return
		}

		claims := claimsAny.(jwt.MapClaims)

		if claims == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No claims found in token"})
			return
		}

		userID := claims["id"]
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token claims"})
			return
		}

		idUser := int(userID.(float64))
		if idUser <= 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Simular un "param" agregándolo al contexto
		c.Set("param_id", idUser)

		c.Next()
	}
}
