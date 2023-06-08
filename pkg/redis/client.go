//Package redis
/*
@Title: client.go
@Description
@Author: kkw 2023/5/16 16:31
*/
package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.kkw.top/gamesTimeLine/internal/config"
	"time"
)

type RdClient struct {
	Client *redis.Client
}

var KwClient *RdClient

func NewClient() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Get("redis.host"),
		Password: config.Get("redis.pass"),         // 如果有密码，则设置密码
		DB:       config.Config.GetInt("redis.db"), // 使用默认数据库
	})
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis 链接失败:%s", err))
	}
	var client *RdClient = &RdClient{}
	client.Client = redisClient
	KwClient = client
}

func (c *RdClient) Get(key string) (string, error) {
	return c.Client.Get(c.Client.Context(), key).Result()
}

func (c *RdClient) Set(key string, val interface{}, expiration time.Duration) error {
	return c.Client.Set(c.Client.Context(), key, val, expiration).Err()
}

func (c *RdClient) Exists(key string) (int64, error) {
	return c.Client.Exists(c.Client.Context(), key).Result()

}
