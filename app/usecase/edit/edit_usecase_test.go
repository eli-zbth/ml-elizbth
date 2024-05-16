package edit

import (
	"errors"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/context"

	"ml-elizabeth/app/domain/models/request"
	mongoMock "ml-elizabeth/app/infrastructure/mongodb/url"
    redisMock "ml-elizabeth/app/infrastructure/redis/url"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEditShortUrl(t *testing.T) {

    usecaseReq:= &request.EditShortUrlRequest{
        ShortUrl: "http://www.tiny.com/tiny",
        NewUrl: "new_tiny",
    }

    t.Run("Should return error when url dosen't exists", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(errors.New(constants.DosentExistsError))

        cacheMock:= new(redisMock.RedisMock)

        useCase :=  NewEditUseCase(storageMock,cacheMock)
        err := useCase.EditShortUrl(c, usecaseReq)

        errExpected := errors.New("Can't update the URL, it dosen't exists")
        assert.Equal(t, errExpected, err )     
           
    })

    t.Run("Should return error when Redis Fail", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateKey",mock.Anything, mock.Anything,mock.Anything).Return(errors.New("Some redis error"))

        useCase :=  NewEditUseCase(storageMock,cacheMock)
        err := useCase.EditShortUrl(c, usecaseReq)

        errExpected := errors.New("Some redis error")
        assert.Equal(t, errExpected, err )     
           
    })

    t.Run("Should return nil when url is updated", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateKey",mock.Anything, mock.Anything,mock.Anything).Return(nil)


        useCase :=  NewEditUseCase(storageMock,cacheMock )
        err := useCase.EditShortUrl(c, usecaseReq)

        assert.Equal(t, nil, err )     
           
    })

}

func TestEditRedirectUrl(t *testing.T) {

    usecaseReq:= &request.EditRedirectURLRequest{
        ShortUrl: "http://www.tiny.com/tiny",
        RedirectUrl: "http://wwww.domain.com/new_long_url_to_redirect",
    }

    t.Run("Should return error when url dosen't exists", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(errors.New(constants.DosentExistsError))

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateValue",mock.Anything, mock.Anything,mock.Anything).Return(nil)


        useCase :=  NewEditUseCase(storageMock,cacheMock )
        err := useCase.EditRedirectUrl(c, usecaseReq)

        errExpected := errors.New("Can't update the URL, it dosen't exists")
        assert.Equal(t, errExpected, err )     

    })

    t.Run("Should return error when redis Fail", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateValue",mock.Anything, mock.Anything,mock.Anything).Return(errors.New("Some Redis error"))


        useCase :=  NewEditUseCase(storageMock,cacheMock )
        err := useCase.EditRedirectUrl(c, usecaseReq)

        errExpected := errors.New("Some Redis error")
        assert.Equal(t, errExpected, err )     

    })

    t.Run("Should return nil when url is updated", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateValue",mock.Anything, mock.Anything,mock.Anything).Return(nil)


        useCase :=  NewEditUseCase(storageMock,cacheMock )
        err := useCase.EditRedirectUrl(c, usecaseReq)

        assert.Equal(t, nil, err )     
           
    })

}

func BoolPointer(b bool) *bool {
    return &b
}

func TestEditUrlStatus(t *testing.T) {
    usecaseReq:= &request.EditUrlStatusRequest{
        ShortUrl: "http://www.tiny.com/tiny",
        IsActive: BoolPointer(true) ,
    }
    
    t.Run("Should return error when url dosen't exists", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(errors.New(constants.DosentExistsError))

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateValue",mock.Anything, mock.Anything,mock.Anything).Return(nil)

        useCase :=  NewEditUseCase(storageMock,cacheMock)
        err := useCase.EditUrlStatus(c, usecaseReq)

        errExpected := errors.New("Can't update the URL, it dosen't exists")
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
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateValue",mock.Anything, mock.Anything,mock.Anything).Return(errors.New("Some redis error"))

        useCase :=  NewEditUseCase(storageMock,cacheMock)
        err := useCase.EditUrlStatus(c, usecaseReq)

        errExpected := errors.New("Some redis error")
        assert.Equal(t, errExpected, err )     

    })

    t.Run("Should return nil when url is updated", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Update",mock.Anything, mock.Anything,mock.Anything,mock.Anything).Return(nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("UpdateValue",mock.Anything, mock.Anything,mock.Anything).Return(nil)

        useCase :=  NewEditUseCase(storageMock,cacheMock)
        err := useCase.EditUrlStatus(c, usecaseReq)

    
        assert.Equal(t, nil, err )     

    })
}