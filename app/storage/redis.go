package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var (
	ctx = context.Background()
)

type redisDB struct {
	log        *log.Logger
	host       string
	expireTime time.Duration
}

func NewRedisDB(log *log.Logger, host string, expireTime time.Duration) *redisDB {
	return &redisDB{log, host, expireTime}
}

func (db *redisDB) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     db.host,
		Password: "", // no password
		DB:       0,  // default DB
	})
}

func (db *redisDB) Set(shortUrl string, originalUrl string) error {
	client := db.getClient()
	err := client.Set(ctx, shortUrl, originalUrl, db.expireTime).Err()
	if err != nil {
		db.log.Println("[ERROR] problem in setting key value to redis. ", err)
		return err
	}
	return nil
}

func (db *redisDB) Get(shortUrl string) (string, error) {
	client := db.getClient()
	val, err := client.Get(ctx, shortUrl).Result()
	if err != nil {
		db.log.Println("[ERROR] DB: GET url: ", shortUrl, err)
		return "", err
	}
	db.log.Println("[INFO] Short URL:", shortUrl, " Original URL:", val)
	return val, nil
}
