package v8

import (
	"encoding/json"
	"os"
	"runtime"
	"sync"
)

type RedisConfig struct {
	Addr            string `json:"addr"`            // 单机模式: "ip:port", 集群模式: "ip1:port1,ip2:port2,ip3:port3"
	User            string `json:"user"`            // 默认空, redis 6.0.0开始支持
	Password        string `json:"password"`        // 默认空
	IsCluster       bool   `json:"isCluster"`       // 是否集群模式
	DialTimeoutSec  int    `json:"dialTimeoutSec"`  // 连接超时时间
	PoolSize        int    `json:"poolSize"`        // 连接池最大连接数
	MinIdleConns    int    `json:"minIdleConns"`    // 连接池最少连接数
	PoolTimeoutSec  int    `json:"poolTimeoutSec"`  // 从连接池获取连接超时时间
	IdleTimeoutSec  int    `json:"idleTimeoutSec"`  // 连接存活时间
	ReadTimeoutSec  int    `json:"readTimeoutSec"`  // 读超时时间
	WriteTimeoutSec int    `json:"writeTimeoutSec"` // 写超时时间
	MaxRetries      int    `json:"maxRetries"`      // 重试次数
	MinRetryBackoff int    `json:"minRetryBackoff"` // 最小重试间隔, 毫秒
	MaxRetryBackoff int    `json:"maxRetryBackoff"` // 最大重试间隔, 毫秒
	DB              int    `json:"db"`              // 集群模式无效
	Log             bool   `json:"log"`             // 是否开启调用log
	SlowLog         int    `json:"slowLog"`         // 慢查询阈值, 毫秒
}

var l sync.Mutex
var rdc *RedisConfig

// 单例
func LoadRedis() {
	if rds != nil {
		return
	}
	l.Lock()
	defer l.Unlock()
	if rds != nil {
		return
	}
	env := os.Getenv("env")

	if env == "" {
		err := os.Setenv("env", "dev")
		if err != nil {
			return
		}
		env = "dev"
	}
	config := &RedisConfig{}

	//path := "conf/" + env + "/redis.ini"
	//utils.LoadIni(path, "redis", config)
	// 默认值
	if config.DialTimeoutSec == 0 {
		config.DialTimeoutSec = 2
	}
	if config.PoolTimeoutSec == 0 {
		config.PoolTimeoutSec = 2
	}
	if config.IdleTimeoutSec == 0 {
		config.IdleTimeoutSec = 120
	}
	if config.ReadTimeoutSec == 0 {
		config.ReadTimeoutSec = 2
	}
	if config.WriteTimeoutSec == 0 {
		config.WriteTimeoutSec = 2
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
		config.MinRetryBackoff = 10
		config.MaxRetries = 100
	}
	// 慢日志
	//if config.SlowLog == 0 {
	//	config.SlowLog = 50
	//}

	// 获取全局配置copy
	m := map[string]interface{}{}
	str, _ := json.Marshal(GetGlobalConfig())
	_ = json.Unmarshal(str, &m)

	m1 := map[string]interface{}{}
	str1, _ := json.Marshal(config)
	_ = json.Unmarshal(str1, &m1)

	for k, v := range m1 {
		m[k] = v
	}

	bs, _ := json.Marshal(m)
	rdc = new(RedisConfig)
	_ = json.Unmarshal(bs, rdc)
}

var globalConfig RedisConfig
var globalConfigOnce sync.Once

var DefaultConfig = RedisConfig{
	DialTimeoutSec:  2,
	PoolSize:        10 * runtime.NumCPU(),
	MinIdleConns:    5,
	PoolTimeoutSec:  2,
	IdleTimeoutSec:  120,
	ReadTimeoutSec:  2,
	WriteTimeoutSec: 2,
	MinRetryBackoff: 2,
	MaxRetryBackoff: 10,
	DB:              0,
	Log:             false,
	SlowLog:         20,
}

func GetGlobalConfig() RedisConfig {
	globalConfigOnce.Do(func() {
		if globalConfig.DialTimeoutSec <= 0 {
			globalConfig.DialTimeoutSec = DefaultConfig.DialTimeoutSec
		}
		if globalConfig.PoolSize <= 0 {
			globalConfig.PoolSize = DefaultConfig.PoolSize
		}
		if globalConfig.PoolTimeoutSec <= 0 {
			globalConfig.PoolTimeoutSec = DefaultConfig.PoolTimeoutSec
		}
		if globalConfig.MinIdleConns <= 0 {
			globalConfig.MinIdleConns = DefaultConfig.MinIdleConns
		}
		if globalConfig.IdleTimeoutSec <= 0 {
			globalConfig.IdleTimeoutSec = DefaultConfig.IdleTimeoutSec
		}
		if globalConfig.ReadTimeoutSec <= 0 {
			globalConfig.ReadTimeoutSec = DefaultConfig.ReadTimeoutSec
		}
		if globalConfig.WriteTimeoutSec <= 0 {
			globalConfig.WriteTimeoutSec = DefaultConfig.WriteTimeoutSec
		}
		if globalConfig.MaxRetries > 0 {
			if globalConfig.MinRetryBackoff <= 0 {
				globalConfig.MinRetryBackoff = DefaultConfig.MinRetryBackoff
			}
			if globalConfig.MaxRetryBackoff <= 0 {
				globalConfig.MaxRetryBackoff = DefaultConfig.MaxRetryBackoff
			}
		}
		if globalConfig.SlowLog <= 0 {
			globalConfig.SlowLog = DefaultConfig.SlowLog
		}
	})
	return globalConfig
}
