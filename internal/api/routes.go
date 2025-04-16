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

	roomType := r.Group("/room-type")
	{
		roomType.POST("", s.roomTypeHandler.Create)
		roomType.GET("", s.roomTypeHandler.GetAll)
		roomType.GET("/:id", s.roomTypeHandler.GetByID)
		roomType.PATCH("/:id", s.roomTypeHandler.Update)
		roomType.DELETE("/:id", s.roomTypeHandler.Delete)
	}

	room := r.Group("/room")
	{
		room.POST("", s.roomHandler.Create)
		room.GET("", s.roomHandler.GetAll)
		room.GET("/:id", s.roomHandler.GetByID)
		room.PATCH("/:id", s.roomHandler.Update)
		room.DELETE("/:id", s.roomHandler.Delete)
	}

	return r
}
