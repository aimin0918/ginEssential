package v8

import (
	"context"
	"errors"
	"ginessential/library/log"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"time"
)

// RedisLock 非自动续期
type RedisLock interface {
	Lock(ctx context.Context, key string, expires time.Duration) (UnLocker, error)
}

type UnLocker interface {
	Unlock(ctx context.Context) error
}

type lockImpl struct {
	client redis.UniversalClient
}

type unLockerImpl struct {
	key    string
	token  string
	client redis.UniversalClient
}

func NewRedisLock(client redis.UniversalClient) RedisLock {
	return &lockImpl{
		client: client,
	}
}

func (l *lockImpl) Lock(ctx context.Context, key string, expires time.Duration) (UnLocker, error) {
	token := uuid.NewV4().String()
	boolRes := l.client.SetNX(ctx, key, token, expires)
	if !boolRes.Val() {
		err := boolRes.Err()
		if err == nil {
			err = errors.New("try lock failed")
		}
		log.InfoWithCtx(ctx, "lock failed", zap.Error(err))
		return nil, err
	}
	return &unLockerImpl{
		key:    key,
		token:  token,
		client: l.client,
	}, nil
}

// see https://redis.io/commands/set
const unlockScript = `if redis.call("get",KEYS[1]) == ARGV[1]
then
    return redis.call("del",KEYS[1])
else
    return 0
end`

func (u *unLockerImpl) Unlock(ctx context.Context) error {
	cmd := u.client.Eval(ctx, unlockScript, []string{u.key}, u.token)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}
