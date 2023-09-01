package log

import (
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

const MAX_TRIM_BYTES = 128

type dsRule interface {
	AffectFields() []string
	Transform(string) string
}

var dsRuleMap = map[string]dsRule{}

// implements dsRule ,and add more rank_define here
var dsRules = []dsRule{dsPasswordRule{}}

func init() {
	for _, r := range dsRules {
		for _, f := range r.AffectFields() {
			dsRuleMap[f] = r
		}
	}
}

type dsPasswordRule struct{}

func (p dsPasswordRule) AffectFields() []string {
	return []string{"password", "Password"}
}
func (p dsPasswordRule) Transform(raw string) string {
	return strings.Repeat("*", len(raw))
}

func DsAny(key string, val interface{}) zap.Field {
	return zap.Any(key, desensitization(val))
}

func desensitization(val interface{}) interface{} {
	if val == nil {
		return val
	}
	valDirectType := reflect.Indirect(reflect.ValueOf(val)).Type()
	if valDirectType.Kind() != reflect.Struct {
		return val
	}
	newVal := reflect.New(valDirectType).Interface()
	_ = copier.CopyWithOption(newVal, val, copier.Option{DeepCopy: true})
	findAndTransform(reflect.ValueOf(newVal))
	return newVal
}

func findAndTransform(val reflect.Value) {
	val = reflect.Indirect(val)
	switch val.Kind() {
	case reflect.Struct:
		valType := val.Type()
		for i := 0; i < val.NumField(); i++ {
			f := reflect.Indirect(val.Field(i))
			if !f.CanSet() {
				continue
			}
			switch f.Kind() {
			case reflect.Struct:
				findAndTransform(f)
			case reflect.Slice:
				transformSlice(f)
			case reflect.String:
				str := f.String()
				rule := dsRuleMap[valType.Field(i).Name]
				if rule != nil {
					str = rule.Transform(str)
					f.Set(reflect.ValueOf(str))
				}
			}
		}
	case reflect.Slice:
		transformSlice(val)
	}
}

func transformSlice(value reflect.Value) {
	elementKind := value.Type().Elem().Kind()
	if elementKind == reflect.Uint8 {
		// trim large []byte
		bs := value.Bytes()
		if len(bs) > MAX_TRIM_BYTES {
			bs = bs[0:MAX_TRIM_BYTES]
			value.Set(reflect.ValueOf(bs))
		}
		return
	}
	// see Kind definition
	// not process simple type: int, float etc..
	if elementKind <= reflect.Func || elementKind == reflect.String {
		return
	}
	for j := 0; j < value.Len(); j++ {
		findAndTransform(value.Index(j))
	}
}
