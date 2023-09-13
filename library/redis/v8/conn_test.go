package v8

import (
	"context"
	"fmt"
	"ginessential/library/cache"
	redisV8 "github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	_ = os.Setenv("WHALE_JAEGER_TRACE_ENABLE", "true")
	_ = os.Setenv("WHALE_JAEGER_AGENT_ADDR", "10.168.0.27:32179")
)

func TestSetKey(t *testing.T) {
	cli, err := GetRedisClient()
	if err != nil {
		t.Errorf("get redis client failed: %v", err)
	}
	result, err := cli.Set(context.Background(), "test_key", "test_value", 60*time.Second).Result()
	fmt.Printf("result: %v, err: %v\n", result, err)

}

func TestPrefix(t *testing.T) {
	prefix := "prefix-"
	key := "test-key"
	client, _ := GetRedisClient()
	prefixedClient, _ := GetPrefixedRedisClient(prefix)

	client.Set(context.Background(), prefix+key, "value", 0)
	defer client.Del(context.Background(), prefix+key)

	result, _ := prefixedClient.Exists(context.Background(), key).Result()

	if result != 1 {
		t.Errorf("test prefix failed")
	}
}

type PrintHook struct {
	Content string
}

func (p *PrintHook) BeforeProcess(ctx context.Context, _ redisV8.Cmder) (context.Context, error) {
	fmt.Println("before: " + p.Content)
	return ctx, nil
}

func (p *PrintHook) AfterProcess(_ context.Context, _ redisV8.Cmder) error {
	fmt.Println("after: " + p.Content)
	return nil
}

func (p *PrintHook) BeforeProcessPipeline(ctx context.Context, _ []redisV8.Cmder) (context.Context, error) {
	fmt.Println("before: " + p.Content)
	return ctx, nil
}

func (p *PrintHook) AfterProcessPipeline(_ context.Context, _ []redisV8.Cmder) error {
	fmt.Println("after: " + p.Content)
	return nil
}

func TestZAdd(t *testing.T) {
	redisCli, _ := GetRedis(context.Background())
	key := cache.OrderKey["ExpireWaitPay"]
	var orderNo string
	for i := 0; i < 100; i++ {
		orderNo = "16631420571010" + strconv.Itoa(i)
		timeStamp := time.Now().Unix()
		redisCli.ZAdd(context.Background(), key, &redisV8.Z{
			Score:  float64(timeStamp),
			Member: orderNo,
		})
	}
}

func TestHMGet(t *testing.T) {
	redisCli, _ := GetRedis(context.Background())
	var field []string
	redisCli.HMSet(context.Background(), "test1", map[string]string{
		"goods_id_3": "3",
	})

	for i := 0; i < 10; i++ {
		field = append(field, fmt.Sprintf("goods_id_%d", i))
	}
	res, err := redisCli.HMGet(context.Background(), "test1", field...).Result()

	for key, val := range res {
		if val == nil {

		}
		t.Log(field[key])
		t.Log(val)
	}

	t.Log("err", err)
	t.Log("res", res)

}
