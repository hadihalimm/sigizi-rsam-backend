package api

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_ORIGIN")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/register", s.authHandler.Register)
		auth.POST("/sign-in", s.authHandler.SignIn)
		auth.POST("/logout", s.authHandler.Logout)
		auth.GET("/check-session", s.authHandler.CheckSession)
	}

	user := r.Group("/user")
	user.Use(s.RequireSession)
	{
		user.GET("", s.userHandler.GetAll)
		user.GET("/:id", s.userHandler.GetByID)
		user.PATCH("/:id", s.userHandler.Update)
		user.DELETE("/:id", s.userHandler.Delete)
		user.POST("/:id/actions/reset-password", s.userHandler.ResetPassword)
		user.POST("/:id/actions/change-password", s.userHandler.UpdatePassword)
		user.POST("/:id/actions/change-name", s.userHandler.UpdateName)
	}

	roomType := r.Group("/room-type")
	roomType.Use(s.RequireSession)
	{
		roomType.POST("", s.roomTypeHandler.Create)
		roomType.GET("", s.roomTypeHandler.GetAll)
		roomType.GET("/:id", s.roomTypeHandler.GetByID)
		roomType.PATCH("/:id", s.roomTypeHandler.Update)
		roomType.DELETE("/:id", s.roomTypeHandler.Delete)
	}

	room := r.Group("/room")
	room.Use(s.RequireSession)
	{
		room.POST("", s.roomHandler.Create)
		room.GET("", s.roomHandler.GetAll)
		room.GET("/:id", s.roomHandler.GetByID)
		room.PATCH("/:id", s.roomHandler.Update)
		room.DELETE("/:id", s.roomHandler.Delete)
		room.GET("/filter", s.roomHandler.FilterByRoomType)
	}

	food := r.Group("/food")
	food.Use(s.RequireSession)
	{
		food.POST("", s.foodHandler.Create)
		food.GET("", s.foodHandler.GetAll)
		food.GET("/:id", s.foodHandler.GetByID)
		food.PATCH("/:id", s.foodHandler.Update)
		food.DELETE("/:id", s.foodHandler.Delete)
	}

	mealType := r.Group("/meal-type")
	mealType.Use(s.RequireSession)
	{
		mealType.POST("", s.mealTypeHandler.Create)
		mealType.GET("", s.mealTypeHandler.GetAll)
		mealType.GET("/:id", s.mealTypeHandler.GetByID)
		mealType.PATCH("/:id", s.mealTypeHandler.Update)
		mealType.DELETE("/:id", s.mealTypeHandler.Delete)
	}

	mealItem := r.Group("/meal-item")
	mealItem.Use(s.RequireSession)
	{
		mealItem.POST("", s.mealItemHandler.Create)
		mealItem.GET("", s.mealItemHandler.GetAll)
		mealItem.GET("/:id", s.mealItemHandler.GetByID)
		mealItem.PATCH("/:id", s.mealItemHandler.Update)
		mealItem.DELETE("/:id", s.mealItemHandler.Delete)
	}

	patient := r.Group("/patient")
	patient.Use(s.RequireSession)
	{
		patient.POST("", s.patientHandler.Create)
		patient.GET("", s.patientHandler.GetAll)
		patient.GET("/:id", s.patientHandler.GetByID)
		patient.PATCH("/:id", s.patientHandler.Update)
		patient.DELETE("/:id", s.patientHandler.Delete)
		patient.GET("/filter", s.patientHandler.FilterByRMN)
		patient.GET("/paginated", s.patientHandler.GetAllWithPaginationAndKeyword)
	}

	dailyPatientMeal := r.Group("/daily-patient-meal")
	dailyPatientMeal.Use(s.RequireSession)
	{
		dailyPatientMeal.POST("", s.dailyPatientMealHandler.Create)
		dailyPatientMeal.GET("", s.dailyPatientMealHandler.GetAll)
		dailyPatientMeal.GET("/:id", s.dailyPatientMealHandler.GetByID)
		dailyPatientMeal.PATCH("/:id", s.dailyPatientMealHandler.Update)
		dailyPatientMeal.DELETE("/:id", s.dailyPatientMealHandler.Delete)
		dailyPatientMeal.GET("/filter", s.dailyPatientMealHandler.FilterByDateAndRoomType)
		dailyPatientMeal.GET("/count", s.dailyPatientMealHandler.CountByDateAndRoomType)
		dailyPatientMeal.GET("/count/diet", s.dailyPatientMealHandler.CountDietCombinationsByDate)
		dailyPatientMeal.GET("/export", s.dailyPatientMealHandler.ExportToExcel)
		dailyPatientMeal.GET("/logs", s.dailyPatientMealHandler.FilterLogsByDate)
	}

	diet := r.Group("/diet")
	diet.Use(s.RequireSession)
	{
		diet.POST("", s.dietHandler.Create)
		diet.GET("", s.dietHandler.GetAll)
		diet.GET("/:id", s.dietHandler.GetByID)
		diet.PATCH("/:id", s.dietHandler.Update)
		diet.DELETE("/:id", s.dietHandler.Delete)
	}

	allergy := r.Group("/allergy")
	allergy.Use(s.RequireSession)
	{
		allergy.POST("", s.allergyHandler.Create)
		allergy.GET("", s.allergyHandler.GetAll)
		allergy.GET("/:id", s.allergyHandler.GetByID)
		allergy.PATCH("/:id", s.allergyHandler.Update)
		allergy.DELETE("/:id", s.allergyHandler.Delete)
	}

	return r
}
