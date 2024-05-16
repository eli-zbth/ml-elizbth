package edit

import (
	"ml-elizabeth/app/domain/models/request"
	"ml-elizabeth/app/domain/repository"
	"ml-elizabeth/app/shared/utils/constants"
	"ml-elizabeth/app/shared/utils/parse"
	"ml-elizabeth/app/shared/utils/context"
	"strconv"
)


type Edit interface {
	EditShortUrl(*context.Context,*request.EditShortUrlRequest) (error)
	EditRedirectUrl(*context.Context, *request.EditRedirectURLRequest) (error)
	EditUrlStatus(*context.Context, *request.EditUrlStatusRequest)(error)
}

type EditUseCase struct {
	StorageRepository repository.StorageRepository
	CacheRepository repository.CacheRepository
}

func NewEditUseCase(storage repository.StorageRepository , cache repository.CacheRepository) Edit {
	return &EditUseCase{
		StorageRepository: storage,
		CacheRepository: cache,
	}
}


func (e *EditUseCase) EditShortUrl(ctx *context.Context, req *request.EditShortUrlRequest) ( error) {
	log := ctx.SetupLogger("edit_usecase","edit_short_url_request")
	
	shortId := parse.DeleteDomain(req.ShortUrl)

	log.Info("Save!")
	err := e.StorageRepository.Update(constants.UpdateShortUrlField,shortId, constants.UpdateShortUrlField, req.NewUrl )
	if err != nil {
		log.Errorf("Error when edit the url in database: %s",err)
			return err
	}

	err = e.CacheRepository.UpdateKey(shortId,  req.NewUrl)
	if err != nil {
		log.Errorf("Error when edit the url in Cache: %s",err)
			return err
	}
	return nil
}


func (e *EditUseCase) EditRedirectUrl(ctx *context.Context, req *request.EditRedirectURLRequest) ( error) {
	log := ctx.SetupLogger("edit_usecase","edit_redirect_url")

	shortId := parse.DeleteDomain(req.ShortUrl)
	err := e.StorageRepository.Update(constants.UpdateShortUrlField,shortId, constants.UpdateStatusField, req.RedirectUrl )
	if err != nil {
		log.Errorf("Error when edit the url in database: %s",err)
		return err
	}

	err = e.CacheRepository.UpdateValue(shortId,  "url",req.RedirectUrl)
	if err != nil {
		log.Errorf("Error when edit the redirect url in cache: %s",err)
			return err
	}
	
	return nil
}


func (e *EditUseCase) EditUrlStatus(ctx *context.Context, req *request.EditUrlStatusRequest) ( error) {
	log := ctx.SetupLogger("edit_usecase","edit_url_status")
	active := strconv.FormatBool(*req.IsActive)
	shortId := parse.DeleteDomain(req.ShortUrl)
	err := e.StorageRepository.Update(constants.UpdateShortUrlField,shortId, constants.UpdateStatusField, active )
	if err != nil {
		log.Errorf("Error when edit the url status in database: %s",err)
		return err
	}

	err = e.CacheRepository.UpdateValue(shortId,  "active", active)
	if err != nil {
		log.Errorf("Error when edit the redirect url in cache: %s",err)
			return err
	}

	return nil
}


