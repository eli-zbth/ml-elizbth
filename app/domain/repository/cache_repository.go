package repository

import "ml-elizabeth/app/domain/entity"


type CacheRepository interface {
    Save(urlRegistry *entity.UrlRegistry) error
    Find(key string)(*entity.UrlRegistry, error)
    UpdateKey(oldkey string, newKey string) error
    UpdateValue(key string, field string, value string) error
}

