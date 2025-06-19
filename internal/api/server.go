package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hadihalimm/sigizi-rsam/internal/config"
	"github.com/hadihalimm/sigizi-rsam/internal/handler"
	"github.com/hadihalimm/sigizi-rsam/internal/model"
	"github.com/hadihalimm/sigizi-rsam/internal/repo"
	"github.com/hadihalimm/sigizi-rsam/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port                    int
	db                      *config.Database
	authHandler             *handler.AuthHandler
	userHandler             *handler.UserHandler
	roomTypeHandler         *handler.RoomTypeHandler
	roomHandler             *handler.RoomHandler
	foodMaterialHandler     *handler.FoodMaterialHandler
	mealTypeHandler         *handler.MealTypeHandler
	foodHandler             *handler.FoodHandler
	patientHandler          *handler.PatientHandler
	dailyPatientMealHandler *handler.DailyPatientMealHandler
	dietHandler             *handler.DietHandler
	allergyHandler          *handler.AllergyHandler
}

func NewServer() *http.Server {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("No .env file found")
	}
	env := os.Getenv("SIGIZI_ENV")
	if env == "" {
		env = "dev"
	}
	envFile := ".env." + env
	err = godotenv.Load(envFile)
	if err != nil {
		log.Print("No %s file found", envFile)
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := config.ConnectToDatabase()
	config.InitSession()
	config.InitLogger()
	defer config.Logger.Sync()

	// db.Gorm.Migrator().DropTable(&model.User{},
	// 	&model.RoomType{}, &model.Room{},
	// 	&model.Food{}, &model.MealType{}, &model.MealItem{},
	// 	&model.Patient{}, &model.DailyPatientMeal{})
	db.Gorm.Migrator().DropTable(&model.FoodMaterialUsage{}, &model.Food{})

	db.Gorm.AutoMigrate(&model.User{},
		&model.RoomType{}, &model.Room{},
		&model.FoodMaterial{}, &model.MealType{}, &model.Food{}, &model.FoodMaterialUsage{},
		&model.Patient{}, &model.DailyPatientMeal{}, &model.DailyPatientMealLog{},
		&model.Diet{}, &model.Allergy{})

	validator := validator.New()

	userRepo := repo.NewUserRepo(db)
	authService := service.NewAuthService(userRepo, validator)
	userService := service.NewUserService(userRepo, validator)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	roomTypeRepo := repo.NewRoomTypeRepo(db)
	roomTypeService := service.NewRoomTypeService(roomTypeRepo, validator)

	roomRepo := repo.NewRoomRepo(db)
	roomService := service.NewRoomService(roomRepo, roomTypeRepo, validator)
	roomTypeHandler := handler.NewRoomTypeHandler(roomTypeService, roomService)
	roomHandler := handler.NewRoomHandler(roomService)

	foodMaterialRepo := repo.NewFoodMaterialRepo(db)
	foodMaterialService := service.NewFoodMaterialService(foodMaterialRepo, validator)
	foodMaterialHandler := handler.NewFoodMaterialHandler(foodMaterialService)

	mealTypeRepo := repo.NewMealTypeRepo(db)
	mealTypeService := service.NewMealTypeService(mealTypeRepo, validator)
	mealTypeHandler := handler.NewMealTypeHandler(mealTypeService)

	foodRepo := repo.NewFoodRepo(db)
	foodService := service.NewFoodService(foodRepo, validator)
	foodHandler := handler.NewFoodHandler(foodService)

	patientRepo := repo.NewPatientRepo(db)
	patientService := service.NewPatientService(patientRepo, validator)
	patientHandler := handler.NewPatientHandler(patientService)

	dailyPatientMealRepo := repo.NewDailyPatientMealRepo(db)
	dailyPatientMealService := service.NewDailyPatientMealService(dailyPatientMealRepo, roomTypeRepo, validator)
	dailyPatientMealHandler := handler.NewDailyPatientMealHandler(dailyPatientMealService)

	dietRepo := repo.NewDietRepo(db)
	dietService := service.NewDietService(dietRepo, validator)
	dietHandler := handler.NewDietHandler(dietService)

	allergyRepo := repo.NewAllergyRepo(db)
	allergyService := service.NewAllergyService(allergyRepo, validator)
	allergyHandler := handler.NewAllergyHandler(allergyService)

	NewServer := &Server{
		port:                    port,
		db:                      db,
		authHandler:             authHandler,
		userHandler:             userHandler,
		roomTypeHandler:         roomTypeHandler,
		roomHandler:             roomHandler,
		foodMaterialHandler:     foodMaterialHandler,
		mealTypeHandler:         mealTypeHandler,
		foodHandler:             foodHandler,
		patientHandler:          patientHandler,
		dailyPatientMealHandler: dailyPatientMealHandler,
		dietHandler:             dietHandler,
		allergyHandler:          allergyHandler,
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}
