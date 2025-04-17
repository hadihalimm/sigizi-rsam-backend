package config

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte("tes123"))

func init() {
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int((12 * time.Hour).Seconds()),
		HttpOnly: true,
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := SessionStore.Get(c.Request, "sigizi-session")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to load session",
			})
			return
		}

		userID, ok := session.Values["userID"].(uint)
		if !ok || userID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid session or no user_id",
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
