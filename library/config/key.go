package config

type KeyType string

const (
	DeductionIsOpen KeyType = "deduction_is_open" // 积分抵扣key
	PointSetting    KeyType = "points_setting"    // 积分配置key
)
