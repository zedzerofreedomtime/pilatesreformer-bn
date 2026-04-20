package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/service"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

func Authenticate(auth *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := strings.TrimSpace(c.GetHeader("Authorization"))
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing bearer token"})
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer"))
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing bearer token"})
			return
		}

		user, err := auth.SessionUser(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid session"})
			return
		}

		c.Set("authUser", user)
		c.Set("sessionToken", token)
		c.Next()
	}
}

func RequireRole(roleID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawUser, ok := c.Get("authUser")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing auth context"})
			return
		}

		user, ok := rawUser.(store.AuthenticatedUser)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid auth context"})
			return
		}

		if user.RoleID != roleID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			return
		}

		c.Next()
	}
}
