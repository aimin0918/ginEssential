package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"ginessential/library/cache"
	redis "ginessential/library/redis/v8"
	"ginessential/models"
	"ginessential/utils"
)

type RootCache struct {
}

func (u *RootCache) GetRootOrderById(ctx context.Context, Id int64) (root models.Root, err error) {
	redisCil, _ := redis.GetRedis(ctx)
	key := fmt.Sprintf(cache.Root["root"], Id)
	cacheVal, err := redisCil.Get(ctx, key).Result()
	if err != nil {
		err = nil

		// 缓存解析失败回表查询
		root, err = models.GetRootById(ctx, Id)
		if err != nil {
			return
		}

		cacheByte, _ := json.Marshal(root)
		redisCil.Set(ctx, key, string(cacheByte), utils.Day)

		return
	}

	_ = json.Unmarshal([]byte(cacheVal), &root)

	return
}
