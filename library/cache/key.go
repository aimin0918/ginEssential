package cache

import (
	"fmt"
)

func addPreKey(s string) string {
	return fmt.Sprintf("ginEssential:%s", s)
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

var CategorysKey = map[string]string{
	"categoryList":      addPreKey("categoryList"),
	"upperCategoryList": addPreKey("upperCategoryList"),
	"whaleCategoryList": addPreKey("whaleCategoryList"),
}

// %d是int类型, %s是string类型

var User = map[string]string{
	"user": addPreKey("user:id_%d"),
}

var Root = map[string]string{
	"root": addPreKey("root:id_%d"),
}

var RootList = map[string]string{
	"rootList": addPreKey("rootList:id_%d"),
}

var SolarDay = map[string]string{
	"solar": addPreKey("solar-day:%s"), // 节气
}
