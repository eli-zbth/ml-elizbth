package application

import (
	"context"
	"fmt"
	"ml-elizabeth/app/infrastructure/http/rest/handlers/health"
	manageRouter "ml-elizabeth/app/infrastructure/http/rest/handlers/manage"
	megaLost "ml-elizabeth/app/infrastructure/lost"
	"ml-elizabeth/app/shared/utils/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ml-elizabeth/app/usecase/manage"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	log "github.com/sirupsen/logrus"
)

const tenSecondRule = 10 * time.Second

// @title UrlShortener API
// @description This is an authentication API used to register create,edit and redirect short url's.This was developed in response to the ML's technical challenge
// @version 1.0
// @contact.name API Support
// @contact.email eli.carren07@gmail.com
// @BasePath /v1
func StartApp() {
	logger.InitLoggerConfig()
	log := logger.New()
	log.WithFields(logger.Fields{
		"layer":  "main",
		"method": "StartApp",
	})

	log.Info("Starting application")

	e := echo.New()

	log.Info("Setup dependencies")
	setupDependencies(e)

	if os.Getenv("ENV") != "production" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	e.Use(echoPrometheus.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	quit := make(chan os.Signal, 1)
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func startServer(e *echo.Echo, quit chan os.Signal) {
	log.Println("Starting server")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Errorln(err.Error())
		close(quit)
	}
}

func gracefulShutdown(e *echo.Echo) {
	log.Println("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), tenSecondRule)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func setupDependencies(e *echo.Echo) {
	logEntry := logger.New()
	logEntry.WithFields(logger.Fields{
		"layer":  "main",
		"method": "setupDependencies",
	})

	randomRepository := megaLost.NewLost()

	manageUseCase := manage.NewManageUseCase(randomRepository)

	manageRouter.NewManageHandler(e, manageUseCase)
	health.NewHealthHandler(e)
}
