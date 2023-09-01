package e

type rankRule struct {
	StateEnabled      int    //等级配置状态: 启用
	StateDisabled     int    //等级配置状态: 禁用
	OrderAmountKey    string //条件名称: 金额
	OrderTimeKey      string //条件名称: 频次
	OrderCountKey     string //条件名称: 数量
	CondKeyLimit      string //条件类型: 限制
	CondKeyNoLimit    string //条件类型: 不限制
	TimeTypeEveryDay  string //时间类型: 每日
	TimeTypeEveryWeek string //时间类型: 每周
}

var RankRuleDefine = rankRule{
	StateDisabled:     0,
	StateEnabled:      1,
	OrderAmountKey:    "orderAmount",
	OrderTimeKey:      "orderTime",
	OrderCountKey:     "orderCount",
	CondKeyLimit:      "limit",
	CondKeyNoLimit:    "noLimit",
	TimeTypeEveryDay:  "everyDay",
	TimeTypeEveryWeek: "onceWeek",
}
