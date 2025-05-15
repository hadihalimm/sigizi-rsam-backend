package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) RequireSession(c *gin.Context) {
	// session, err := config.SessionStore.Get(c.Request, "sigizi-rsam")
	// if err != nil || session.Values["userID"] == nil {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Session invalid or expired",
	// 	})
	// 	return
	// }

	// userID := session.Values["userID"].(uint)
	// c.Set("userID", userID)
	c.Next()
}
