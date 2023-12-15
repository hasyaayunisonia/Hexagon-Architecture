package main

import (
	"context"
	productsRepo "hexagon-architecture/internal/domain/products/repository"
	"os"
	"os/signal"
	"syscall"

	"hexagon-architecture/config"
	"hexagon-architecture/internal/api"
	productsDB "hexagon-architecture/internal/domain/products/repository/db"
	"hexagon-architecture/internal/infrastructure"
	"hexagon-architecture/internal/service"
	"hexagon-architecture/pkg/http"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
// Get config.
	cfg := config.GetConfig()

	// Init db.
	db, _ := config.NewDB(cfg.DB, cfg.App.Env)

	// Init products.
	var products productsRepo.Repository = productsDB.New(db, "products")

	// Init service.
	service := service.New(
		products,
	)

	ctx := context.Background()

	// Init Opentelemetry.
	otelShutdown, err := infrastructure.SetupOTelSDK(ctx, cfg.App.Name, cfg.App.Version, cfg.Otel.Host, cfg.App.Env)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = otelShutdown(ctx)
	}()

	// Init web server.
	httpServer := http.New(http.Config{
		Port:            cfg.App.Port,
		ReadTimeout:     cfg.App.ReadTimeout,
		WriteTimeout:    cfg.App.WriteTimeout,
		GracefulTimeout: cfg.App.GracefulTimeout,
	})
	r := httpServer.Router()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))
	r.Use(otelfiber.Middleware())

	// Register api route.
	api.New(service, cfg.App.Env).Register(r)

	// Run web server.
	httpServerChan := httpServer.Run()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-httpServerChan:
		if err != nil {
			return
		}
	case <-sigChan:
	}

}