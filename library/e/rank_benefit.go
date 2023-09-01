package e

type rankBenefit struct {
	ValidDateTypeForever string //生效日期类型: 永久
	ValidDateTypeDay     string //生效日期类型: 每日
	ValidDateTypeWeek    string //生效日期类型: 每周
	ValidDateTypeMonth   string //生效日期类型: 每月
	ValidDateTypeYear    string //生效日期类型: 每年
	ValidDateTypeRange   string //生效日期类型: 日期范围

	BenefitTypePoint    string //权益类型: 积分
	BenefitTypeDiscount string //权益类型: 折扣

	StateEnabled  int //状态: 开启
	StateDisabled int //状态: 禁用

	PointValidTypeAfter        string // 滚动过期
	PointValidTypeLimitCurrent string //积分时间生效类型: 当前
	PointValidTypeForever      string //积分时间生效类型: 永久

	PointValidDimensionMonth string //积分时间生效类型: 每月
	PointValidDimensionYear  string //积分时间生效类型: 每年

	DiscountTimeDimensionDay     string //折扣时间维度类型: 每日
	DiscountTimeDimensionWeek    string //折扣时间维度类型: 每周
	DiscountTimeDimensionTotal   string //折扣时间维度类型: 总计
	DiscountTimeDimensionNoLimit string //折扣时间维度类型: 不限制

	DiscountDishesDimensionDay     string //折扣菜品维度限制类型: 每日
	DiscountDishesDimensionWeek    string //折扣菜品维度限制类型: 每周
	DiscountDishesDimensionTotal   string //折扣菜品维度限制类型: 总计
	DiscountDishesDimensionNoLimit string //折扣菜品维度限制类型: 不限制
	DiscountDishesDimensionOrder   string //折扣菜品维度限制类型: 每笔订单

	PointBenefitLimitTypeLimit    string // 积分权益限制类型: 限制
	PointBenefitLimitTypeNotLimit string // 积分权益限制类型: 不限制

	RuleCondKeyGoodsCategory string //规则条件名称: 菜品分类
	RuleCondKeyGoods         string //规则条件名称: 菜品
	RuleCondKeyTotalPrice    string //规则条件名称: 总金额

	RuleSymbolContain      string //规则条件类型: 包含
	RuleSymbolNotContain   string //规则条件类型: 不包含
	RuleSymbolGreater      string //规则条件类型: 大于
	RuleSymbolGreaterEqual string //规则条件类型: 大于等于
}

var RankBenefit = rankBenefit{
	ValidDateTypeForever: "forever",
	ValidDateTypeRange:   "range",
	ValidDateTypeDay:     "day",
	ValidDateTypeWeek:    "week",
	ValidDateTypeMonth:   "month",
	ValidDateTypeYear:    "year",

	BenefitTypePoint:    "point",
	BenefitTypeDiscount: "discount",

	StateEnabled:  1,
	StateDisabled: 2,

	PointValidTypeAfter:        "after",
	PointValidTypeLimitCurrent: "limit_current",
	PointValidTypeForever:      "forever",

	PointValidDimensionMonth: "month",
	PointValidDimensionYear:  "year",

	DiscountTimeDimensionDay:     "day",
	DiscountTimeDimensionWeek:    "week",
	DiscountTimeDimensionTotal:   "total",
	DiscountTimeDimensionNoLimit: "not_limit",

	DiscountDishesDimensionDay:     "day_dishes",
	DiscountDishesDimensionWeek:    "week_dishes",
	DiscountDishesDimensionTotal:   "total_dishes",
	DiscountDishesDimensionNoLimit: "not_limit_dishes",
	DiscountDishesDimensionOrder:   "order_dishes",

	PointBenefitLimitTypeLimit:    "limit",
	PointBenefitLimitTypeNotLimit: "not_limit",

	RuleCondKeyGoodsCategory: "goods_category",
	RuleCondKeyGoods:         "goods",
	RuleCondKeyTotalPrice:    "total_price",

	RuleSymbolContain:      "contain",
	RuleSymbolNotContain:   "not_contain",
	RuleSymbolGreater:      "greater",
	RuleSymbolGreaterEqual: "greaterAndEqual",
}
