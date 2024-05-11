package manage

import "github.com/labstack/echo/v4"

type Handler interface {
	CreateUrlRequest(c echo.Context) error
}
