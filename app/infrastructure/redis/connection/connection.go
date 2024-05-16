package connection

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)


type RedisOptionsRepository interface {
	GetConnection() *redis.Client
}


type RedisConnection struct {
	clientOptions *redis.Options
}

func NewRedisConnector() *RedisConnection {
    clientOptions := &redis.Options{
        Addr:viper.GetString("REDIS_URL"),
        // Password:viper.GetString("REDIS_PASSWORD"),
    }
    return &RedisConnection{
		clientOptions: clientOptions,
	}
}

func (m *RedisConnection) GetConnection() *redis.Client {
    return redis.NewClient(m.clientOptions)
}
