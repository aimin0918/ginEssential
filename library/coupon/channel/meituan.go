package channel

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"ginessential/library/coupon/config"
	"ginessential/library/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type meituan struct {
	base
}

const (
	QueryByMobileUri = "tuangou/coupon/queryByMobile"
	StatusQueryUri   = "tuangou/coupon/status/query"
	BatchConsumeUri  = "tuangou/coupon/batchConsume"
	CouponCancelUri  = "tuangou/coupon/cancel"

	CouponCanUseStatus    = 10
	CouponCanNotUseStatus = 20
)

type UserCouponQueryReq struct {
	VendorShopId string `json:"vendorShopId"`
	Mobile       string `json:"mobile"`
}

type UserCouponQueryResp struct {
	CouponCode        string       `json:"couponCode"`
	DealID            int64        `json:"dealId"`
	DealTitle         string       `json:"dealTitle"`
	DealType          int          `json:"dealType"`
	CouponEndTime     int64        `json:"couponEndTime"`
	DealValue         float64      `json:"dealValue"`
	DealPrice         float64      `json:"dealPrice"`
	DealMenu          [][]DealMenu `json:"dealMenu"`
	CouponBuyPlatform int          `json:"couponBuyPlatform"`
}
type DealMenu struct {
	Content       string   `json:"content"`
	Specification string   `json:"specification"`
	Price         string   `json:"price"`
	Total         string   `json:"total"`
	Type          string   `json:"type"`
	Images        []string `json:"images"`
	NotDishes     bool     `json:"notDishes"`
}

type Response struct {
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Content interface{} `json:"content"`
}

type CouponStatusQueryReq struct {
	VendorShopId string `json:"vendorShopId"`
	CouponCode   string `json:"couponCode"`
}

type CouponStatusQueryResp struct {
	Status                     int                  `json:"status"`
	CouponStartTime            string               `json:"couponStartTime"`
	CouponEndTime              string               `json:"couponEndTime"`
	RealAmount                 float64              `json:"realAmount"`
	DealType                   int                  `json:"dealType"`
	DealID                     string               `json:"dealId"`
	DealTitle                  string               `json:"dealTitle"`
	DealPrice                  float64              `json:"dealPrice"`
	DealValue                  float64              `json:"dealValue"`
	UserPhoneTail              string               `json:"userPhoneTail"`
	DealSkuMappingDetail       DealSkuMappingDetail `json:"dealSkuMappingDetail"`
	DealPromotionMappingDetail string               `json:"dealPromotionMappingDetail"`
	CanShareOtherPromotion     bool                 `json:"canShareOtherPromotion"`
	DealRuleCouponLimit        int                  `json:"dealRuleCouponLimit"`
	OpenID                     string               `json:"openId"`
}
type DealSkuMappingDetail struct {
	Count      int      `json:"count"`
	VendorSkus []string `json:"vendorSkus"`
}

type CouponChargeReq struct {
	VendorOrderID  string      `json:"vendorOrderId"`
	ToPayAmount    float64     `json:"toPayAmount"`
	OrderSkus      []OrderSkus `json:"orderSkus"`
	CouponCodes    []string    `json:"couponCodes"`
	VendorShopID   string      `json:"vendorShopId"`
	EID            string      `json:"eId"`
	EName          string      `json:"eName"`
	NumberOfDiners int         `json:"numberOfDiners"`
}
type OrderSkus struct {
	VendorSkuID   string  `json:"vendorSkuId"`
	VendorSkuName string  `json:"vendorSkuName"`
	Unit          string  `json:"unit"`
	UnitPrice     float64 `json:"unitPrice"`
	Count         int     `json:"count"`
}

type CouponCancelReq struct {
	VendorShopID string `json:"vendorShopId"`
	CouponCode   string `json:"couponCode"`
	EID          string `json:"eId"`
	EName        string `json:"eName"`
}

func NewMeiTuan(couponConfig *config.CouponConfig) *meituan {
	return &meituan{
		base: base{couponConfig},
	}
}

func (m *meituan) getSign(ctx context.Context, args interface{}) (string, int64) {
	ts := time.Now().UnixMilli()
	content, _ := json.Marshal(args)

	source := fmt.Sprintf("%sappKey%sts%dversion%s%s", m.couponConfig.AppSecret, m.couponConfig.AppKey, ts, m.couponConfig.Version, content)
	h := sha1.New()
	h.Write([]byte(source))

	return hex.EncodeToString(h.Sum(nil)), ts
}

func (m *meituan) QueryByMobile(ctx context.Context, req interface{}) (content interface{}, err error) {
	body := m.request(ctx, QueryByMobileUri, req)
	content, err = checkResponse(ctx, body)
	if err != nil {
		return
	}

	return
}

func (m *meituan) CouponStatusQuery(ctx context.Context, req interface{}) (content interface{}, err error) {
	body := m.request(ctx, StatusQueryUri, req)
	content, err = checkResponse(ctx, body)
	if err != nil {
		return
	}

	return
}

func (m *meituan) CouponCharge(ctx context.Context, req interface{}) (content interface{}, err error) {
	body := m.request(ctx, BatchConsumeUri, req)
	content, err = checkResponse(ctx, body)
	if err != nil {
		return
	}

	return
}

func (m *meituan) CouponCancel(ctx context.Context, req interface{}) (content interface{}, err error) {
	body := m.request(ctx, CouponCancelUri, req)
	content, err = checkResponse(ctx, body)
	if err != nil {
		return
	}

	return
}

func (m *meituan) request(ctx context.Context, sufferUrl string, args interface{}) (body []byte) {
	sign, ts := m.getSign(ctx, args)
	v := url.Values{}
	v.Add("sign", sign)
	v.Add("ts", strconv.FormatInt(ts, 10))
	v.Add("appKey", m.couponConfig.AppKey)
	v.Add("version", m.couponConfig.Version)

	query := v.Encode()
	de, _ := url.QueryUnescape(query)

	reqUrl := fmt.Sprintf("%s%s?%s", m.couponConfig.Host, sufferUrl, de)

	content, _ := json.Marshal(args)
	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(content))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WarnWithCtx(ctx, "请求失败")
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WarnWithCtx(ctx, "获取返回值失败")
		return
	}

	return
}

func checkResponse(ctx context.Context, body []byte) (content interface{}, err error) {
	resp := Response{}
	_ = json.Unmarshal(body, &resp)

	if resp.Code != 200 {
		err = errors.New(resp.Msg)
		return
	}

	content = resp.Content
	return
}
