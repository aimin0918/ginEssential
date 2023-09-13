package v8

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// ClusterClientWrapper wrapper ClusterClient to support batch cmd
// TODO @yueshan.zhang: calculate the hash key
type ClusterClientWrapper struct {
	*redis.ClusterClient
}

func (w *ClusterClientWrapper) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	pipe := w.Pipeline()
	var cmds []*redis.IntCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.Del(ctx, key))
	}

	_, err := pipe.Exec(ctx)
	return collectIntCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	pipe := w.Pipeline()
	var cmds []*redis.IntCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.Unlink(ctx, key))
	}

	_, err := pipe.Exec(ctx)
	return collectIntCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	pipe := w.Pipeline()
	var cmds []*redis.IntCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.Exists(ctx, key))
	}

	_, err := pipe.Exec(ctx)
	return collectIntCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	pipe := w.Pipeline()
	var cmds []*redis.IntCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.Touch(ctx, key))
	}

	_, err := pipe.Exec(ctx)
	return collectIntCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	pipe := w.Pipeline()
	var cmds []*redis.SliceCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.MGet(ctx, key))
	}

	_, err := pipe.Exec(ctx)
	return collectSliceCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	pipe := w.Pipeline()
	var cmds []*redis.StatusCmd

	for i := 0; i < len(values)-1; i += 2 {
		cmds = append(cmds, pipe.MSet(ctx, values[i], values[i+1]))
	}

	_, err := pipe.Exec(ctx)
	return collectStatusCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	pipe := w.Pipeline()
	var cmds []*redis.StringSliceCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.BLPop(ctx, timeout, key))
	}

	_, err := pipe.Exec(ctx)
	return collectStringSliceCmdResults(err, cmds, ctx)
}

func (w *ClusterClientWrapper) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	pipe := w.Pipeline()
	var cmds []*redis.StringSliceCmd

	for _, key := range keys {
		cmds = append(cmds, pipe.BRPop(ctx, timeout, key))
	}

	_, err := pipe.Exec(ctx)
	return collectStringSliceCmdResults(err, cmds, ctx)
}

func collectIntCmdResults(err error, cmds []*redis.IntCmd, ctx context.Context) *redis.IntCmd {
	resultCmd := redis.NewIntCmd(ctx)

	if err != nil {
		resultCmd.SetErr(err)
		return resultCmd
	}

	for _, cmd := range cmds {
		resultCmd.SetVal(resultCmd.Val() + cmd.Val())
	}

	return resultCmd
}

func collectSliceCmdResults(err error, cmds []*redis.SliceCmd, ctx context.Context) *redis.SliceCmd {
	resultCmd := redis.NewSliceCmd(ctx)

	if err != nil {
		resultCmd.SetErr(err)
		return resultCmd
	}

	var results []interface{}
	for _, cmd := range cmds {
		results = append(results, cmd.Val()...)
	}

	resultCmd.SetVal(results)
	return resultCmd
}

func collectStringSliceCmdResults(err error, cmds []*redis.StringSliceCmd, ctx context.Context) *redis.StringSliceCmd {
	resultCmd := redis.NewStringSliceCmd(ctx)

	if err != nil {
		resultCmd.SetErr(err)
		return resultCmd
	}

	var results []string
	for _, cmd := range cmds {
		results = append(results, cmd.Val()...)
	}

	resultCmd.SetVal(results)
	return resultCmd
}

func collectStatusCmdResults(err error, cmds []*redis.StatusCmd, ctx context.Context) *redis.StatusCmd {
	resultCmd := redis.NewStatusCmd(ctx)

	if err != nil {
		resultCmd.SetErr(err)
		return resultCmd
	}

	status := "OK"
	for _, cmd := range cmds {
		if cmd.Val() != "OK" {
			status = cmd.Val()
		}
	}

	resultCmd.SetVal(status)
	return resultCmd
}

func (w *ClusterClientWrapper) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	pipe := w.Pipeline()
	var cmds []*redis.IntCmd

	for _, geo := range geoLocation {
		cmds = append(cmds, pipe.GeoAdd(ctx, key, geo))
	}

	_, err := pipe.Exec(ctx)
	return collectIntCmdResults(err, cmds, ctx)
}
