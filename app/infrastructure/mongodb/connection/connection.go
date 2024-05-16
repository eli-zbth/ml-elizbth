package connection

import (
	"context"
	"ml-elizabeth/app/shared/utils/constants"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var collection *mongo.Collection

type MongoDbRepository interface {
	GetConnection() (*mongo.Collection, error)
}

type MongoDBConnection struct {
	clientOptions *options.ClientOptions
}

func NewMongoConnector() *MongoDBConnection {
    clientOptions := options.Client().ApplyURI(viper.GetString("MONGO_DB_URI"))
    return &MongoDBConnection{
		clientOptions: clientOptions,
	}
}

func (m *MongoDBConnection) GetConnection() (*mongo.Collection, error) {
	var err error
	if collection == nil {

        client, err := mongo.Connect(context.Background(), m.clientOptions)
		if err != nil {
			log.Errorf("error trying to open a connect to DB: %s", err)
			return nil, err
		} else {
            collection := client.Database(constants.MongodatabaseName).Collection(constants.MongocollectionName)
			if err != nil {
				log.Errorf("error trying to get collection for DB: %s", err)
                return nil, err
			}
            return collection,nil
		}
	}
	return collection, err
}


