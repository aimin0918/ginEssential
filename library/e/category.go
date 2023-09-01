package e

type category struct {
	ShowGoodsTypeGoods     string //指定菜品
	ShowGoodsTypeGoodsRule string //规则菜品
}

var Category = category{
	ShowGoodsTypeGoods:     "goods",
	ShowGoodsTypeGoodsRule: "goods_rule",
}
