package dredis

import (
	"github.com/go-redis/redis"
	"holmes/conf"
	"log"
	"strconv"
	"time"
)

var (
	Redis *redis.Client
)

// init redis
func InitRedis() {
	timeout, err := strconv.Atoi(conf.Conf.Database.Redis.Timeout)
	if err != nil {
		log.Fatalln(err)
		return
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:         conf.Conf.Database.Redis.Host + ":" + conf.Conf.Database.Redis.Port,
		Password:     conf.Conf.Database.Redis.Auth, // password set
		DB:           conf.Conf.Database.Redis.Db,   // use default DB
		ReadTimeout:  time.Duration(timeout) * time.Millisecond,
		WriteTimeout: time.Duration(timeout) * time.Millisecond,
		DialTimeout:  time.Duration(timeout) * time.Millisecond,
	})
}
