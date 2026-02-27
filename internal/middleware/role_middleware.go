package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware allows the request only if AuthMiddleware has set user_role to required.
func RoleMiddleware(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleI, exists := c.Get("user_role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found"})
			return
		}
		role := roleI.(string)
		if role != required {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
