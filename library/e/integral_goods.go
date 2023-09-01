package e

type integralGoodsType struct {
	VirtualGoodsType int
	RealGoodsType    int
}

var IntegralGoodsType = integralGoodsType{
	VirtualGoodsType: 1, //虚拟商品
	RealGoodsType:    2, //实际商品
}

var IntegralGoodsTypeValue = map[int]string{
	IntegralGoodsType.VirtualGoodsType: "虚拟商品",
	IntegralGoodsType.RealGoodsType:    "实际商品",
}

type integralGoodsState struct {
	OffLine int
	OnLine  int
}

var IntegralGoodsState = integralGoodsState{
	OnLine:  1,
	OffLine: -1,
}

type integralGoodsIsAuto struct {
	NoAuto int
	Auto   int
}

var IntegralGoodsIsAuto = integralGoodsIsAuto{
	NoAuto: 0,
	Auto:   1,
}
