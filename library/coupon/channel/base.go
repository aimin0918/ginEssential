package channel

import (
	"context"
	"ginessential/library/coupon/config"
)

type base struct {
	couponConfig *config.CouponConfig
}

func (b *base) GetSafeCouponConfig(ctx context.Context) *config.CouponConfig {
	return &config.CouponConfig{
		AppKey:    b.couponConfig.AppKey,
		AppSecret: b.couponConfig.AppSecret,
		Channel:   b.couponConfig.Channel,
		Version:   b.couponConfig.Version,
		Host:      b.couponConfig.Host,
	}
}
