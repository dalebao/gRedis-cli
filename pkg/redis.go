package pkg

import (
	"errors"
	"github.com/go-redis/redis"
	"strconv"
)

var Client *redis.Client

type RedisConfig struct {
	Addr     string
	Port     string
	Password string
	DB       string
}

func (redisConfig *RedisConfig) Dial() error {
	db := 0
	if redisConfig.DB != "" {
		db, _ = strconv.Atoi(redisConfig.DB)
	}
	Client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password, // no password set
		DB:       db,                   // use default DB
	})
	_, err := Client.Ping().Result()
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
