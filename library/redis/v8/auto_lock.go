package v8

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	PubSubPrefix        = "{redis_lock}_"
	DefaultExpiration   = 10
	DefaultSpinInterval = 100
)

// AutoRedisLock 自动续期redis
type AutoRedisLock struct {
	ctx         context.Context
	key         string
	value       string
	redisClient Redis
	expiration  time.Duration
	cancelFunc  context.CancelFunc
}

func NewAutoRedisLock(ctx context.Context, redisClient Redis, key string, value string) *AutoRedisLock {
	return &AutoRedisLock{
		ctx:         ctx,
		key:         key,
		value:       value,
		redisClient: redisClient,
		expiration:  time.Duration(DefaultExpiration) * time.Second}
}

func NewAutoRedisLockWithExpireTime(ctx context.Context, redisClient Redis, key string, value string, expiration time.Duration) *AutoRedisLock {
	return &AutoRedisLock{
		ctx:         ctx,
		key:         key,
		value:       value,
		redisClient: redisClient,
		expiration:  expiration,
	}
}

// TryLock try get lock only once, if get the lock return true, else return false
func (lock *AutoRedisLock) TryLock() (bool, error) {
	success, err := lock.redisClient.SetNX(lock.ctx, lock.key, lock.value, lock.expiration).Result()
	if err != nil {
		return false, err
	}
	if success {
		ctx, cancelFunc := context.WithCancel(context.Background())
		lock.cancelFunc = cancelFunc
		lock.renew(ctx)
	}
	return success, nil
}

// Lock blocked until get lock
func (lock *AutoRedisLock) Lock() error {
	for {
		success, err := lock.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}

// Unlock release the lock
func (lock *AutoRedisLock) Unlock(isLua bool) (err error) {
	var res interface{}
	defer lock.cancelFunc()
	if isLua {
		script := redis.NewScript(fmt.Sprintf(
			`if redis.call("get", KEYS[1]) == "%s" then return redis.call("del", KEYS[1]) else return 0 end`,
			lock.value))
		runCmd := script.Run(lock.ctx, lock.redisClient, []string{lock.key})
		res, err = runCmd.Result()
	} else {
		res, err = lock.redisClient.Del(lock.ctx, lock.key).Result()
	}
	if err != nil {
		return err
	}
	if tmp, ok := res.(int64); ok {
		if tmp == 1 {
			return nil
		}
	}
	err = fmt.Errorf("unlock script fail: %s", lock.key)
	return err
}

// LockWithTimeout blocked until get lock or timeout
func (lock *AutoRedisLock) LockWithTimeout(d time.Duration) error {
	timeNow := time.Now()
	for {
		success, err := lock.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
		deltaTime := d - time.Since(timeNow)
		if !success {
			err := lock.subscribeLockWithTimeout(deltaTime)
			if err != nil {
				return err
			}
		}
	}
}

func (lock *AutoRedisLock) SpinLock(times int) error {
	for i := 0; i < times; i++ {
		success, err := lock.TryLock()
		if err != nil {
			return err
		}
		if success {
			return nil
		}
		time.Sleep(time.Millisecond * DefaultSpinInterval)
	}
	return fmt.Errorf("max spin times reached")
}

// subscribeLock blocked until lock is released
func (lock *AutoRedisLock) subscribeLock() error {
	pubSub := lock.redisClient.Subscribe(lock.ctx, getPubSubTopic(lock.key))
	_, err := pubSub.Receive(lock.ctx)
	if err != nil {
		return err
	}
	<-pubSub.Channel()
	return nil
}

// subscribeLock blocked until lock is released or timeout
func (lock *AutoRedisLock) subscribeLockWithTimeout(d time.Duration) error {
	timeNow := time.Now()
	pubSub := lock.redisClient.Subscribe(lock.ctx, getPubSubTopic(lock.key))
	_, err := pubSub.ReceiveTimeout(lock.ctx, d)
	if err != nil {
		return err
	}
	deltaTime := time.Since(timeNow) - d
	select {
	case <-pubSub.Channel():
		return nil
	case <-time.After(deltaTime):
		return fmt.Errorf("timeout")
	}
}

// publishLock publish a message about lock is released
func (lock *AutoRedisLock) publishLock() error {
	err := lock.redisClient.Publish(lock.ctx, getPubSubTopic(lock.key), "release lock").Err()
	if err != nil {
		return err
	}
	return nil
}

// renew renew the expiration of lock, and can be canceled when call Unlock
func (lock *AutoRedisLock) renew(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(lock.expiration / 3)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				_, _ = lock.redisClient.Expire(lock.ctx, lock.key, lock.expiration).Result()
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

// getPubSubTopic key -> PubSubPrefix + key
func getPubSubTopic(key string) string {
	return PubSubPrefix + key
}
