package e

type pointsOrderState struct {
	WaitPay      int // 待支付
	WaitingPay   int // 支付中
	PaySuccess   int // 支付成功
	Shipped      int // 已发货
	Success      int // 订单完成
	Cancel       int // 订单取消
	ApplyRefund  int // 订单申请退单
	Refund       int // 订单退单
	RejectRefund int // 订单拒绝退单
}

var PointsOrderStatus = pointsOrderState{
	WaitPay:      10,
	WaitingPay:   20,
	PaySuccess:   30,
	Shipped:      40,
	Success:      80,
	Cancel:       -1,
	ApplyRefund:  -2,
	Refund:       -3,
	RejectRefund: -4,
}

var PointsOrderStatusValue = map[int]string{
	PointsOrderStatus.WaitPay:      "待支付",
	PointsOrderStatus.WaitingPay:   "支付中",
	PointsOrderStatus.PaySuccess:   "支付成功",
	PointsOrderStatus.Shipped:      "已发货",
	PointsOrderStatus.Success:      "订单完成",
	PointsOrderStatus.Cancel:       "订单取消",
	PointsOrderStatus.ApplyRefund:  "订单申请退单",
	PointsOrderStatus.Refund:       "订单退单",
	PointsOrderStatus.RejectRefund: "订单拒绝退单",
}

const PointOrderNoPrefix = "JFSC"
