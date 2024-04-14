package service

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"zaiyun.app/config"
)

var RedisClient *redis.Client

func init() {
	// 从配置文件或环境变量中获取Redis连接参数
	redisCfg := config.RedisConfig{
		Address:  "localhost:6379",
		Password: "ganto123",
		DB:       0,
	}

	// 创建Redis客户端选项
	opts := &redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	}

	// 初始化Redis客户端
	client := redis.NewClient(opts)

	// 连接Redis并检查连接状态
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	RedisClient = client
}

// GetRedisJWT 从Redis中获取与用户名关联的JWT
func GetRedisJWT(username string) (error, string) {
	key := fmt.Sprintf("jwt:%s", username)
	val, err := RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			err = nil // 用户名对应的JWT不存在，返回空字符串和nil错误
		}
		return err, ""
	}
	return nil, val
}

// SetRedisJWT 将JWT保存到Redis，并与用户名关联
func SetRedisJWT(jwtToken string, username string) error {
	key := fmt.Sprintf("jwt:%s", username)
	err := RedisClient.Set(context.Background(), key, jwtToken, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set JWT in Redis: %w", err)
	}
	return nil
}
