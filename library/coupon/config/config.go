package config

import (
	"errors"
	"oceanlearn.teach/ginessential/utils"
	"os"
)

type CouponConfig struct {
	AppKey    string `json:"AppKey"`
	AppSecret string `json:"AppSecret"`
	Channel   int32  `json:"Channel"`
	Version   string `json:"Version"`
	Host      string `json:"Host"`
}

var configMap = map[string]*CouponConfig{}

func GetCouponConfig(storeConfigKey string) (*CouponConfig, error) {
	if _, ok := configMap[storeConfigKey]; !ok {
		err := loadCouponConfig(storeConfigKey)
		if err != nil {
			return nil, err
		}
	}
	couponConfig := configMap[storeConfigKey]
	if couponConfig == nil {
		return nil, errors.New("invalid config")
	}
	return couponConfig, nil
}

func SetCouponConfig(storeConfigKey string, couponConfig *CouponConfig) error {
	configMap[storeConfigKey] = couponConfig
	return nil
}

func loadCouponConfig(storeConfigKey string) error {
	var couponConfig = &CouponConfig{}
	env := os.Getenv("env")

	if env == "" {
		err := os.Setenv("env", "dev")
		if err != nil {
			return err
		}
		env = "dev"
	}

	path := "conf/" + env + "/app.ini"
	utils.LoadIni(path, storeConfigKey, couponConfig)
	return SetCouponConfig(storeConfigKey, couponConfig)
}
