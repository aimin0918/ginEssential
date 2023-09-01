package e

type rankBenefitUsage struct {
	SourceTypeOrder int //优惠记录类型：订单
	SourceTypeGoods int //优惠记录类型：商品
}

var RankBenefitUsage = rankBenefitUsage{
	SourceTypeOrder: 1,
	SourceTypeGoods: 2,
}
