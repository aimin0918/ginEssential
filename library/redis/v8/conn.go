package v8

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"sync"
	"time"
)

var redisLock sync.Mutex
var rds redis.UniversalClient

type Redis interface {
	redis.UniversalClient
	RedisLock
}

type redisImpl struct {
	redis.UniversalClient
	RedisLock
}

func GetRedis(ctx context.Context) (Redis, error) {
	client, err := GetRedisClient()
	if err != nil {
		return nil, err
	}

	return &redisImpl{
		UniversalClient: client,
		RedisLock:       NewRedisLock(client),
	}, nil
}

// GetPrefixedRedis get prefixed Redis, it will add a prefix to the keys in most operations
func GetPrefixedRedis(prefix string) (Redis, error) {
	client, err := GetPrefixedRedisClient(prefix)
	if err != nil {
		return nil, err
	}
	return &redisImpl{
		UniversalClient: client,
		RedisLock:       NewRedisLock(client),
	}, nil
}

func GetRedisClient() (redis.UniversalClient, error) {
	if rds != nil {
		return rds, nil
	}
	return GetPrefixedRedisClient("")
}

// GetPrefixedRedisClient get prefixed RedisClient, it will add a prefix to the keys in most operations
func GetPrefixedRedisClient(prefix string) (redis.UniversalClient, error) {
	redisLock.Lock()
	defer redisLock.Unlock()

	LoadRedis()

	client := NewRedisClient(rdc)
	applyHooks(client, rdc, prefix)
	rds = client

	return rds, nil
}

func applyHooks(client redis.UniversalClient, config *RedisConfig, prefix string) {
	// prefix hook must process first
	if prefix != "" {
		client.AddHook(&PrefixHook{prefix: prefix})
	}

	if config.Log || config.SlowLog > 0 {
		client.AddHook(&LogHook{
			CommonLog: config.Log,
			SlowLog:   config.SlowLog,
		})
	}
}

func NewRedisClient(config *RedisConfig) redis.UniversalClient {
	if config.IsCluster {
		return createClusterClient(config)
	} else {
		return createClient(config)
	}
}

func createClient(redisConfig *RedisConfig) redis.UniversalClient {
	return redis.NewClient(
		&redis.Options{
			Network:      "tcp",
			Addr:         redisConfig.Addr,
			Password:     redisConfig.Password,
			DialTimeout:  time.Duration(redisConfig.DialTimeoutSec) * time.Second,  // 连接超时时间
			MinIdleConns: redisConfig.MinIdleConns,                                 // 连接池，至少存活连接数
			PoolSize:     redisConfig.PoolSize,                                     // 连接池，最大连接数
			PoolTimeout:  time.Duration(redisConfig.PoolTimeoutSec) * time.Second,  // 从连接池，获取连接的超时时间，假设所有的 conn 都在忙
			IdleTimeout:  time.Duration(redisConfig.IdleTimeoutSec) * time.Second,  // 连接空闲时间，超过这个时间，连接会被关闭，重新 new
			ReadTimeout:  time.Duration(redisConfig.ReadTimeoutSec) * time.Second,  // 读超时时间
			WriteTimeout: time.Duration(redisConfig.WriteTimeoutSec) * time.Second, // 写超时时间

			// redis 命令执行失败，重试次数，默认 0
			// 如果命令返回 error:redis: nil 的错误，不会重试
			// write、read timeout 会重试
			// conn，连接池里 conn 无效，会重试
			MaxRetries:      redisConfig.MaxRetries,
			MinRetryBackoff: time.Duration(redisConfig.MinRetryBackoff) * time.Millisecond, // 重试最小间隔
			MaxRetryBackoff: time.Duration(redisConfig.MaxRetryBackoff) * time.Millisecond, // 重试最大间隔
			DB:              redisConfig.DB,
		},
	)
}

func createClusterClient(redisConfig *RedisConfig) redis.UniversalClient {
	client := redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs:        strings.Split(redisConfig.Addr, ","),
			Password:     redisConfig.Password,
			DialTimeout:  time.Duration(redisConfig.DialTimeoutSec) * time.Second,  // 连接超时时间
			MinIdleConns: redisConfig.MinIdleConns,                                 // 连接池，至少存活连接数
			PoolSize:     redisConfig.PoolSize,                                     // 连接池，最大连接数
			PoolTimeout:  time.Duration(redisConfig.PoolTimeoutSec) * time.Second,  // 从连接池，获取连接的超时时间，假设所有的 conn 都在忙
			IdleTimeout:  time.Duration(redisConfig.IdleTimeoutSec) * time.Second,  // 连接空闲时间，超过这个时间，连接会被关闭，重新 new
			ReadTimeout:  time.Duration(redisConfig.ReadTimeoutSec) * time.Second,  // 读超时时间
			WriteTimeout: time.Duration(redisConfig.WriteTimeoutSec) * time.Second, // 写超时时间
			MaxRedirects: 8,
			// redis 命令执行失败，重试次数，默认 0
			// 如果命令返回 error:redis: nil 的错误，不会重试
			// write、read timeout 会重试
			// conn，连接池里 conn 无效，会重试
			MaxRetries:      redisConfig.MaxRetries,
			MinRetryBackoff: time.Duration(redisConfig.MinRetryBackoff) * time.Millisecond, // 重试最小间隔
			MaxRetryBackoff: time.Duration(redisConfig.MaxRetryBackoff) * time.Millisecond, // 重试最大间隔
		},
	)

	return &ClusterClientWrapper{
		ClusterClient: client,
	}
}
