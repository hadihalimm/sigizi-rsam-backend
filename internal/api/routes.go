package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", s.authHandler.Register)
		auth.POST("/sign-in", s.authHandler.SignIn)
		auth.POST("/logout", s.authHandler.Logout)
		auth.GET("/check-session", s.authHandler.CheckSession)
	}

	user := api.Group("/user")
	user.Use(s.RequireSession)
	{
		user.POST("/:id/actions/reset-password", s.userHandler.ResetPassword)
		user.POST("/:id/actions/change-password", s.userHandler.UpdatePassword)
		user.POST("/:id/actions/change-name", s.userHandler.UpdateName)
	}

	roomType := api.Group("/room-type")
	roomType.Use(s.RequireSession)
	{
		roomType.GET("", s.roomTypeHandler.GetAll)
		roomType.GET("/:id", s.roomTypeHandler.GetByID)
	}

	room := api.Group("/room")
	room.Use(s.RequireSession)
	{
		room.GET("", s.roomHandler.GetAll)
		room.GET("/:id", s.roomHandler.GetByID)
		room.GET("/filter", s.roomHandler.FilterByRoomType)
	}

	foodMaterial := api.Group("/food-material")
	foodMaterial.Use(s.RequireSession)
	{
		foodMaterial.GET("", s.foodMaterialHandler.GetAll)
		foodMaterial.GET("/:id", s.foodMaterialHandler.GetByID)
	}

	mealType := api.Group("/meal-type")
	mealType.Use(s.RequireSession)
	{
		mealType.GET("", s.mealTypeHandler.GetAll)
		mealType.GET("/:id", s.mealTypeHandler.GetByID)
	}

	food := api.Group("/food")
	food.Use(s.RequireSession)
	{
		food.GET("", s.foodHandler.GetAll)
		food.GET("/:id", s.foodHandler.GetByID)
	}

	patient := api.Group("/patient")
	patient.Use(s.RequireSession)
	{
		patient.POST("", s.patientHandler.Create)
		patient.GET("", s.patientHandler.GetAll)
		patient.GET("/:id", s.patientHandler.GetByID)
		patient.PATCH("/:id", s.patientHandler.Update)
		patient.DELETE("/:id", s.patientHandler.Delete)
		patient.GET("/filter", s.patientHandler.FilterByRMN)
		patient.GET("/paginated", s.patientHandler.GetAllWithPaginationAndKeyword)
		patient.GET("/find-from-simrs", s.patientHandler.FindFromSIMRS)
	}

	dailyPatientMeal := api.Group("/daily-patient-meal")
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
		dailyPatientMeal.POST("/copy-from-yesterday", s.dailyPatientMealHandler.CopyFromYesterday)
	}

	diet := api.Group("/diet")
	diet.Use(s.RequireSession)
	{
		diet.GET("", s.dietHandler.GetAll)
		diet.GET("/:id", s.dietHandler.GetByID)
	}

	allergy := api.Group("/allergy")
	allergy.Use(s.RequireSession)
	{
		allergy.GET("", s.allergyHandler.GetAll)
		allergy.GET("/:id", s.allergyHandler.GetByID)
	}

	admin := api.Group("/admin")
	admin.Use(s.RequireSession, s.RequireAdminRole)
	{
		user := admin.Group("/user")
		{
			user.GET("", s.userHandler.GetAll)
			user.GET("/:id", s.userHandler.GetByID)
			user.PATCH("/:id", s.userHandler.Update)
			user.DELETE("/:id", s.userHandler.Delete)
		}

		roomType := admin.Group("/room-type")
		{
			roomType.POST("", s.roomTypeHandler.Create)
			roomType.PATCH("/:id", s.roomTypeHandler.Update)
			roomType.DELETE("/:id", s.roomTypeHandler.Delete)
			roomType.POST("/simrs-sync", s.roomTypeHandler.SyncFromSIMRS)
		}

		room := admin.Group("/room")
		{
			room.POST("", s.roomHandler.Create)
			room.PATCH("/:id", s.roomHandler.Update)
			room.DELETE("/:id", s.roomHandler.Delete)
		}

		foodMaterial := admin.Group("/food-material")
		{
			foodMaterial.POST("", s.foodMaterialHandler.Create)
			foodMaterial.PATCH("/:id", s.foodMaterialHandler.Update)
			foodMaterial.DELETE("/:id", s.foodMaterialHandler.Delete)
		}

		mealType := admin.Group("/meal-type")
		{
			mealType.POST("", s.mealTypeHandler.Create)
			mealType.PATCH("/:id", s.mealTypeHandler.Update)
			mealType.DELETE("/:id", s.mealTypeHandler.Delete)
		}

		food := admin.Group("/food")
		{
			food.POST("", s.foodHandler.Create)
			food.PATCH("/:id", s.foodHandler.Update)
			food.DELETE("/:id", s.foodHandler.Delete)
		}

		diet := admin.Group("/diet")
		{
			diet.POST("", s.dietHandler.Create)
			diet.PATCH("/:id", s.dietHandler.Update)
			diet.DELETE("/:id", s.dietHandler.Delete)
		}

		allergy := admin.Group("/allergy")
		{
			allergy.POST("", s.allergyHandler.Create)
			allergy.PATCH("/:id", s.allergyHandler.Update)
			allergy.DELETE("/:id", s.allergyHandler.Delete)
		}
	}

	return r
}
