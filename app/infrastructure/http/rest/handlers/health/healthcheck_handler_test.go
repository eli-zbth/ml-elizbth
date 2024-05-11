package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHealthHandler(t *testing.T) {
	t.Run("when Healthcheck response is OK", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

		NewHealthHandler(e)
		healthcheck := healthHandler{}

		_ = healthcheck.HealthCheck(echoContext)

		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
