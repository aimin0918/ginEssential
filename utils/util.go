package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/forgoer/openssl"

	"math"
	"math/rand"
	"net/url"
	e2 "oceanlearn.teach/ginessential/library/e"
	"os"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/go-ini/ini"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"

	"oceanlearn.teach/ginessential/library/log"
)

// Setup Initialize the utils
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		index := strings.Index(abPath, "whgo_xjd/")
		if index != -1 {
			return abPath[0 : index+len("whgo_xjd/")]
		} else {
			log.Fatal("项目文件夹必须使用whgo_xjd命名", zap.String("path", abPath))
		}
		return abPath
	}
	return ""
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}
	return string(result)
}

func LoadIni(path, section string, v interface{}) {
	abPath := GetCurrentAbPathByCaller()
	cfg, err := ini.Load(abPath + path)
	if err != nil {
		log.Fatal(fmt.Sprintf("setting.Setup, fail to parse '%v': %v", path, err))
	}

	err = cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cfg.MapTo %s err: %v", section, err))
	}
}

func GenRandNum(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func NewOrderNo(prefix string, platform int, customerId int64) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%d%d%s%s", prefix, time.Now().Unix(), platform, Sup(customerId, 3), GenRandNum(2))
}

// 对长度不足n的数字前面补0, 超过n的取模
func Sup(i int64, n int) string {
	j := i % int64(math.Pow10(n))
	m := fmt.Sprintf("%d", j)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func GetYearByOrderNo(prefix string, orderNumber string) (year int, err error) {
	left := len(prefix)
	orderNo := orderNumber[left:]

	if len(orderNo) != 16 {
		err = errors.New("订单号长度不配")
		return
	}

	for _, r := range orderNo {
		if !unicode.IsDigit(r) {
			err = errors.New("订单号格式错误")
			return
		}
	}

	timestamp, _ := strconv.ParseInt(orderNo[:10], 10, 64)

	year, _ = strconv.Atoi(time.Unix(timestamp, 0).Format("2006"))
	return
}

func CoverPhone(phone string) string {
	if phone != "" {
		return phone[:3] + "****" + phone[7:]
	}

	return phone
}

func CoverName(name string) string {
	if name != "" {
		return name[:1] + strings.Repeat("*", len(name)-1)
	}

	return name
}

func InArray(val, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}

	return false
}

//float64乘以特定倍数，返回int64结果
func Float64ToInt(f float64, mulNum int64) int64 {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", f), 64)
	decimalValue := decimal.NewFromFloat(num)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(mulNum))

	return decimalValue.IntPart()
}

//int64除以指定倍数，返回float64
func IntToFloat(i int64, divNum int64) float64 {
	de := decimal.NewFromInt(i)
	de = de.Div(decimal.NewFromInt(divNum))
	f, _ := de.Float64()
	return f
}

//加密
func AesEncrypt(data, aesKey string) string {
	dst, _ := openssl.AesECBEncrypt([]byte(data), []byte(aesKey), openssl.PKCS7_PADDING)
	return base64.StdEncoding.EncodeToString(dst)
}

//解密
func AesDecrypt(data, aesKey string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Info(fmt.Sprintf("%v", err))
		}
	}()
	baseDecrData, _ := base64.StdEncoding.DecodeString(data)
	dst, _ := openssl.AesECBDecrypt(baseDecrData, []byte(aesKey), openssl.PKCS7_PADDING)
	return string(dst)
}

//手机号加密
func PhoneEncrypt(phone string) string {
	return AesEncrypt(phone, e2.AESKEY)
}

//手机号解密
func PhoneDecrypt(phoneCode string) string {
	return AesDecrypt(phoneCode, e2.AESKEY)
}

func UrlSafeB64encode(b []byte) string {
	str := base64.StdEncoding.EncodeToString(b)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	return str

}

func UrlSafeB64decode(str string) (result []byte) {
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	mod4 := len(str) % 4
	if mod4 != 0 {
		str = str + "===="[0:mod4]
	}
	result, _ = base64.StdEncoding.DecodeString(str)
	return
}

func Chunk(arr []string, size int) [][]string {
	var chunks [][]string
	for i := 0; i < len(arr); i += size {
		end := i + size
		if end > len(arr) {
			end = len(arr)
		}
		chunks = append(chunks, arr[i:end])
	}
	return chunks
}

// CustomerEncrypt 用户openID加密
func CustomerEncrypt(customerId string) string {
	return url.QueryEscape(AesEncrypt(customerId, e2.AESKEY))
}

// CustomerDecrypt 用户openID解密
func CustomerDecrypt(customerId string) string {
	id, _ := url.QueryUnescape(customerId)
	return AesDecrypt(id, e2.AESKEY)
}

func CustomerShareUrl(inviteCode string) string {

	url := "https://whgo-xjd.develop.meetwhale.com/qr/page_index?env=dev&type=index&invite_code=" + inviteCode
	if os.Getenv("env") == "prod" {
		url = "https://scrm.xijiade.cn/qr/page_index?env=pro&type=index&invite_code=" + inviteCode
	}
	return url
}
