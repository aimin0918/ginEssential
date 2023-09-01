package e

type cart struct {
	StatusInOpen       int //购物车开台状态
	StatusNotOpen      int //购物车为开台
	StatusPendingOrder int //购物车中有待支付订单
}

var Cart = cart{
	StatusInOpen:       1,
	StatusNotOpen:      0,
	StatusPendingOrder: 2,
}

type cartMessageAction struct {
	CartInfo        string //购物车信息Action
	CreateInOrder   string //创建订单中
	CartClosed      string //已关台
	CartOpenSuccess string // 开台成功
}

var CartMessageAction = cartMessageAction{
	CartInfo:        "cartInfo",
	CreateInOrder:   "createInOrder",
	CartClosed:      "cartClosed",
	CartOpenSuccess: "cartOpenSuccess",
}
