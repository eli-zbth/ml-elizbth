package create

import (
	"encoding/json"
	"errors"
	"io"
	"ml-elizabeth/app/shared/validations"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ml-elizabeth/app/usecase/create"

	"ml-elizabeth/app/domain/models/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateHandlerValidations(t *testing.T) {

    t.Run("Should return bad request when request has invalid body  ", func(t *testing.T) {

        body :=  `}}{{{}}`

        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        useCaseMock:= new(create.CreateUseCaseMock)
        createHandler :=   NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)

       
        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "failed to parse body request")
		assert.Equal(t, errExpected, err)        
    })


    t.Run("Should return bad request when the body does not include the url parameter", func(t *testing.T) {

        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
        body := `{"someParamether": "email@email.com"}`

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        useCaseMock:= new(create.CreateUseCaseMock)
    
        createHandler :=   NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)

       
        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "Url is a required field")
		assert.Equal(t, errExpected, err)
       
    })

    t.Run("Should return bad request when the url Field is not a valid url", func(t *testing.T) {
        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
        body := `{"url": "email@email.com"}`

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        useCaseMock:= new(create.CreateUseCaseMock)
       
		
        createHandler :=  NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)

    
        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "Url must be a valid URL")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return bad request when custom_id length is less than 3 chars", func(t *testing.T) {
        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
        body := `{"url": "http://www.dominio.com/url" , "custom_url":"fd"}`

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        useCaseMock:= new(create.CreateUseCaseMock)
       
		
        createHandler :=  NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)

    
        errExpected:= echo.NewHTTPError(http.StatusBadRequest, "CustomId must be at least 3 characters in length")
		assert.Equal(t, errExpected, err)        
    })


    t.Run("Should return error 500 when an error occurs inside the usecase", func(t *testing.T) {
        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
        body := `{"url": "https://www.domain.cl/valid-url"}`

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        useCaseMock:= new(create.CreateUseCaseMock)
        useCaseMock.On("CreateUrl",mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error"))
      
      
		createHandler :=   NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)
		
        errExpected:= echo.NewHTTPError(http.StatusInternalServerError, "Unexpected Error")
		assert.Equal(t, errExpected, err)        
    })

    t.Run("Should return 200 when short uls was created successfully", func(t *testing.T) {
        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
        body := `{"url": "https://www.domain.cl/valid-url"}`

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        mockReponse :=  &response.CreateUrlResponse{
            ShortUrl: "http://www.fakedominio.col/tinyurl",
        }

        useCaseMock:= new(create.CreateUseCaseMock)
        useCaseMock.On("CreateUrl",mock.Anything, mock.Anything).Return(mockReponse, nil)
      
    
		createHandler :=   NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)


        var resp *response.CreateUrlResponse
        b, _ := io.ReadAll(rec.Body)
        _ = json.Unmarshal(b, &resp)

        assert.Equal(t, http.StatusOK, rec.Code)
		assert.EqualValues(t, nil, err)
    })

    t.Run("Should return 200 when short uls was created successfully with custom id", func(t *testing.T) {
        e := echo.New()
        e.Validator, _ = validations.NewCustomValidator(validator.New())
        body := `{"url": "https://www.domain.cl/valid-url" , "custom_url":"custom_url"}`

		req := httptest.NewRequest(http.MethodPost, "/create", strings.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)

        mockReponse :=  &response.CreateUrlResponse{
            ShortUrl: "http://www.fakedominio.col/custom_url",
        }

        useCaseMock:= new(create.CreateUseCaseMock)
        useCaseMock.On("CreateUrl",mock.Anything, mock.Anything).Return(mockReponse, nil)
    
		createHandler := NewCreateHandler(e, useCaseMock)
        err := createHandler.CreateUrlRequest(echoContext)

        var resp *response.CreateUrlResponse
        b, _ := io.ReadAll(rec.Body)
        _ = json.Unmarshal(b, &resp)

        assert.Equal(t, http.StatusOK, rec.Code)
		assert.EqualValues(t,nil,err)
    })

}