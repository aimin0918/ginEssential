package e

// pointDetail -
type pointDetail struct {
	EventAdd    int //加积分
	EventReduce int //减积分

	StateNormal  int //积分状态正常
	StateExpired int //积分状态已过期
	StateUsed    int //积分已被使用

	SourceTypeOrder     int //关联类型是订单
	SourceTypeGoods     int //关联类型是商品
	SourceTypeReturnID  int //关联类型是退款流水号(部分退)
	SourceTypeTc        int //关联类型是天财
	SourceTypeAdmin     int //关联类型是后台
	SourceTypeIntegral  int //关联类型是积分商城
	SourceTypeActivity  int //关联类型是会员任务活动
	SourceTypeCustomer  int //关联类型是用户整体
	SourceTypeDeduction int //关联类型是积分当前花

}

var PointDetail = pointDetail{
	EventAdd:    1,
	EventReduce: 2,

	StateNormal:  0, //积分状态正常
	StateExpired: 1, //积分状态已过期
	StateUsed:    2, //积分已被使用

	SourceTypeOrder:     1,
	SourceTypeGoods:     2,
	SourceTypeReturnID:  3,
	SourceTypeTc:        4,
	SourceTypeAdmin:     5,
	SourceTypeIntegral:  6,
	SourceTypeActivity:  7,
	SourceTypeCustomer:  8,
	SourceTypeDeduction: 9,
}

var PointDetailSourceTypeMap = map[int]string{
	PointDetail.SourceTypeOrder:     "订单积分",
	PointDetail.SourceTypeReturnID:  "订单部分退款",
	PointDetail.SourceTypeAdmin:     "后台积分操作",
	PointDetail.SourceTypeIntegral:  "积分商城",
	PointDetail.SourceTypeActivity:  "会员活动",
	PointDetail.SourceTypeCustomer:  "用户整体积分",
	PointDetail.SourceTypeDeduction: "积分当钱花",
}
