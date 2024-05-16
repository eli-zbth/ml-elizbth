package url

import (
	"github.com/stretchr/testify/mock"

	"ml-elizabeth/app/domain/entity"
)

type RedisMock struct {
	mock.Mock
}


func (m *RedisMock) Save(urlRegistry *entity.UrlRegistry) error{
    args := m.Called(urlRegistry)
    return args.Error(0)
}

func (m *RedisMock) Find(key string)(*entity.UrlRegistry, error){
    args := m.Called(key)
    if args.Get(0) != nil {
		return args.Get(0).(*entity.UrlRegistry), nil
	}
    return nil, args.Error(1)
}

func (m *RedisMock) UpdateKey(oldkey string, newKey string) error {
    args := m.Called(oldkey, newKey)
	return args.Error(0)
}

func (m *RedisMock) UpdateValue(key string, field string, value string) error {
    args := m.Called(key, field,value)
	return args.Error(0)
}