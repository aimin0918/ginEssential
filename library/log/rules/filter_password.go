package rules

import (
	"github.com/forestgiant/sliceutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

func init() {
	// handler传入nil 代表使用defaultHandler
	registerLogFilter(filterPassword, nil)
}

// 自定义过滤条件 原型为 FilterRule
func filterPassword(field *zap.Field) []string {
	fieldsToBeFiltered := make([]string, 0)
	if field.Type == zapcore.StringType {
		keyStringVal := strings.ToLower(field.Key)
		if sliceutil.Contains([]string{"pass", "password"}, keyStringVal) {
			// 返回需要过滤的字段 (对象是string 就为field.Key, struct或者map的话 就是field和mapKey)
			fieldsToBeFiltered = append(fieldsToBeFiltered, keyStringVal)
		}
	}
	if field.Type == zapcore.ReflectType {
		// reflectType包含了map和struct类型, 返回的strings就表示需要过滤的内部字段或者key
		fieldsToBeFiltered = append(fieldsToBeFiltered, "pass", "password")
	}
	if field.Type == zapcore.StringerType {
		fieldsToBeFiltered = append(fieldsToBeFiltered, "pass", "password")
	}

	// 返回 nil 代表无字段需要过滤
	return fieldsToBeFiltered
}
