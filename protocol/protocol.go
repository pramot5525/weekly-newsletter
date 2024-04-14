package protocol

import (
	"fmt"
	"log"
	"time"
	"weekly-newsletter/internal/handler"
	"weekly-newsletter/internal/repository"
	"weekly-newsletter/internal/service"
	"weekly-newsletter/pkg/config"
	"weekly-newsletter/pkg/database"
	router "weekly-newsletter/route"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type svc struct {
	app   *fiber.App
	db    *gorm.DB
	close func()
}

func ServeHTTP() error {
	svc, err := setup()
	if err != nil {
		log.Fatalf("error setting up service: %v", err)
	}
	defer svc.close()

	err = svc.app.Listen(":" + config.GetString(config.AppPort))
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

	err = svc.app.ShutdownWithTimeout(10 * time.Second)
	if err != nil {
		log.Fatal("error shutting down server", err)
	}

	fmt.Println("shutdown gracefully")

	return nil
}

func setup() (*svc, error) {
	err := config.Loadenv(".env")
	if err != nil {
		return nil, fmt.Errorf("setup: %w", err)
	}

	db, err := database.ConnectPostgres(
		config.GetString(config.PostgresHost),
		config.GetString(config.PostgresPort),
		config.GetString(config.PostgresUser),
		config.GetString(config.PostgresPassword),
		config.GetString(config.PostgresDatabase),
	)

	if err != nil {
		return nil, fmt.Errorf("setup: %w", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		return nil, fmt.Errorf("setup: %w", err)
	}

	app := fiber.New()

	repo := repository.New(db)
	service := service.NewUserService(repo)
	handler := handler.New(service)
	router.NewRouter(app, handler)

	return &svc{
		app: app,
		db:  db,
		close: func() {
			database.DisconnectPostgres(db)
		},
	}, nil
}
