package manage

import (
	"fmt"
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/usecase/manage"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type manageHandler struct {
	manageUseCase manage.ManageUseCase
}

const (
	invalidBodyRequestMsg       = "invalid body request: %s"
	corruptedBodyRequestMsg     = "corrupted body request: %s"
	failedToParseBodyRequestMsg = "failed to parse body request"
	logoutSuccessMsg            = "user successfully logged out"
)

func NewManageHandler(e *echo.Echo, manageUseCase manage.ManageUseCase) Handler {
	h := &manageHandler{manageUseCase}
	h.manageUseCase = manageUseCase
	userGroup := e.Group("/manage")
	userGroup.POST("/create", h.CreateUrlRequest)
	return h
}

func (h *manageHandler) CreateUrlRequest(c echo.Context) error {
	var createUrlRequest *request.CreateUrlRequest
	log.Infof("acaaaaa")

	if err := c.Bind(&createUrlRequest); err != nil {
		// log.Errorf(corruptedBodyRequestMsg, err)
		return echo.NewHTTPError(http.StatusBadRequest, failedToParseBodyRequestMsg)
	}

	// if err := c.Bind(project); err != nil {
	//     fmt.Println("Error happened due to::", err)
	//     return err
	// }

	fmt.Println(createUrlRequest)

	log.Info(createUrlRequest.Test)

	return c.JSON(http.StatusOK, "ook")
}
