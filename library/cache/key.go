package cache

import (
	"fmt"
)

func addPreKey(s string) string {
	return fmt.Sprintf("ginEssential:%s", s)
}

var CategorysKey = map[string]string{
	"categoryList":      addPreKey("categoryList"),
	"upperCategoryList": addPreKey("upperCategoryList"),
	"whaleCategoryList": addPreKey("whaleCategoryList"),
}

var User = map[string]string{
	"user": addPreKey("user:id:%s"),
}

var Root = map[string]string{
	"root": addPreKey("root:id:%s"),
}

var Activity = map[string]string{
	"activity": addPreKey("activity:id_%d"),
}

var PointAutoRedisLockKey = map[string]string{
	"point_calculate_lock": addPreKey("point_calculate_lock:%d"),
}
