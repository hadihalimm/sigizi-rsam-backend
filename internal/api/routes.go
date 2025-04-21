package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/register", s.authHandler.Register)
		auth.POST("/sign-in", s.authHandler.SignIn)
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

	food := r.Group("/food")
	{
		food.POST("", s.foodHandler.Create)
		food.GET("", s.foodHandler.GetAll)
		food.GET("/:id", s.foodHandler.GetByID)
		food.PATCH("/:id", s.foodHandler.Update)
		food.DELETE("/:id", s.foodHandler.Delete)
	}

	mealType := r.Group("/meal-type")
	{
		mealType.POST("", s.mealTypeHandler.Create)
		mealType.GET("", s.mealTypeHandler.GetAll)
		mealType.GET("/:id", s.mealTypeHandler.GetByID)
		mealType.PATCH("/:id", s.mealTypeHandler.Update)
		mealType.DELETE("/:id", s.mealTypeHandler.Delete)
	}

	mealItem := r.Group("/meal-item")
	{
		mealItem.POST("", s.mealItemHandler.Create)
		mealItem.GET("", s.mealItemHandler.GetAll)
		mealItem.GET("/:id", s.mealItemHandler.GetByID)
		mealItem.PATCH("/:id", s.mealItemHandler.Update)
		mealItem.DELETE("/:id", s.mealItemHandler.Delete)
	}

	patient := r.Group("/patient")
	{
		patient.POST("", s.patientHandler.Create)
		patient.GET("", s.patientHandler.GetAll)
		patient.GET("/:id", s.patientHandler.GetByID)
		patient.PATCH("/:id", s.patientHandler.Update)
		patient.DELETE("/:id", s.patientHandler.Delete)
	}

	dailyPatientMeal := r.Group("/daily-patient-meal")
	{
		dailyPatientMeal.POST("", s.dailyPatientMealHandler.Create)
		dailyPatientMeal.GET("", s.dailyPatientMealHandler.GetAll)
		dailyPatientMeal.GET("/:id", s.dailyPatientMealHandler.GetByID)
		dailyPatientMeal.PATCH("/:id", s.dailyPatientMealHandler.Update)
		dailyPatientMeal.DELETE("/:id", s.dailyPatientMealHandler.Delete)
		dailyPatientMeal.GET("/filter", s.dailyPatientMealHandler.FilterByDateAndRoomType)
	}

	return r
}
