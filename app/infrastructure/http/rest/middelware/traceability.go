package middleware

import (
	"ml-elizabeth/app/shared/utils/trace"
	"ml-elizabeth/app/shared/utils/logger"
	"github.com/labstack/echo/v4"

)


type EchoCtxCustom struct {
	echo.Context
	Log logger.Logger
}

func NewEchoCtxCustom(c echo.Context, log logger.Logger) *EchoCtxCustom {
	return &EchoCtxCustom{
		c,
		log,
	}
}



func Traceability(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		traceparent := c.Request().Header.Get("traceparent")

		lenSpanID := 16
		spanID := trace.NewSpanID(lenSpanID)
		t := trace.New(spanID, traceparent)
		loggerTrace := logger.NewWithTrace(t)
		cc := NewEchoCtxCustom(c, loggerTrace)

		return next(cc)
		
	}
}
