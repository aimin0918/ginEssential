package e

type pointIsBalance struct {
	AddMore     int // 增加大于消耗
	ConsumeMore int // 增加小于消耗
	Equal       int // 持平
}

var PointIsBalance = pointIsBalance{
	AddMore:     0,
	ConsumeMore: 1,
	Equal:       2,
}
