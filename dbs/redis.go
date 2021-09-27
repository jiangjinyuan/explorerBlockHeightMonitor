package dbs

import (
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// RedisDB redis client
var RedisDB *redis.Client

// InitRedisDB 初始化 redis 数据库
func InitRedisDB(address, password string) error {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:            address,
		Password:        password,
		DialTimeout:     3 * time.Second,
		ReadTimeout:     3 * time.Second,
		WriteTimeout:    3 * time.Second,
		PoolTimeout:     5 * time.Second,
		MaxRetries:      3,
		MaxRetryBackoff: 1 * time.Minute,
	})
	status := RedisDB.Ping()
	if status.Err() != nil {
		log.Error(status.Err())
		return status.Err()
	}
	return nil
}

func InitRedisDBCluster(address []string, password string) error {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:           address,
		Password:        password,
		DialTimeout:     3 * time.Second,
		ReadTimeout:     3 * time.Second,
		WriteTimeout:    3 * time.Second,
		PoolTimeout:     5 * time.Second,
		MaxRetries:      3,
		MaxRetryBackoff: 1 * time.Minute,
		RouteRandomly:   true,
		ReadOnly:        true,
	})
	status := client.Ping()
	if status.Err() != nil {
		log.Error(status.Err())
		return status.Err()
	}
	return nil
}
