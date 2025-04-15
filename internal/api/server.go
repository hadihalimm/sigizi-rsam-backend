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
	port        int
	db          *config.Database
	authHandler *handler.AuthHandler
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := config.ConnectToDatabase()
	db.Gorm.Migrator().DropTable(&model.User{})
	db.Gorm.AutoMigrate(&model.User{})

	validator := validator.New()
	userRepo := repo.NewUserRepo(db)
	authService := service.NewAuthService(userRepo, validator)
	authHandler := handler.NewAuthHandler(authService)

	NewServer := &Server{
		port:        port,
		db:          db,
		authHandler: authHandler,
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
