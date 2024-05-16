package redirect

import (
	"errors"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/context"

	"ml-elizabeth/app/domain/entity"
	"ml-elizabeth/app/domain/models/request"
	mongoMock "ml-elizabeth/app/infrastructure/mongodb/url"
    redisMock "ml-elizabeth/app/infrastructure/redis/url"


	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRedirectUseCase(t *testing.T) {

    usecaseReq:= &request.RedirectRequest{
        Id: "tiny_id",
    }

    t.Run("Should return error when url dosen't exists", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Find",mock.Anything,mock.Anything).Return(nil,errors.New(constants.DosentExistsError))

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Find",mock.Anything,mock.Anything).Return(nil,errors.New(constants.DosentExistsError))

        useCase :=  NewRedirectUseCase(storageMock,cacheMock)
        resp,err := useCase.Redirect(c, usecaseReq)

        errExpected := errors.New("Can't update the URL, it dosen't exists")
        assert.Equal(t, errExpected, err )     
        assert.Equal(t, "", resp )
           
    })

    t.Run("Should return error when exists en Redis but url is not active", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        mockResponse := &entity.UrlRegistry{
            URL     :"http://wwww.domain.com/original_long_url",
            Key     :"tiny_id" ,
            Active  : false,
        }
        
        
        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Find",mock.Anything,mock.Anything).Return(nil,nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Find",mock.Anything,mock.Anything).Return(mockResponse,nil)

        useCase :=  NewRedirectUseCase(storageMock,cacheMock)
        resp,err := useCase.Redirect(c, usecaseReq)

        assert.Equal(t, errors.New("Can't redirect, the url expired"), err )     
        assert.Equal(t, "", resp )     
           
    })

    t.Run("Should return error when dosent's exist in redis but exist in database and url is not active", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        mockResponse := &entity.UrlRegistry{
            URL     :"http://wwww.domain.com/original_long_url",
            Key     :"tiny_id" ,
            Active  : false,
        }
        

        mockCacheResponse := &entity.UrlRegistry{
            URL     :"",
            Key     :"" ,
            Active  : false,
        }
        
        
        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Find",mock.Anything,mock.Anything).Return(mockResponse,nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Find",mock.Anything,mock.Anything).Return(mockCacheResponse,nil)

        useCase :=  NewRedirectUseCase(storageMock,cacheMock)
        resp,err := useCase.Redirect(c, usecaseReq)

        assert.Equal(t, errors.New("Can't redirect, the url expired"), err )     
        assert.Equal(t, "", resp )     
           
    })

    t.Run("Should return redirect_url when exist in redis and url is not active", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        mockResponse := &entity.UrlRegistry{
            URL     :"http://wwww.domain.com/original_long_url",
            Key     :"tiny_id" ,
            Active  : true,
        }
        
        
        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Find",mock.Anything,mock.Anything).Return(nil,nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Find",mock.Anything,mock.Anything).Return(mockResponse,nil)

        useCase :=  NewRedirectUseCase(storageMock,cacheMock)
        resp,err := useCase.Redirect(c, usecaseReq)

        assert.Equal(t,nil, err )     
        assert.Equal(t, "http://wwww.domain.com/original_long_url", resp )     
           
    })

    t.Run("Should return redirect_url when dosent's exist in redis but exist in database and url is active", func(t *testing.T) {

        body :=  `{}`
        e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/fakeendpoint", strings.NewReader(body))
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
        c := context.NewContext(ctx)

        mockResponse := &entity.UrlRegistry{
            URL     :"http://wwww.domain.com/original_long_url",
            Key     :"tiny_id" ,
            Active  : true,
        }
        

        mockCacheResponse := &entity.UrlRegistry{
            URL     :"",
            Key     :"" ,
            Active  : false,
        }
        
        storageMock:= new(mongoMock.MongoMock)
        storageMock.On("Find",mock.Anything,mock.Anything).Return(mockResponse,nil)

        cacheMock:= new(redisMock.RedisMock)
        cacheMock.On("Find",mock.Anything,mock.Anything).Return(mockCacheResponse,nil)

        useCase :=  NewRedirectUseCase(storageMock,cacheMock)
        resp,err := useCase.Redirect(c, usecaseReq)

        assert.Equal(t,nil, err )     
        assert.Equal(t, "http://wwww.domain.com/original_long_url", resp )     
           
    })
}
