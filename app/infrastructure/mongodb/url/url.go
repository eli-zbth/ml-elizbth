package url

import (
	"context"
	"errors"
	"ml-elizabeth/app/domain/entity"
	"ml-elizabeth/app/domain/repository"
	"ml-elizabeth/app/shared/utils/constants"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"ml-elizabeth/app/infrastructure/mongodb/connection"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoOptionsRepository struct {
	conn   connection.MongoDbRepository

}

type MongoUrlRaw struct {
	ID           primitive.ObjectID `bson:"_id"`
	key         string
	Url      	string
	Active       bool
}

func NewMongoOptionsRepository(conn connection.MongoDbRepository) repository.StorageRepository {
	return &MongoOptionsRepository{conn: conn }
}

func (m *MongoOptionsRepository) Save(url string, id string, active bool) error {

	collection,_ := m.conn.GetConnection()
	urlRaw := &entity.UrlRegistry{
		URL: url,
		Key: id,
		Active: true,
	}
	_, err := collection.InsertOne(context.TODO(), urlRaw)
	if err != nil {
		isDuplicatedError := strings.Contains(string(err.Error()),constants.MongoDbDuplicateError)

		if isDuplicatedError{
			return errors.New(constants.DuplicateUrlError)
		}

		return err
	}
	return nil
}

func (m *MongoOptionsRepository) Update(filterkey string, filtervalue string, updateKey string,updatevalue string) (error ) {
	collection,_ := m.conn.GetConnection()
	
	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}

	filter := bson.M{filterkey: filtervalue}

	update := bson.M{
		"$set":   bson.M{updateKey: updatevalue},
	}

	updateResult := collection.FindOneAndUpdate(context.TODO(), filter,update, &updateOptions)
	
	if updateResult.Err() != nil {
		isDosentExistError := strings.Contains(updateResult.Err().Error(),constants.MongoDosentExistsError)

		if  isDosentExistError {
			return errors.New(constants.DosentExistsError)
		}
		return updateResult.Err() 
	}

	return nil
}


func (m *MongoOptionsRepository) Find(filterkey string, filtervalue string) ( *entity.UrlRegistry, error ) {
	collection,_ := m.conn.GetConnection()
	

	filter := bson.M{filterkey: filtervalue}

	var result MongoUrlRaw
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		isDosentExistError := strings.Contains(err.Error(),constants.MongoDosentExistsError)

		if  isDosentExistError {
			return nil, errors.New(constants.DosentExistsError)
		}
		return nil,err
	}

	return &entity.UrlRegistry{
		URL: result.Url,
		Key: result.key,
		Active: result.Active,
	} , nil
} 