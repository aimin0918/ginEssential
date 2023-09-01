package channel

import (
	"errors"
	"fmt"
	"oceanlearn.teach/ginessential/library/coupon/config"
)

func newCoupon(couponConfig *config.CouponConfig) (Coupon, error) {
	switch config.Channel(couponConfig.Channel) {
	case config.CouponChannelUnknown:
		return nil, errors.New(fmt.Sprintf("Undefined channel %d", couponConfig.Channel))
	case config.CouponChannelMeiTuan:
		return NewMeiTuan(couponConfig), nil
	}
	return nil, errors.New(fmt.Sprintf("Undefined channel %d", couponConfig.Channel))
}

func NewCouponLib(storeName string) (Coupon, error) {
	couponConfig, err := config.GetCouponConfig(storeName)
	if err != nil {
		return nil, err
	}
	return newCoupon(couponConfig)
}
