package redirect

import (
	"ml-elizabeth/app/shared/validations"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

    "ml-elizabeth/app/usecase/redirect"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)


func setupEchoTest(body string, method string, url string) (*echo.Echo, echo.Context, *httptest.ResponseRecorder ){
    e := echo.New()
    e.Validator, _ = validations.NewCustomValidator(validator.New())
    req := httptest.NewRequest(method, url, strings.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

    rec := httptest.NewRecorder()
    echoContext := e.NewContext(req, rec)
    return e,echoContext,rec
}

var url string
var method string

func TestRedirectUrlRequest(t *testing.T) {
    url = "/"
    method = http.MethodGet

    t.Run("Should return bad request when has invalid body", func(t *testing.T) {
        body :=  `}}{{{}}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(redirect.RedirectUseCaseMock)
        handler := NewRedirecttHandler(e, useCaseMock)
        err := handler.RedirectRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "Id is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request id's length is less than 3 chars", func(t *testing.T) {

        body :=  `}}{{{}}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(redirect.RedirectUseCaseMock)
        handler := NewRedirecttHandler(e, useCaseMock)
        err := handler.RedirectRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "Id is a required field")
		assert.Equal(t, errExpected, err)        
    })

}