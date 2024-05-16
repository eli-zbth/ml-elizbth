package repository

import "ml-elizabeth/app/domain/entity"


type StorageRepository interface {
	Save(url string, id string, active bool) error
    Update(filterkey string, filtervalue string, updateKey string,updatevalue string) error
    Find(filterkey string, filtervalue string) (*entity.UrlRegistry, error)
}
