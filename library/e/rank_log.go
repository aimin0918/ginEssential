package e

type rankLog struct {
	ChangeTypePromotion         int8 //等级变动类型: 升级
	ChangeTypeDemotion          int8 //等级变动类型: 降级
	ChangeReasonPromotionAmount int8 //等级升级原因：金额
	ChangeReasonPromotionRate   int8 //等级升级原因：频次
	ChangeReasonDemotionAmount  int8 //等级降级原因：金额
	ChangeReasonDemotionRate    int8 //等级降级原因: 频次
	ChangeReasonDemotionReturn  int8 //等级降级原因: 退单
	ChangeReasonAdmin           int8 //等级变化原因: 手动
	ChangeReasonRegister        int8 //等级变化原因: 首次注册
}

var RankLog = rankLog{
	ChangeTypePromotion:         1,
	ChangeTypeDemotion:          2,
	ChangeReasonPromotionAmount: 1,
	ChangeReasonPromotionRate:   2,
	ChangeReasonDemotionAmount:  3,
	ChangeReasonDemotionReturn:  4,
	ChangeReasonDemotionRate:    5,
	ChangeReasonAdmin:           6,
	ChangeReasonRegister:        7,
}
