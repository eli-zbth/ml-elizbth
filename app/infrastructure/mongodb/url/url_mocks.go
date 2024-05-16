package url

import (
	"github.com/stretchr/testify/mock"

	"ml-elizabeth/app/domain/entity"
)

type MongoMock struct {
	mock.Mock
}

func (m *MongoMock) Save(url string, id string, active bool) error  {
	args := m.Called(url, id, active)
	return args.Error(0)
}

func (m *MongoMock) Update(filterkey string, filtervalue string, updateKey string,updatevalue string) error   {
	args := m.Called(filterkey, filtervalue, updateKey,updatevalue)
	return args.Error(0)
}

func (m *MongoMock) Find(filterkey string, filtervalue string) ( *entity.UrlRegistry, error )  {
	args := m.Called(filterkey, filtervalue)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.UrlRegistry), nil
	}
	return nil, args.Error(1)
}