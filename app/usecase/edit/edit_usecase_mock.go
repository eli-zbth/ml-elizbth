package edit

import (

	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/shared/utils/context"

	"github.com/stretchr/testify/mock"
)

type EditUseCaseMock struct {
	mock.Mock
}


func (m *EditUseCaseMock) EditShortUrl(ctx *context.Context, req *request.EditShortUrlRequest) ( error) {
    args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *EditUseCaseMock) EditRedirectUrl(ctx *context.Context, req *request.EditRedirectURLRequest) ( error) {
    args := m.Called(ctx, req)
	return args.Error(0)
}


func (m *EditUseCaseMock) EditUrlStatus(ctx *context.Context, req *request.EditUrlStatusRequest) ( error) {
    args := m.Called(ctx, req)
	return args.Error(0)
}


