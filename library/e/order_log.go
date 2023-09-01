package e

type orderLog struct {
	EventCreateOrder      int8 //事件: 创建订单
	EventRetrunOrder      int8 //事件: 整单退
	EventPartRetrunOrder  int8 //事件: 部分退款
	EventPayedOrder       int8 //事件: 支付成功
	EventCancelOrder      int8 //事件: 取消订单
	EventUpdateDeskNo     int8 //事件: 更新桌号
	EventTcStateUpdate    int8 //事件: 自提状态更新
	EventTcOrderAccept    int8 //事件: 天财已接收订单
	EventPushOrderToTC    int8 //事件: 推送订单到天财
	EventRePushOrderToTC  int8 //事件: 重新推送订单到天财
	EventPushOrderToTCRes int8 //事件: 推送订单到天财结果
}

var OrderLog = orderLog{
	EventCreateOrder:      1,
	EventRetrunOrder:      2,
	EventPartRetrunOrder:  3,
	EventPayedOrder:       4,
	EventCancelOrder:      5,
	EventUpdateDeskNo:     6,
	EventTcStateUpdate:    7,
	EventTcOrderAccept:    8,
	EventPushOrderToTC:    9,
	EventRePushOrderToTC:  10,
	EventPushOrderToTCRes: 11,
}
