package e

type couponLog struct {
	EventMeiTuanCouponCharge    int8 //事件: 美团核销
	EventMeiTuanCouponChargeRes int8 //事件: 美团核销结果
	EventMeiTuanCouponCancel    int8 //事件: 美团核销结果
	EventMeiTuanCouponCancelRes int8 //事件: 美团核销结果
}

var CouponLog = couponLog{
	EventMeiTuanCouponCharge:    1,
	EventMeiTuanCouponChargeRes: 2,
	EventMeiTuanCouponCancel:    3,
	EventMeiTuanCouponCancelRes: 4,
}
