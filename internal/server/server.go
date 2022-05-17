package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	"github.com/bangadam/go-fiber-boilerplate/pkg/logger"
	postgress "github.com/bangadam/go-fiber-boilerplate/pkg/postgres"
	"github.com/bangadam/go-fiber-boilerplate/pkg/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"

	baseApp "github.com/bangadam/internal/app"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	viperPkg "github.com/spf13/viper"
)

func Run() {
	// Load config
	_, b, _, _ := runtime.Caller(0)
	configPath := filepath.Join(filepath.Dir(b), "../..")
	config := &viper.EnvConfig{
		FileName: "config",
		FileType: "yaml",
		Path:     configPath,
		IdleTimeout: viperPkg.GetDuration("server.idleTimeout"),
	}
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}

	// Load connection postgres
	pg, err := postgress.Connect()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := pg.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	// Logger steup
	zap, err := logger.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	// Load Fiber
	app := fiber.New(fiber.Config{
		IdleTimeout: config.IdleTimeout,
	})

	app.Use(
		recover.New(),
		compress.New(),
		etag.New(),
		cors.New(),
		fiberlog.New(),
	)
	
	// set Handler
	rh := &baseApp.Handler{
		Postgres: sqlDB,
		R: app,
		Logger: zap,
	}

	rh.SetRoutes()

	// Listen and serve
	go func () {
		if err := app.Listen(viperPkg.GetString("server.port")); err != nil {
			log.Panic("Failed to start server: ", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	var _ = <- c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup taks...")
	sqlDB.Close()
	fmt.Println("Service was stopped")
}