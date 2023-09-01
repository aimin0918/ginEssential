package rules

import (
	"fmt"
	"github.com/forestgiant/sliceutil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
	"strings"
)

type FilterHandler func(field *zap.Field, needFilterKeys []string) *zap.Field
type FilterRule func(field *zap.Field) []string

type LogFilter struct {
	handler FilterHandler
	rule    FilterRule
}

var filterRegistry = make([]LogFilter, 0)

func Filter(field *zap.Field) *zap.Field {
	fieldParam := field
	for _, filter := range getAllFilters() {
		keysToBeFiltered := filter.rule(field)
		if keysToBeFiltered != nil && len(keysToBeFiltered) > 0 {
			if filter.handler == nil {
				filter.handler = filterHandlerDefault
			}
			fieldParam = filter.handler(fieldParam, keysToBeFiltered)
		}
	}

	return fieldParam
}

func registerLogFilter(rule FilterRule, handler FilterHandler) {
	filter := LogFilter{
		handler: handler,
		rule:    rule,
	}
	filterRegistry = append(filterRegistry, filter)
}

func filterHandlerDefault(field *zap.Field, needFilterKeys []string) *zap.Field {
	retField := &zap.Field{
		Key:       field.Key,
		Type:      field.Type,
		Integer:   field.Integer,
		String:    field.String,
		Interface: nil,
	}

	if field.Type == zapcore.StringType {
		filterZapStringFieldVal(retField, "********", needFilterKeys...)
	} else if field.Type == zapcore.ReflectType {
		if field.Interface == nil {
			return field
		}
		dupInterfaceT := reflect.TypeOf(field.Interface).Elem()
		dupInterface := reflect.New(dupInterfaceT)
		dupInterface.Elem().Set(reflect.ValueOf(field.Interface).Elem())
		retField.Interface = dupInterface.Interface()
		filterZapInterfaceFieldVal(retField, "********", needFilterKeys...)
	} else if field.Type == zapcore.StringerType {
		if field.Interface == nil {
			return field
		}
		//reflect.TypeOf(field.Interface)
		dupInterfaceT := reflect.TypeOf(field.Interface)
		dupInterface := reflect.New(dupInterfaceT).Elem()
		dupInterface.Set(reflect.ValueOf(field.Interface))
		retField.Interface = dupInterface.Interface()
		filterZapStringerFieldVal(retField, "********", needFilterKeys...)
	} else {
		// doing nothing
		return field
	}

	return retField
}

func filterZapInterfaceFieldVal(field *zap.Field, intoVal string, keys ...string) {
	if field.Type == zapcore.ReflectType {

		t := reflect.TypeOf(field.Interface).Elem()
		val := reflect.ValueOf(field.Interface).Elem()

		if t.Kind() == reflect.Struct {
			for i := 0; i < t.NumField(); i++ {
				if keywordExistInKeyList(t.Field(i).Name, keys...) {
					val.Field(i).SetString(intoVal)
				}
			}

		} else if t.Kind() == reflect.Map {
			mapIter := val.MapRange()
			for mapIter.Next() {
				mapKey := mapIter.Key().String()
				if keywordExistInKeyList(mapKey, keys...) {
					val.SetMapIndex(reflect.ValueOf(mapKey), reflect.ValueOf(intoVal))
				}

			}
		}
	}
}

func filterZapStringerFieldVal(field *zap.Field, intoVal string, keys ...string) {
	val, ok := field.Interface.(fmt.Stringer)
	if ok {
		strVal := val.String()
		stringFields := strings.Fields(strVal)
		newStringFields := make([]string, 0)
		if len(stringFields) > 0 {
			for _, eachField := range stringFields {
				kvPair := strings.Split(eachField, ":")
				if len(kvPair) == 2 {
					if keywordExistInKeyList(kvPair[0], keys...) {
						// 执行过滤
						kvPair[1] = intoVal
						newStringFields = append(newStringFields, strings.Join(kvPair, ":"))
						continue
					}
				}
				newStringFields = append(newStringFields, eachField)
			}
		}
		field.Type = zapcore.StringType
		field.String = strings.Join(newStringFields, " ")
		field.Interface = nil
	}
}

func filterZapStringFieldVal(field *zap.Field, intoVal string, keys ...string) {
	if keywordExistInKeyList(field.Key, keys...) {
		field.String = intoVal
	}
}

func keywordExistInKeyList(keyword string, keyList ...string) bool {
	lowerCaseKey := strings.ToLower(keyword)
	if sliceutil.Contains(keyList, lowerCaseKey) {
		return true
	}
	return false
}

func getAllFilters() []LogFilter {
	return filterRegistry
}
