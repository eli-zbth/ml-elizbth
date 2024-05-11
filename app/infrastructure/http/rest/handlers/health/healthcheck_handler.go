package health

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type healthHandler struct {
}

type health struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
	Region      string `json:"region"`
}

func NewHealthHandler(e *echo.Echo) {
	h := &healthHandler{}

	e.GET("/health", h.HealthCheck)
}

func (p *healthHandler) HealthCheck(c echo.Context) error {

	healthCheck := health{
		Status:      "UP",
		Version:     os.Getenv("APP_VERSION"),
		Environment: os.Getenv("ENV"),
		Region:      os.Getenv("REGION"),
	}

	return c.JSON(http.StatusOK, healthCheck)
}
