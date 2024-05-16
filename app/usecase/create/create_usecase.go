package create

import (
	"ml-elizabeth/app/domain/entity"
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/models/response"
	"ml-elizabeth/app/domain/repository"
	"ml-elizabeth/app/shared/utils/context"
	"ml-elizabeth/app/shared/utils/parse"

	"github.com/spf13/viper"
)


type Create interface {
	CreateUrl(*context.Context , *request.CreateUrlRequest) (*response.CreateUrlResponse, error)
}
type CreateUseCase struct {
	StorageRepository repository.StorageRepository
	CacheRepository repository.CacheRepository
}	

func NewCreateUsecase(storage repository.StorageRepository , cache repository.CacheRepository) Create {
	return &CreateUseCase{
		StorageRepository: storage,
		CacheRepository: cache,
	}
}

func (c *CreateUseCase) CreateUrl(ctx *context.Context, createUrlRequest *request.CreateUrlRequest) (*response.CreateUrlResponse, error) {
	log := ctx.SetupLogger("create_usecase","create_url_usecase")
	

	id:=createUrlRequest.CustomId
	if id == "" {
		id,_ = parse.GenerateId()
	}
	
	err := c.StorageRepository.Save(createUrlRequest.Url,id,true)
	if err != nil {
		log.Errorf("Error when save the url in database: %s",err)
		return nil,err
	}


	registredUrl:= &entity.UrlRegistry{
		URL: createUrlRequest.Url,
		Key: id,
		Active: true,
	}
	
	err = c.CacheRepository.Save(registredUrl)
	if err != nil {
		log.Errorf("Error when save the url in database: %s",err)
		return nil,err
	}

	shortDomain:= viper.GetString("SHORT_URL_DOMAIN")
	return &response.CreateUrlResponse{
		ShortUrl: shortDomain+id,
	}, nil
}
