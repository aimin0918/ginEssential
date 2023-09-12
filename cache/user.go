package cache

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"ginessential/library/cache"
//	redis "ginessential/library/redis/v8"
//	"ginessential/models"
//	"ginessential/utils"
//)
//
//type UserCache struct {
//
//}
//
//func (u *UserCache) GetUserOrderById (ctx context.Context, id int64) (user models.Users, err error) {
//	redisCil, _ := redis.GetRedis(ctx)
//	key := fmt.Sprintf(cache.User["user"], id)
//	cacheVal, err := redisCil.Get(ctx, key).Result()
//	if err != nil {
//		err = nil
//
//		// 缓存解析失败回表查询
//		user, err = models.GetUserDetail(ctx, id)
//		if err != nil {
//			return
//		}
//
//		cacheByte, _ := json.Marshal(user)
//		redisCil.Set(ctx, key, string(cacheByte), utils.Day)
//
//		return
//	}
//
//	_ = json.Unmarshal([]byte(cacheVal), &user)
//
//	return
//}
//
//func (u *UserCache) GetRootOrderById (ctx context.Context, id int64) (root models.Root, err error) {
//	redisCil, _ := redis.GetRedis(ctx)
//	key := fmt.Sprintf(cache.Root["root"], id)
//	cacheVal, err := redisCil.Get(ctx, key).Result()
//	if err != nil {
//		err = nil
//
//		// 缓存解析失败回表查询
//		root, err = models.GetRootDetail(ctx, id)
//		if err != nil {
//			return
//		}
//
//		cacheByte, _ := json.Marshal(root)
//		redisCil.Set(ctx, key, string(cacheByte), utils.Day)
//
//		return
//	}
//
//	_ = json.Unmarshal([]byte(cacheVal), &root)
//
//	return
//}
