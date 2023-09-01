package e

type shopGoods struct {
	TcStatusDelete   int8
	TcStatusNoDelete int8
}

// ShopGoods - shopGoods表常量定义
var ShopGoods = shopGoods{
	TcStatusNoDelete: 0,

	TcStatusDelete: 1,
}
