package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/sigizi-rsam/internal/config"
)

func (s *Server) RequireSession(c *gin.Context) {
	session, err := config.SessionStore.Get(c.Request, "sigizi-rsam")
	if err != nil || session.Values["userID"] == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Session invalid or expired",
		})
		return
	}

	userID := session.Values["userID"].(uint)
	c.Set("userID", userID)
	c.Set("role", session.Values["role"])
	c.Next()
}

func (s *Server) RequireAdminRole(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	c.Next()
}
