package v8

import (
	"context"
	"fmt"
	"ginessential/library/log"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
	"time"
)

type PrefixHook struct {
	prefix           string
	skipReflectCheck bool
}

func (p *PrefixHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (p *PrefixHook) AfterProcess(_ context.Context, _ redis.Cmder) error {
	return nil
}

func (p *PrefixHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (p *PrefixHook) AfterProcessPipeline(_ context.Context, _ []redis.Cmder) error {
	return nil
}

type OpentracingHook struct {
	config *RedisConfig
	tracer opentracing.Tracer
}

const _tracingHookOpKey = "whale.redis.span"

func (o *OpentracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	if parentSp := opentracing.SpanFromContext(ctx); parentSp != nil {
		sp, _ := opentracing.StartSpanFromContextWithTracer(ctx, o.tracer, "redis."+cmd.Name())
		return context.WithValue(ctx, _tracingHookOpKey, sp), nil
	}
	return ctx, nil
}

func (o *OpentracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	if ctx.Value(_tracingHookOpKey) == nil {
		return nil
	}

	sp := ctx.Value(_tracingHookOpKey).(opentracing.Span)

	if sp == nil {
		return nil
	}
	sp.SetTag("component", "REDIS")
	sp.SetTag("addr", o.config.Addr)

	err := cmd.Err()
	if err != nil {
		sp.SetTag("error", true)
	}
	var fields []opentracinglog.Field
	fields = append(fields, opentracinglog.String("cmd", cmd.Name()))
	fields = append(fields, opentracinglog.String("args", fmt.Sprint(cmd.Args())))
	if err != nil {
		fields = append(fields, opentracinglog.String("err_msg", err.Error()))
	}
	sp.LogFields(fields...)
	sp.Finish()

	return nil
}

const _OpPipeline = "redis.pipeline"

func (o *OpentracingHook) BeforeProcessPipeline(ctx context.Context, _ []redis.Cmder) (context.Context, error) {
	if parentSp := opentracing.SpanFromContext(ctx); parentSp != nil {
		sp, _ := opentracing.StartSpanFromContextWithTracer(ctx, o.tracer, _OpPipeline)
		return context.WithValue(ctx, _tracingHookOpKey, sp), nil
	}
	return ctx, nil
}

func (o *OpentracingHook) AfterProcessPipeline(ctx context.Context, _ []redis.Cmder) error {
	sp := ctx.Value(_tracingHookOpKey).(opentracing.Span)

	if sp == nil {
		return nil
	}
	sp.SetTag("component", "REDIS")
	sp.SetTag("addr", o.config.Addr)
	sp.Finish()

	return nil
}

const _logHookTimeKey = "whale.redis.time"

type LogHook struct {
	CommonLog bool
	SlowLog   int
}

func (l *LogHook) BeforeProcess(ctx context.Context, _ redis.Cmder) (context.Context, error) {
	return context.WithValue(ctx, _logHookTimeKey, time.Now()), nil
}

func (l *LogHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	beforeTime := ctx.Value(_logHookTimeKey).(time.Time)

	if l.CommonLog {
		if cmd.Err() != nil {
			log.Warn(
				"error running redis command ", zap.Any("CMD:", cmd), zap.Error(cmd.Err()),
				zap.String("Took Time", time.Since(beforeTime).String()),
			)
		} else {
			log.Info(
				"running redis command", zap.Any("CMD:", cmd),
				zap.String("Took Time", time.Since(beforeTime).String()),
			)
		}
	}

	if l.SlowLog > 0 {
		slowLog(beforeTime, cmd, int64(l.SlowLog))
	}
	return nil
}

func (l *LogHook) BeforeProcessPipeline(ctx context.Context, _ []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (l *LogHook) AfterProcessPipeline(_ context.Context, _ []redis.Cmder) error {
	return nil
}

func slowLog(before time.Time, cmd redis.Cmder, threshold int64) {
	if time.Since(before).Milliseconds() >= threshold {
		if cmd.Name() == "brpop" {
			return
		}
		log.Warn(
			"redis_slow_log", zap.Int64("took_time", time.Since(before).Milliseconds()), zap.Any("redis_command", fmt.Sprintf("%s %s", cmd.FullName(), cmd.String())),
		)
	}
}
