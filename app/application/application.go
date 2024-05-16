package application

import (
	"context"
	"fmt"

	createRouter "ml-elizabeth/app/infrastructure/http/rest/handlers/create"
	editRouter "ml-elizabeth/app/infrastructure/http/rest/handlers/edit"
	healthRouter "ml-elizabeth/app/infrastructure/http/rest/handlers/health"
	redirectRouter "ml-elizabeth/app/infrastructure/http/rest/handlers/redirect"

	m "ml-elizabeth/app/infrastructure/http/rest/middelware"

	"ml-elizabeth/app/shared/config"
	"ml-elizabeth/app/shared/utils/logger"
	"ml-elizabeth/app/shared/validations"
	"os"
	"os/signal"
	"syscall"
	"time"

	mongoConnector "ml-elizabeth/app/infrastructure/mongodb/connection"
	redisConnector "ml-elizabeth/app/infrastructure/redis/connection"
	StorageRepository "ml-elizabeth/app/infrastructure/mongodb/url"
	CacheRepository "ml-elizabeth/app/infrastructure/redis/url"
	"ml-elizabeth/app/usecase/create"
	"ml-elizabeth/app/usecase/edit"
	"ml-elizabeth/app/usecase/redirect"

	"github.com/go-playground/validator/v10"

	echoPrometheus "github.com/globocom/echo-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "ml-elizabeth/docs/swagger"

	log "github.com/sirupsen/logrus"
)

const tenSecondRule = 10 * time.Second

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func StartApp() {
	logger.InitLoggerConfig()
	log := logger.NewLogger()
	log.WithFields(logger.Fields{
		"layer":  "application",
		"method": "start_pp",
	})
	log.Info("Starting application setup")


	config.LoadConfig()

	e := echo.New()
	log.Info("Setup dependencies")
	setupDependencies(e)

	// if os.Getenv("ENV") != "production" {
		e.GET("/docs/swagger/*", echoSwagger.WrapHandler)
	// }

	e.Use(echoPrometheus.MetricsMiddleware())
	e.Use(m.Traceability)

	
	v, err := validations.NewCustomValidator(validator.New())
	if err != nil {
		log.Errorf("could not create custom validator: %v", err)
	}
	e.Validator = v

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	quit := make(chan os.Signal, 1)

	log.Info("Starting application setup")
	go startServer(e, quit)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	gracefulShutdown(e)
}

func startServer(e *echo.Echo, quit chan os.Signal) {
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
	mongoDbConn := mongoConnector.NewMongoConnector()
	storageRepository := StorageRepository.NewMongoOptionsRepository(mongoDbConn)   

	redisConn := redisConnector.NewRedisConnector()
	cacheRepository := CacheRepository.NewRedisRepository(redisConn);

	createUseCase := create.NewCreateUsecase(storageRepository,cacheRepository)
	editUseCase := edit.NewEditUseCase(storageRepository,cacheRepository)
	redirectUseCase := redirect.NewRedirectUseCase(storageRepository,cacheRepository)

	healthRouter.NewHealthHandler(e)
	createRouter.NewCreateHandler(e,createUseCase)
	editRouter.NewEditHandler(e,editUseCase)
	redirectRouter.NewRedirecttHandler(e, redirectUseCase)
}
