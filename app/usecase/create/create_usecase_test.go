package create

import (
	"errors"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/spf13/viper"
	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/context"

	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/models/response"
	mongoMock "ml-elizabeth/app/infrastructure/mongodb/url"
    redisMock "ml-elizabeth/app/infrastructure/redis/url"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUseCase(t *testing.T) {

    usecaseReq:= &request.CreateUrlRequest{
        Url: "http://www.domain.com/longDomain",
        CustomId: "custom_id",
    }

    t.Run("Should return error when url already exists", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Save",mock.Anything, mock.Anything,mock.Anything).Return(errors.New(constants.DuplicateUrlError))

        cacheMock:= new(redisMock.RedisMock)
        useCase :=   NewCreateUsecase(storageMock,cacheMock)
        _,err := useCase.CreateUrl(c, usecaseReq)

        errExpected := errors.New("Cannot create short url, it already exists")
        assert.Equal(t, errExpected, err )     
           
    })

    t.Run("Should return error when redis fails", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Save",mock.Anything, mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Save",mock.Anything, mock.Anything,mock.Anything).Return(errors.New("Some Redis error"))

        useCase :=   NewCreateUsecase(storageMock,cacheMock)
        _,err := useCase.CreateUrl(c, usecaseReq)

        errExpected := errors.New("Some Redis error")
        assert.Equal(t, errExpected, err )     
           
    })

    t.Run("Should return new URL", func(t *testing.T) {
        viper.Set("SHORT_URL_DOMAIN", "http://api-domain.com/")
        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeEndpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Save",mock.Anything, mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Save",mock.Anything, mock.Anything,mock.Anything).Return(nil)

        useCase :=   NewCreateUsecase(storageMock,cacheMock)
        resp,err := useCase.CreateUrl(c, usecaseReq)

        expectedResp:= &response.CreateUrlResponse{
            ShortUrl: "http://api-domain.com/custom_id" ,
        }

        assert.Equal(t, nil, err )     
        assert.Equal(t, expectedResp, resp )     
           
    })
}
