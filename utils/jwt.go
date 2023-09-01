package utils

import (
	"github.com/dgrijalva/jwt-go"
	e2 "oceanlearn.teach/ginessential/library/e"

	"time"
)

var jwtSecret []byte

type Claims struct {
	Platform   int    `json:"platform"`
	UserId     int64  `json:"user_id"`
	Uid        string `json:"union_id"`
	UserName   string `json:"user_name"`
	CustomerId int64  `json:"customer_id"`
	jwt.StandardClaims
}

// platform :平台标识 , userId ：粉丝表中的id , userName: 微信是openid 支付宝是user_id ,  customerId:会员id
func GenerateToken(platform int, userId int64, uId, userName string, customerId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(e2.LOGIN_TIME_OUT * time.Hour)
	claims := Claims{
		platform,
		userId,
		uId,
		userName,
		customerId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

//检查token是否过期
func ValidToken(token string) bool {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return false
	}
	return tokenClaims.Valid
}

//从token中解析出来用户信息
func GetClaims(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			return claims, nil
		}
	}
	return nil, err
}
