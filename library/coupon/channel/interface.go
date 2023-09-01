package channel

import "context"

type Coupon interface {
	QueryByMobile(context.Context, interface{}) (interface{}, error)

	CouponStatusQuery(context.Context, interface{}) (interface{}, error)

	CouponCharge(context.Context, interface{}) (interface{}, error)

	CouponCancel(context.Context, interface{}) (interface{}, error)
}
