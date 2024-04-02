package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"example/gingonic/utils" // Adjust import path as necessary

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies the token and extracts user info
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing."})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization header format must be 'Bearer {token}'."})
			return
		}

		token := parts[1]
		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token.", "details": err.Error()})
			return
		}

		// Log claims for debugging
		fmt.Printf("Claims: %#v\n", claims)

		// Attach user information to context
		c.Set("email", claims["email"])
		c.Set("userId", claims["userId"])

		c.Next()
	}
}
