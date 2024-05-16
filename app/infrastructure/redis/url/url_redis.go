package url

import (
	"context"

	"ml-elizabeth/app/domain/entity"
	"ml-elizabeth/app/domain/repository"
	"ml-elizabeth/app/infrastructure/redis/connection"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type RedisRepository struct {
	conn   connection.RedisOptionsRepository
}

type RedisMap struct {
	Url     string     
    Active   string   
}

func NewRedisRepository(conn connection.RedisOptionsRepository) repository.CacheRepository {
	return &RedisRepository{conn: conn }
}

func (r *RedisRepository) Save(urlRegistry *entity.UrlRegistry) error {
	client := r.conn.GetConnection()


    strconv.FormatBool(urlRegistry.Active)
 
    active := strconv.FormatBool(urlRegistry.Active)
    err := client.HSet(context.Background(), urlRegistry.Key,"url", urlRegistry.URL, "active",active ).Err()
    if err != nil {
        log.Errorf("Error trying to save data in Redis: %s",err)
        return err
    }

    return nil
}

func toBool(value string) bool {
    b, err := strconv.ParseBool(value)
    if err != nil {
        return false
    }
    return b
}

func (r *RedisRepository) Find(key string) (*entity.UrlRegistry, error) {
	client := r.conn.GetConnection()

    redisResult, err := client.HGetAll(context.Background(), key).Result()
    if err != nil {
        log.Errorf("Error trying to get data from Redis: %s",err)
        return nil,err
    }
    urlRegistry := &entity.UrlRegistry{
        URL:  redisResult["url"],
        Key: key ,
        Active: toBool(redisResult["active"]),
    }
    return urlRegistry, nil
}

func (r *RedisRepository) UpdateKey(oldKey string,newKey string)  error {
	client := r.conn.GetConnection()
    ctx:= context.Background()

    redisResult, err := client.HGetAll(ctx, oldKey).Result()
    if err != nil {
        log.Errorf("Error trying to get data from Redis: %s",err)
        return err
    }

    err = client.Del(ctx, oldKey).Err()
    if err != nil {
        log.Errorf("Error trying to delete get data from Redis: %s",err)
    }

    err = client.HSet(ctx,newKey,"url", redisResult["url"], "active", redisResult["active"] ).Err()
    if err != nil {
        log.Errorf("Error trying to update key data in Redis: %s",err)
        return err
    }
 
    return nil
}

func (r *RedisRepository) UpdateValue(key string, field string, value string)  error {
	client := r.conn.GetConnection()

    err := client.HSet(context.Background(),key,field,value).Err()
    if err != nil {
        log.Errorf("Error trying to update Key data in Redis: %s",err)
        return err
    }
    return nil
}
