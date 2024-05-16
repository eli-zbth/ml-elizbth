package redirect

import (
	"errors"
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/repository"
	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/context"
)


type Redirect interface {
	Redirect(*context.Context,*request.RedirectRequest) (string, error)
}

type RedirectUseCase struct {
	StorageRepository repository.StorageRepository
	CacheRepository repository.CacheRepository
}

func NewRedirectUseCase(StorageRepository repository.StorageRepository, cache repository.CacheRepository) Redirect {
	return &RedirectUseCase{
		StorageRepository: StorageRepository,
		CacheRepository: cache,
	}
}

func (r *RedirectUseCase) Redirect(ctx *context.Context,req *request.RedirectRequest) ( string, error) {
	log := ctx.SetupLogger("redirect_usecase","redirect_request")


	res,err := r.CacheRepository.Find(req.Id )
	
	if err != nil || res.URL==""{
		log.Infof("URL donsen't exists in Cache, continue with database: %s",req.Id)
		
		res,err = r.StorageRepository.Find(constants.UpdateShortUrlField,req.Id )
		if err != nil {
			log.Errorf("Error when lookup the URL: %s",err)
			return "",err
		}
	}

	if  res.Active {
		return res.URL	,nil
	}else{
		return "",errors.New("Can't redirect, the url expired")
	}
	
}