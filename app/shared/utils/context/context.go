package context

import (
	"context"

	logger "ml-elizabeth/app/shared/utils/logger"
	m "ml-elizabeth/app/infrastructure/http/rest/middelware"

	"github.com/labstack/echo/v4"
)

const ctxKey int = 0

type Context struct {
	Ctx context.Context
	Log logger.Logger
}



func NewContext(c echo.Context) *Context {
	log := logger.NewLogger()
	newContext, ok := c.(*m.EchoCtxCustom)
	if ok {
		log = newContext.Log
	}
	return &Context{
		Ctx: context.WithValue(context.Background(), ctxKey, log),
		Log: log,
	}
}


func (c *Context) SetupLogger(layer string ,method string) (logger.Logger) {
	c.Log.WithFields(logger.Fields{
		"layer":  layer,
		"method": method,
	})
	return c.Log
}



