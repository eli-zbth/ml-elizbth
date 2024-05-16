package redirect

import (

	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/shared/utils/context"

	"github.com/stretchr/testify/mock"
)

type RedirectUseCaseMock struct {
	mock.Mock
}

func (m *RedirectUseCaseMock) Redirect(ctx *context.Context, req *request.RedirectRequest) (string, error) {
    args := m.Called(ctx, req)
	return "",args.Error(0)
}
