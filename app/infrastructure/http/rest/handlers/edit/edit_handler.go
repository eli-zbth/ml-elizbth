package edit

import (
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/context"
	"ml-elizabeth/app/usecase/edit"
	"net/http"
	"github.com/labstack/echo/v4"
)

type EditHandlerInterface interface {
	EditShortUrlRequest(c echo.Context) error
	EditRedirectUrlRequest(c echo.Context) error
	EditUrlStatus(c echo.Context) error
}

type EditHandler struct {
	editUseCase edit.Edit
}

func NewEditHandler(e *echo.Echo, editUseCase edit.Edit) EditHandlerInterface {
	h := &EditHandler{editUseCase}
	h.editUseCase = editUseCase
	userGroup := e.Group("/edit")
	userGroup.POST("/short_url", h.EditShortUrlRequest)
	userGroup.POST("/redirect_url", h.EditRedirectUrlRequest)
	userGroup.POST("/url_status", h.EditUrlStatus)
	return h
}


// EditShortUrl godoc
// @Summary      Edit Short url
// @Description  Change id for shorturl and can customizate it 
// @Tags         Edit
// @Accept       json
// @Produce      json
// @Param bodyRequest body request.EditShortUrlRequest true "shortUrl newUrl"
// @Success      200              {string}  string    "url updated successfully"
// @Failure      400  {object}  echo.HTTPError{message=string}
// @Failure      500  {object}  echo.HTTPError{message=string}
// @Router       /edit/short_url [post]
func (e *EditHandler) EditShortUrlRequest(c echo.Context) error {
	var req *request.EditShortUrlRequest
	ctx := context.NewContext(c)
	log := ctx.SetupLogger("edit_usecase","edit_short_url_request")

	if err := c.Bind(&req); err != nil {
		log.Errorf(constants.CorruptedBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, constants.FailedToParseBodyRequestMsg )
	}

	err := c.Validate(req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = e.editUseCase.EditShortUrl(ctx,req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Infof("URL successfully Edited: %s",req.NewUrl)
	return c.JSON(http.StatusOK, constants.UpdateUrlSuccess)
}


// EditRedirect godoc
// @Summary      Edit Redirect url
// @Description  Change web url assigned to a shorturl
// @Tags         Edit
// @Accept       json
// @Produce      json
// @Param bodyRequest body request.EditRedirectURLRequest  true "shortUrl RedirectUrl"
// @Success      200              {string}  string    "url updated successfully"
// @Failure      400  {object}  echo.HTTPError{message=string}
// @Failure      500  {object}  echo.HTTPError{message=string}
// @Router       /edit/redirect_url [post]
func (e *EditHandler) EditRedirectUrlRequest(c echo.Context) error {
	var req *request.EditRedirectURLRequest
	ctx := context.NewContext(c)
	log := ctx.SetupLogger("edit_usecase","edit_redirect_url_request")

	if err := c.Bind(&req); err != nil {
		log.Errorf(constants.CorruptedBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, constants.FailedToParseBodyRequestMsg )
	}

	err := c.Validate(req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = e.editUseCase.EditRedirectUrl(ctx,req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Infof("Redirect URL successfully Edited: %s", req.RedirectUrl)
	return c.JSON(http.StatusOK, constants.UpdateUrlSuccess)
}

// EditUrlStatus godoc
// @Summary      Edit url status
// @Description  Active/deactive short url
// @Tags         Edit
// @Accept       json
// @Produce      json
// @Param bodyRequest body request.EditUrlStatusRequest true "shortUrl IsActive"
// @Failure      400  {object}  echo.HTTPError{message=string}
// @Failure      500  {object}  echo.HTTPError{message=string}
// @Router       /edit/url_status [post]
func (e *EditHandler) EditUrlStatus(c echo.Context) error {
	var req *request.EditUrlStatusRequest
	ctx := context.NewContext(c)
	log := ctx.SetupLogger("edit_usecase","edit_url_status_request")

	if err := c.Bind(&req); err != nil {
		log.Errorf(constants.CorruptedBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, constants.FailedToParseBodyRequestMsg )
	}

	err := c.Validate(req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = e.editUseCase.EditUrlStatus(ctx,req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Infof("URL status was successfully Edited: %s", req.ShortUrl)
	return c.JSON(http.StatusOK, constants.UpdateUrlStatusSucess)

}

