package redirect

import (

	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/context"
	"ml-elizabeth/app/usecase/redirect"

	"net/http"

	"github.com/labstack/echo/v4"
)


type RedirectHandlerInterface interface {
	RedirectRequest(c echo.Context) error
}

type RedirecttHandler struct {
	redirectUseCase redirect.Redirect
}

func NewRedirecttHandler(e *echo.Echo, redirectUseCase redirect.Redirect) RedirectHandlerInterface{
	h := &RedirecttHandler{redirectUseCase}
	h.redirectUseCase =redirectUseCase
	e.GET("/:id", h.RedirectRequest)
	return h
}


// Redirect godoc
// @Summary      Redirect
// @Description  Get shorturl and redirect to web url
// @Tags         Redirect
// @Accept       json
// @Produce      json
// @Param path body request.RedirectRequest true "Id"
// @Success      302  {string}  string    "Url status updated successfully"
// @Failure      400  {object}  echo.HTTPError{message=string}
// @Failure      500  {object}  echo.HTTPError{message=string}
// @Router       / [get]
func (r *RedirecttHandler) RedirectRequest(c echo.Context) error {
	ctx := context.NewContext(c)
	log := ctx.SetupLogger("redirect_usecase","redirect_request")

	req:=&request.RedirectRequest{
		Id: c.Param("id"),
	}
	err := c.Validate(req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url,err := r.redirectUseCase.Redirect(ctx,req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Infof("Redirect successfully")
	return c.Redirect(http.StatusTemporaryRedirect,url)

}