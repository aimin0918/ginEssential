package e

import "fmt"

func addPreKey(s string) string {
	return fmt.Sprintf("xjd%s", s)
}

type kafkaMessage struct {
	CustomerRegister string // 用户注册

	PointChange string // 积分变动
	RuleChange  string // 等级变动

	CreateOrder       string // 创建订单
	TcOrderExpired    string // 天财订单过期消息
	PushTcOrderFailed string // 天财订单推送失败
	OrderHeadList     string //订单头表
	OrderLineList     string //订单明细表

	SubscribePopupEvent    string //订阅消息事件
	SubscribeChangeEvent   string //删除订阅事件
	SubscribeSentEvent     string //订阅消息发送结果
	SubscribeSendToWxEvent string //订阅消息发送
}

var KafkaMessage = kafkaMessage{
	CustomerRegister: addPreKey("CustomerRegister"),

	PointChange: addPreKey("PointChange"),
	RuleChange:  addPreKey("RuleChange"),

	CreateOrder:            addPreKey("CreateOrder"),
	TcOrderExpired:         addPreKey("TcOrderExpired"),
	PushTcOrderFailed:      addPreKey("PushTcOrderFailed"),
	OrderHeadList:          addPreKey("OrderHeadList"),
	OrderLineList:          addPreKey("OrderLineList"),
	SubscribePopupEvent:    addPreKey("SubscribePopupEvent"),
	SubscribeChangeEvent:   addPreKey("SubscribeChangeEvent"),
	SubscribeSentEvent:     addPreKey("SubscribeSentEvent"),
	SubscribeSendToWxEvent: addPreKey("SubscribeSendToWxEvent"),
}
