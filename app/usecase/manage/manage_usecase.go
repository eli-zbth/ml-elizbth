package manage

import (
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/models/response"
	"ml-elizabeth/app/domain/repository"
	"ml-elizabeth/app/shared/utils/logger"
)

type ManageUseCase interface {
	UsecaseCreateUrl(*request.CreateUrlRequest) (*response.CreateUrlResponse, error)
}

type manageUseCase struct {
	manageRepository repository.RandomRepository
}

func NewManageUseCase(manageRepository repository.RandomRepository) ManageUseCase {
	return &manageUseCase{
		manageRepository: manageRepository,
	}
}

func (m *manageUseCase) UsecaseCreateUrl(createUrlRequest *request.CreateUrlRequest) (*response.CreateUrlResponse, error) {
	log := logger.New()
	log.WithFields(logger.Fields{
		"layer":  "manage_usecase",
		"method": "CreateUrl",
	})

	log.Info("Hasta aca vamos bien")

	return &response.CreateUrlResponse{
		Test: "yeeeeei",
	}, nil
}
