package log

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/xid"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc/metadata"
)

const traceIdKey = "trace_id"

func SetTraceId(ctx context.Context) context.Context {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if _, ok := span.Context().(jaeger.SpanContext); ok {
			return ctx
		}
	}
	inMd, inTraceId := getIncomingTraceId(ctx)
	outMd, outTraceId := getOutgoingTraceId(ctx)
	if inTraceId == "" {
		traceId, _ := ctx.Value(traceIdKey).(string)
		if traceId == "" {
			traceId = xid.New().String()
		}
		ctx = setIncomingTraceId(ctx, inMd, traceId)
		ctx = setOutgoingTraceId(ctx, outMd, traceId)
	} else if outTraceId != inTraceId {
		ctx = setOutgoingTraceId(ctx, outMd, inTraceId)
	}
	return ctx
}

func getIncomingTraceId(ctx context.Context) (metadata.MD, string) {
	md, _ := metadata.FromIncomingContext(ctx)
	traces := md[traceIdKey]
	if len(traces) == 0 {
		return md, ""
	}
	return md, traces[len(traces)-1]
}

func getOutgoingTraceId(ctx context.Context) (metadata.MD, string) {
	md, _ := metadata.FromOutgoingContext(ctx)
	traces := md[traceIdKey]
	if len(traces) == 0 {
		return md, ""
	}
	return md, traces[len(traces)-1]
}

func setIncomingTraceId(ctx context.Context, md metadata.MD, traceId string) context.Context {
	if md == nil {
		md = metadata.MD{}
	}
	md[traceIdKey] = []string{traceId}
	return metadata.NewIncomingContext(ctx, md)
}

func setOutgoingTraceId(ctx context.Context, md metadata.MD, traceId string) context.Context {
	if md == nil {
		md = metadata.MD{}
	}
	md[traceIdKey] = []string{traceId}
	return metadata.NewOutgoingContext(ctx, md)
}

func GetTraceId(ctx context.Context) string {
	span := opentracing.SpanFromContext(ctx)
	if span != nil {
		if span, ok := span.Context().(jaeger.SpanContext); ok {
			return span.TraceID().String()
		}
	}
	_, traceId := getIncomingTraceId(ctx)
	if traceId != "" {
		return traceId
	}
	traceId, _ = ctx.Value(traceIdKey).(string)
	return traceId
}
