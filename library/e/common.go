package e

type pointRateType struct {
	PointRateCustomer int64 // 会员类型
	PointRateSolar    int64 // 节气类型
	PointRateBirthday int64 // 生日
}

var PointRateType = pointRateType{
	PointRateCustomer: 1,
	PointRateSolar:    2,
	PointRateBirthday: 3,
}
