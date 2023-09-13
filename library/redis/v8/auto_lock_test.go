package v8

import (
	"context"
	"fmt"
	"ginessential/library/cache"
	"testing"
	"time"
)

func TestName(t *testing.T) {

}

func TestAutoRedisLock_Lock(t *testing.T) {
	client, _ := GetRedis(context.Background())
	for i := 1; i <= 100; i++ {
		go func(i int) {
			autoRedisLock := NewAutoRedisLock(context.Background(), client, fmt.Sprintf(cache.PointAutoRedisLockKey["point_calculate_lock"], 1000001), fmt.Sprintf("%d", 1000001))
			if err := autoRedisLock.Lock(); err == nil {
				defer func() {
					autoRedisLock.Unlock(false)
				}()
			}
			time.Sleep(3 * time.Second)
			fmt.Println("i lock success:", i)
		}(i)
	}
	time.Sleep(1 * time.Hour)
}
