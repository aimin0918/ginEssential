package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_Not_Vip_User:              "未注册会员",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",

	ERROR_CHECK_RANKS_SORT:    "当前等级已存在",
	ERROR_CHECK_RANKS_Title:   "当前等级名称已存在",
	ERROR_CHECK_RANKS_Amount:  "当前等级金额条件不满足条件",
	ERROR_CHECK_RANKS_Count:   "当前等级频次不满足条件",
	ERROR_POINT_BENEFIT_EXITS: "积分权益只能创建一个，目前已被创建！",

	//支付订单状态错误
	ERROR_ORDER_PAY_STATE_FAIL: "订单状态发生改变，请稍后再试",
	ERROR_ORDER_PAY_DATA_FAIL:  "获取订单支付出错",

	GET_ORDER_FAIl:              "获取订单失败",
	GET_ORDER_NO_FAIl:           "获取订单号失败",
	CREATE_ORDER_FAIl:           "创建订单失败",
	ORDER_STOCK_NOT_ENOUGN:      "库存不足",
	CART_UPDATED:                "订单有变动请重新支付",
	CART_EXPIRED:                "已超时请重新加购",
	ORDER_CREATED:               "订单已下单",
	ORDER_ZTTIME_INVALID:        "自提时间无效",
	ORDER_ZTTIME_EXPIRED:        "取餐时间已过，请重新下单！",
	ERROR_ORDER_PAY_AMOUNT_FAIL: "支付金额与订单金额不一致",
	BUSINESS_STATUS_FAIL:        "店铺已打烊",
	COUPON_NOT_EXIST:            "优惠券不存在",
	CUSTOMER_NOTFIND:            "该用户不存在",
	ORDER_AMOUNT_ZERO:           "订单金额不可为0,请继续加购",
	SHOP_NOT_OPEN:               "店铺未营业",
	ORDER_PAY_FAIL:              "支付失败",
	ORDER_POINT_DEDUCTION_FAIL:  "订单积分抵扣出错，请重新下单！",
	//验签失败
	SIGNFAILED: "验签失败",

	SUBSCRIBE:   "已订阅",
	UNSUBSCRIBE: "未订阅",

	ERROR_GOOD_NOT_FIND:  "商品信息查询不到",
	SHOP_CONFIG_NOT_OPEN: "店铺配置未打开",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
