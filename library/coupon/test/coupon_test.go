package test

import (
	"context"
	"encoding/json"
	"fmt"
	"ginessential/library/coupon/channel"
	"testing"
)

func TestMeiTuan(t *testing.T) {
	c, err := channel.NewCouponLib("meituan")
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()

	content, err := c.QueryByMobile(ctx, channel.UserCouponQueryReq{
		VendorShopId: "33352",
		Mobile:       "18656478575",
	})

	var resp []channel.UserCouponQueryResp
	contentJson, _ := json.Marshal(content)
	_ = json.Unmarshal(contentJson, &resp)

	t.Log(fmt.Sprintf("%+v\n", resp))
}

func TestMeiTuanQuery(t *testing.T) {
	c, err := channel.NewCouponLib("meituan")
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()

	content, err := c.CouponStatusQuery(ctx, channel.CouponStatusQueryReq{
		VendorShopId: "136805",
		CouponCode:   "234937394594",
	})

	couponInfo := channel.CouponStatusQueryResp{}
	contentJson, _ := json.Marshal(content)
	_ = json.Unmarshal(contentJson, &couponInfo)

	t.Log(fmt.Sprintf("%+v\n", couponInfo))
}
