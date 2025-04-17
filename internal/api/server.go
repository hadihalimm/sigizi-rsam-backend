package api

import (
	"fmt"
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
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port            int
	db              *config.Database
	authHandler     *handler.AuthHandler
	roomTypeHandler *handler.RoomTypeHandler
	roomHandler     *handler.RoomHandler
	foodHandler     *handler.FoodHandler
	mealTypeHandler *handler.MealTypeHandler
	mealItemHandler *handler.MealItemHandler
	patientHandler  *handler.PatientHandler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := config.ConnectToDatabase()

	db.Gorm.Migrator().DropTable(&model.User{},
		&model.RoomType{}, &model.Room{},
		&model.Food{}, &model.MealType{}, &model.MealItem{},
		&model.Patient{})

	db.Gorm.AutoMigrate(&model.User{},
		&model.RoomType{}, &model.Room{},
		&model.Food{}, &model.MealType{}, &model.MealItem{},
		&model.Patient{})

	validator := validator.New()

	userRepo := repo.NewUserRepo(db)
	authService := service.NewAuthService(userRepo, validator)
	authHandler := handler.NewAuthHandler(authService)

	roomTypeRepo := repo.NewRoomTypeRepo(db)
	roomTypeService := service.NewRoomTypeService(roomTypeRepo, validator)
	roomTypeHandler := handler.NewRoomTypeHandler(roomTypeService)

	roomRepo := repo.NewRoomRepo(db)
	roomService := service.NewRoomService(roomRepo, validator)
	roomHandler := handler.NewRoomHandler(roomService)

	foodRepo := repo.NewFoodRepo(db)
	foodService := service.NewFoodService(foodRepo, validator)
	foodHandler := handler.NewFoodHandler(foodService)

	mealTypeRepo := repo.NewMealTypeRepo(db)
	mealTypeService := service.NewMealTypeService(mealTypeRepo, validator)
	mealTypeHandler := handler.NewMealTypeHandler(mealTypeService)

	mealItemRepo := repo.NewMealItemRepo(db)
	mealItemService := service.NewMealItemService(mealItemRepo, validator)
	mealItemHandler := handler.NewMealItemHandler(mealItemService)

	patientRepo := repo.NewPatientRepo(db)
	patientService := service.NewPatientService(patientRepo, validator)
	patientHandler := handler.NewPatientHandler(patientService)

	NewServer := &Server{
		port:            port,
		db:              db,
		authHandler:     authHandler,
		roomTypeHandler: roomTypeHandler,
		roomHandler:     roomHandler,
		foodHandler:     foodHandler,
		mealTypeHandler: mealTypeHandler,
		mealItemHandler: mealItemHandler,
		patientHandler:  patientHandler,
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}
