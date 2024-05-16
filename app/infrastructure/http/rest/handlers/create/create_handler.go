package create

import (
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/shared/utils/constants"
    "ml-elizabeth/app/usecase/create"
	"ml-elizabeth/app/shared/utils/context"
	"net/http"


	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
    createUseCase   create.Create
}


type CreateHandlerInterface interface {
	CreateUrlRequest(c echo.Context) error
}

func NewCreateHandler(e *echo.Echo, createUseCase create.Create) CreateHandlerInterface {
	h := &CreateHandler{createUseCase}
	h.createUseCase = createUseCase
	e.POST("/create", h.CreateUrlRequest)
	return h
}

// CreateUrl godoc
// @Summary      Create short Url
// @Description  Take the Url and assing a short id, then save the data in database and cache	
// @Tags         Create
// @Accept       json
// @Produce      json
// @Param bodyRequest body request.CreateUrlRequest true "url CustomId"
// @Success      200  {object}  response.CreateUrlResponse
// @Failure      400  {object}  echo.HTTPError{message=string}
// @Failure      500  {object}  echo.HTTPError{message=string}
// @Router       /create [post]
func (h *CreateHandler) CreateUrlRequest(c echo.Context) error {
	var req *request.CreateUrlRequest
	ctx := context.NewContext(c)
	log := ctx.SetupLogger("create_usecase","create_url_request")
    
	if err := c.Bind(&req); err != nil {
		log.Errorf(constants.CorruptedBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, constants.FailedToParseBodyRequestMsg )
	}

	err := c.Validate(req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.createUseCase.CreateUrl(ctx, req)
	if err != nil {
		log.Errorf(constants.InvalidBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Infof("URL created: %s",resp.ShortUrl)
	return c.JSON(http.StatusOK, resp)
}
