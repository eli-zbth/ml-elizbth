package manage

import (
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/models/response"

	"github.com/stretchr/testify/mock"
)

type useCaseMock struct {
	mock.Mock
}

func (m *useCaseMock) UsecaseCreateUrl(req *request.CreateUrlRequest) (*response.CreateUrlResponse, error) {
	args := m.Called(req)

	if args.Get(0) != nil {
		return args.Get(0).(*response.CreateUrlResponse), nil
	}

	return nil, args.Error(1)
}
