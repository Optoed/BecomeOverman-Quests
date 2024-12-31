package main

import (
	"BecomeOverMan/internal/handlers"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	_ "BecomeOverMan/internal/models"
	"BecomeOverMan/internal/repositories"
	"BecomeOverMan/internal/services"

	_ "github.com/golang-migrate/migrate/v4/source/file" // Импорт для работы с миграциями через файлы
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Настройка миграций
	m, err := migrate.New(
		"file://migrations", // Путь к папке с миграциями
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal("Error initializing migration:", err)
	}

	// Применение миграций
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Failed to apply migrations:", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	baseRepo := repositories.NewBaseRepository(db)
	baseService := services.NewBaseService(baseRepo)
	baseHandler := handlers.NewBaseHandler(baseService)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", baseHandler.CheckConnection)

		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
