package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", s.authHandler.Register)
		auth.POST("/login", s.authHandler.Login)
		auth.POST("/logout", s.authHandler.Logout)
	}

	return r
}
