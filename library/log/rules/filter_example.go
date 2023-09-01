package rules

import (
	"github.com/forestgiant/sliceutil"
	"go.uber.org/zap"
	"strings"
)

// example, 或者详细请参考filter_password.go的实现
func init() {
	//registerLogFilter(filterRuleExample, handlerExample)
}

// 自定义过滤条件 原型为 FilterRule
func filterRuleExample(field *zap.Field) []string {

	keyStringVal := strings.ToLower(field.Key) + "#example never never never hit rules#"
	if sliceutil.Contains([]string{"order", "ordernum"}, keyStringVal) {
		return []string{keyStringVal}
	}

	// 返回 nil 代表不需要过滤
	return nil
}

// 自定义过滤器 原型为 FilterHandler
func handlerExample(field *zap.Field, needFilterKeys []string) *zap.Field {
	// 参考 defaultHandler 的实现
	return nil
}
