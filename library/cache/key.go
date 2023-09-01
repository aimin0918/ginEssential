package cache

import (
	"fmt"
)

var CustomerKey = map[string]string{
	"wxCustomer":  addPreKey("wx:%s"),
	"aliCustomer": addPreKey("ali:%s"),
	"customer":    addPreKey("customer:%d"),
}

var RankRuleKey = map[string]string{
	"rankKey":              addPreKey("rankKey:id_%d"),
	"firstRankKey":         addPreKey("firstRankKey:company_id_%d"),
	"rankListKey":          addPreKey("rankListKey:company_id_%d"),
	"userOrderCount":       addPreKey("userOrderCount:%d"),
	"userOrderAmount":      addPreKey("userOrderAmount:%d:%d"),
	"pendingBackupBenefit": addPreKey("pendingBackupBenefit"),
}

func addPreKey(s string) string {
	return fmt.Sprintf("whgo_xjd:%s", s)
}

var CategorysKey = map[string]string{
	"categoryList":      addPreKey("categoryList"),
	"upperCategoryList": addPreKey("upperCategoryList"),
	"whaleCategoryList": addPreKey("whaleCategoryList"),
}

var GoodsKey = map[string]string{
	"categoryGoodsV3Key": addPreKey("categoryGoodsV3Key:categoryGoodsList_%d"),
	"backGoodsKey":       addPreKey("backGoodsKey:backGoods"),
	"quantityShopKey":    addPreKey("quantityShopKey:quantityKey_%d"),
	"guQingShopKey":      addPreKey("guQingShopKey:shop_%d"),
	"shopGoods":          addPreKey("shopGoods:shop_%d:channel_%d"),
	"goodsInfo":          addPreKey("goodsInfo"),
	"goodsDetail":        addPreKey("goodsDetail:shop_%d"),
	"shopGoodsPrice":     addPreKey("shop_goods_price:%d"),
	"goodsUnKey":         addPreKey("goodsUnKey:goodsUn:shop_%d"),
	"goodsRecommend":     addPreKey("goodsRecommend:shop_%d"),
	"recommend":          addPreKey("goodsRecommend:info"),
}

var OrderKey = map[string]string{
	"TcDineInPendingOrder":       addPreKey("tc_dinein_pending_order:%s"),
	"TcDineInShopPendingOrder":   addPreKey("tc_dinein_shop_pending_order:%s:%d"),
	"TcDeliveryPendingOrder":     addPreKey("tc_delivery_pending_order:%s"),
	"TcDeliveryShopPendingOrder": addPreKey("tc_delivery_shop_pending_order:%s:%d"),
	"TcReturnOrderReqID":         addPreKey("tc_return_order_req_id:%s"),
	"CartInfo":                   addPreKey("cartinfo:%s:%d"),
	"CartBenefitUsage":           addPreKey("benefit_usage:%s:%d"),
	"CartLifeCycle":              addPreKey("cartlifecycle"),
	"PendingOrderInfo":           addPreKey("pending-order:%s"),
	"VirtualOrder":               addPreKey("virtual_order:%s"),
	"CartInfoLock":               addPreKey("cartinfo:lock:%s:%d"),
	"DeliveryCode":               addPreKey("order:shop_%d:dineWay_%d:%s"),
	"OrderGoods":                 addPreKey("order_goods:orderNo_%s"),
	"ExpireWaitPay":              addPreKey("expire_wait_pay"),
	"ztTimeCount":                addPreKey("order:ztTimeCount_%d_%s"),
}

var ShopKey = map[string]string{
	"shopCityListKey": addPreKey("shopCityListKey:shopCityList"),
	"getShopById":     addPreKey("shop:%d"),
	"backendShopsKey": addPreKey("BackendShopsKey:BackendShops"),
	"shopIdsKey":      addPreKey("shopIdsKey:shopIds"),
	"tcMcIdsKey":      addPreKey("tcMcIdsKey:mcIds"),
	//优化版店铺列表
	"shopCityList": addPreKey("shopCityList:p_%s_c_%s"),
}

var DeliveryTimeKey = map[string]string{
	//系统自提时间配置
	"DeliveryTimeKey": addPreKey("DeliveryTimeKey:DeliveryTimes"),
	//第二版缓存key，预计要作废
	"DeliveryTimeListKey": addPreKey("DeliveryTimeListKey:DeliveryTimeList"),
	//第三版时间切片缓存key
	"DeliveryTimeShopKey": addPreKey("DeliveryTimeShopKey:DeliveryTimeList_%s_%d"),
}

var JobRecord = map[string]string{
	"RankJobMonth": addPreKey("RankJobMonth:%d"),
	"RankJobDay":   addPreKey("RankJobDay:%d"),
}

var PointJobRecord = map[string]string{
	"PointJobDay": addPreKey("PointJobDay:%s"),
}

var Renovation = map[string]string{
	"Renovation": addPreKey("Renovation:%s"),
}

var TcsGroup = map[string]string{
	"GroupInfo": addPreKey("tcs_group_info"),
}

var PopupWindow = map[string]string{
	"popup":       addPreKey("popup:%s_%d:%s:%s"),
	"PopupWindow": addPreKey("PopupWindow:%s_%d:%s"),
}

var Coupon = map[string]string{
	"CouponList": addPreKey("couponList:customer:%d"),
}

var Config = map[string]string{
	"config": addPreKey("config:%s"),
}

var IntegralGoods = map[string]string{
	"pageCache":   addPreKey("integralGoods:type_%d"),
	"goodsDetail": addPreKey("integralGoodsDetail:goodsId_%d"),
	"goodsStock":  addPreKey("integralGoodsStock"),
}

var UserAddress = map[string]string{
	"userAddress":     addPreKey("userAddress:customer_id_%d"),
	"userAddressList": addPreKey("userAddressList"),
}

var Postage = map[string]string{
	"postage": addPreKey("postage:id_%d"),
}

var PointsOrder = map[string]string{
	"detail": addPreKey("p_o_d:order_no:%s"),
}

var Activity = map[string]string{
	"activity": addPreKey("activity:id_%d"),
}

var SolarDay = map[string]string{
	"solar": addPreKey("solar-day:%s"), // 节气
}

var DeductionPointKey = map[string]string{
	"deduction_point": addPreKey("deduction_point:%s"),
	"birthday":        "birthday:%d", // 生日
	"member":          "member:%d",   // 会员
	"solar":           "solar:%d",    // 节气
}

var PointAutoRedisLockKey = map[string]string{
	"point_calculate_lock": addPreKey("point_calculate_lock:%d"),
}

var WechatAccessTokenKey = map[string]string{
	"wechat_office_account": addPreKey("access_token:wechat_office_account"), // 公众号token key
	"wechat_mini_program":   addPreKey("access_token:wechat_mini_program"),   // 小程序号token key
	"wechat_work":           addPreKey("access_token:wechat_work"),           // 企微token key
}
