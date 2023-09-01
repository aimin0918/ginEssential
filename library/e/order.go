package e

import "time"

const (
	OrderPayExpireDuration = 15 * time.Minute
)

type orderState struct {
	WaitPay    int // 待支付
	WaitingPay int // 支付中
	PaySuccess int // 支付成功
	InQueue    int //排队中

	MadeIn int // 制作中
	// Seated   int // 已落座
	TakeFood int // 可取餐

	Success    int // 订单完成
	Cancel     int // 订单取消
	Refund     int // 订单退单
	AutoRefund int // 订单自动退单
}

var OrderStatus = orderState{
	WaitPay:    10,
	WaitingPay: 20,
	PaySuccess: 30,
	InQueue:    40,
	// Seated:     45,
	MadeIn:     50,
	TakeFood:   60,
	Success:    80,
	Cancel:     -1,
	Refund:     -2,
	AutoRefund: -3,
}

type dineWay struct {
	EatIn    int // 堂食
	TakeSelf int // 自提
	InLine   int // 排队点餐
}

var DineWay = dineWay{
	EatIn:    1,
	TakeSelf: 2,
	InLine:   3,
}

type payType struct {
	WxPayType      int    //微信支付方式
	ApiPayType     int    //支付宝支付方式
	CloudPayType   int    //云闪付支付方式
	TcWxPayType    string //天财微信
	TcAliPayType   string //天财支付宝
	TcWxpayTypeId  int    //天财微信支付id
	TcAlipayTypeId int    //天财支付宝支付id
}

var PayType = payType{
	WxPayType:      1,
	ApiPayType:     2,
	CloudPayType:   3,
	TcWxPayType:    "6",
	TcAliPayType:   "8",
	TcWxpayTypeId:  6,
	TcAlipayTypeId: 8,
}

var PayTypeName = map[int]string{
	PayType.WxPayType:    "微信",
	PayType.ApiPayType:   "支付宝",
	PayType.CloudPayType: "云闪付",
}

// tcOrderState - 天财订单状态
type tcOrderState struct {
	DineInAcceptOrder int8 //堂食:门店POS接收到订单,2
	DineInCreateOrder int8 //堂食:门店POS下单成功,3
	DineInCancelOrder int8 //堂食:门店POS取消下单,4
	DineInCloseDesk   int8 //堂食:关台,5

	DeliveryAcceptOrder   int8 //自提:商户已接单状态1
	DeliveryFailOrer      int8 //自提:商户接单失败2
	DeliveryCancel        int8 //自提:商户取消3
	DeliveryUpToSend      int8 //自提: 外送起送4
	DeliveryDone          int8 //自提: 外送送达5
	DeliveryUnAccepted    int8 //自提: 商家未接单6
	DeliveryMadeIn        int8 //自提: 制作中9
	DeliveryPendingTakeUp int8 //自提:待取餐12
	DeliveryTakedUp       int8 //自提: 已取餐13
}

var TcOrderState = tcOrderState{
	DineInAcceptOrder: 1, //堂食:门店POS接收到订单,2
	DineInCreateOrder: 2, //堂食:门店POS下单成功,3
	DineInCancelOrder: 3, //堂食:门店POS取消下单,4
	DineInCloseDesk:   4, //堂食:关台,5

	DeliveryAcceptOrder:   6,  //自提:商户已接单状态1
	DeliveryFailOrer:      7,  //自提:商户接单失败2
	DeliveryCancel:        8,  //自提:商户取消3
	DeliveryUpToSend:      9,  //自提: 外送起送4
	DeliveryDone:          10, //自提: 外送送达5
	DeliveryUnAccepted:    11, //自提: 商家未接单6
	DeliveryMadeIn:        12, //自提: 制作中9
	DeliveryPendingTakeUp: 13, //自提:待取餐12
	DeliveryTakedUp:       14, //自提: 已取餐13
}

type pointDeductionType struct {
	PointDeductionDefault int
	PointDeductionYes     int
}

var PointDeductionType = pointDeductionType{
	PointDeductionDefault: 0,
	PointDeductionYes:     1,
}
