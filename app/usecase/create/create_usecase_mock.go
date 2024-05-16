package create

import (

	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/models/response"
	"ml-elizabeth/app/shared/utils/context"
	"github.com/stretchr/testify/mock"
)

type CreateUseCaseMock struct {
	mock.Mock
}

func (m *CreateUseCaseMock) CreateUrl(ctx *context.Context, req *request.CreateUrlRequest) (*response.CreateUrlResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) != nil {
		return args.Get(0).(*response.CreateUrlResponse), nil
	}
	return nil, args.Error(1)
}
