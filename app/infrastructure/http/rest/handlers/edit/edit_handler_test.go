package edit

import (
	"errors"

	"ml-elizabeth/app/shared/validations"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ml-elizabeth/app/usecase/edit"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestEditShortUrlRequest(t *testing.T) {

    url = "/edit/short_url"
    method = http.MethodPost


    t.Run("Should return bad request when has invalid body  ", func(t *testing.T) {

        body :=  `}}{{{}}`

        e, echoContext,_  := setupEchoTest(body,method,url)
 
        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditShortUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "failed to parse body request")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when all field are missing ", func(t *testing.T) {
        
        body := `{"url": "https://www.domain.cl/valid-url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditShortUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when NewUrl field missing ", func(t *testing.T) {
        
        body := `{"short_url": "https://www.domain.cl/valid-url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditShortUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "NewUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when short_url is invalid URL", func(t *testing.T) {
        
        body := `{"short_url": "invalid_url" , "new_url": "new_url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditShortUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl must be a valid URL")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when new is less than 3 chars", func(t *testing.T) {
        
        body := `{"short_url": "http://www.tiny.com/tiny_url" , "new_value": "da"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditShortUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "NewUrl must be at least 3 characters in length")
		assert.Equal(t, errExpected, err)        
    })


    t.Run("Should return error 500 when an error occurs inside the usecase", func(t *testing.T) {

        body := `{"short_url": "http://www.tiny.com/tiny_url", "new_value": "new_tiny_url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        useCaseMock.On("EditShortUrl",mock.Anything, mock.Anything).Return(errors.New("Unexpected Error"))
      
		createHandler :=   NewEditHandler(e, useCaseMock)
        err := createHandler.EditShortUrlRequest(echoContext)
		
        errExpected:= echo.NewHTTPError(http.StatusInternalServerError, "Unexpected Error")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return 200 when short url was updated successfully", func(t *testing.T) {

        body := `{"short_url": "http://www.tiny.com/tiny_url", "new_value": "new_tiny_url"}`

        e, echoContext,rec  := setupEchoTest(body,method,url)


        useCaseMock:= new(edit.EditUseCaseMock)
        useCaseMock.On("EditShortUrl",mock.Anything, mock.Anything).Return(nil)
      
    
		createHandler :=   NewEditHandler(e, useCaseMock)
        err := createHandler.EditShortUrlRequest(echoContext)

        assert.Equal(t, err, nil)
        assert.Equal(t, http.StatusOK, rec.Code)
    })

}

func TestEditRedirectUrlRequest(t *testing.T) {
    url = "/edit/redirect_url"
    method = http.MethodPost

    t.Run("Should return bad request when has invalid body  ", func(t *testing.T) {

        body :=  `}}{{{}}`

        e, echoContext,_  := setupEchoTest(body,method,url)
 
        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "failed to parse body request")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when all field are missing ", func(t *testing.T) {
        
        body := `{"url": "https://www.domain.cl/valid-url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when short_url field is missing ", func(t *testing.T) {
        
        body := `{"redirect_url": "https://www.domain.cl/valid-url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })


    t.Run("Should return bad request when redirect field is missing ", func(t *testing.T) {
        
        body := `{"short_url": "https://www.tiny.cl/tiny"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "RedirectUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when short_url is invalid URL ", func(t *testing.T) {
        
        body := `{"short_url": "bad_url" , "redirect_url":"http://dominio.com/ultra_long_redirect_url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl must be a valid URL")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when redirect_url is invalid URL ", func(t *testing.T) {
        
        body := `{"short_url": "https://www.tiny.cl/tiny" , "redirect_url":"htt_dfeueBadurl"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "RedirectUrl must be a valid URL")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return error 500 when an error occurs inside the usecase", func(t *testing.T) {
 
        body := `{"short_url": "http://www.tiny.com/tiny_url", "redirect_url": "http://www.dominio.com/new_redirect_url"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        useCaseMock.On("EditRedirectUrl",mock.Anything, mock.Anything).Return(errors.New("Unexpected Error"))
      
		handler := NewEditHandler(e, useCaseMock)
        err := handler.EditRedirectUrlRequest(echoContext)
		
        errExpected:= echo.NewHTTPError(http.StatusInternalServerError, "Unexpected Error")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return 200 when redirect url was updated successfully", func(t *testing.T) {
        body := `{"short_url": "http://www.tiny.com/tiny_url", "redirect_url": "http://www.dominio.com/new_redirect_url"}`
        e, echoContext,rec  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        useCaseMock.On("EditRedirectUrl",mock.Anything, mock.Anything).Return(nil)
      
		createHandler :=   NewEditHandler(e, useCaseMock)
        err := createHandler.EditRedirectUrlRequest(echoContext)

        assert.Equal(t, err, nil)
        assert.Equal(t, http.StatusOK, rec.Code)
    })

}

func TestEditUrlStatusRequest(t *testing.T) {
    url = "/edit/url_status"
    method = http.MethodPost

    t.Run("Should return bad request when has invalid body  ", func(t *testing.T) {

        body :=  `}}{{{}}`

        e, echoContext,_  := setupEchoTest(body,method,url)
 
        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "failed to parse body request")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when all field are missing ", func(t *testing.T) {
        
        body := `{"any": "any_value"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when short_url field is missing ", func(t *testing.T) {
        
        body := `{"is_active": true}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl is a required field")
		assert.Equal(t, errExpected, err)        
    })


    t.Run("Should return bad request when is_active field is missing ", func(t *testing.T) {
        
        body := `{"short_url": "https://www.tiny.cl/tiny"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "IsActive is a required field")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when short_url is invalid URL ", func(t *testing.T) {
        
        body := `{"short_url": "bad_url" , "is_active": false}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "ShortUrl must be a valid URL")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when is_active is invalid bool ", func(t *testing.T) {
        
        body := `{"short_url": "https://www.tiny.cl/tiny" , "is_active":"true"}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)

        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "failed to parse body request")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return error 500 when an error occurs inside the usecase", func(t *testing.T) {
 
        body := `{"short_url": "https://www.tiny.cl/tiny" , "is_active":true}`

        e, echoContext,_  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        useCaseMock.On("EditUrlStatus",mock.Anything, mock.Anything).Return(errors.New("Unexpected Error"))
      
		handler := NewEditHandler(e, useCaseMock)
        err := handler.EditUrlStatus(echoContext)
		
        errExpected:= echo.NewHTTPError(http.StatusInternalServerError, "Unexpected Error")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return 200 when redirect url was updated successfully", func(t *testing.T) {
        body := `{"short_url": "https://www.tiny.cl/tiny" , "is_active":true}`
        e, echoContext,rec  := setupEchoTest(body,method,url)

        useCaseMock:= new(edit.EditUseCaseMock)
        useCaseMock.On("EditUrlStatus",mock.Anything, mock.Anything).Return(nil)
      
		createHandler :=   NewEditHandler(e, useCaseMock)
        err := createHandler.EditUrlStatus(echoContext)

        assert.Equal(t, err, nil)
        assert.Equal(t, http.StatusOK, rec.Code)
    })

}