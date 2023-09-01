package utils

import (
	"fmt"
	"testing"
)

//func TestRsa(t *testing.T) {
//	var mingwen = "测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试"
//	//RSA的内容使用base64打印
//	privateKey, publicKey, _ := GenRSAKey(1024)
//	log.Println("rsa私钥:\t", base64.StdEncoding.EncodeToString(privateKey))
//	log.Println("rsa公钥:\t", base64.StdEncoding.EncodeToString(publicKey))
//	//分段加密
//	miwen, err := RsaEncryptBlock([]byte(mingwen), publicKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("加密后：\t", base64.StdEncoding.EncodeToString(miwen))
//	//分段解密
//	jiemi, err := RsaDecryptBlock(miwen, privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("解密后：\t", string(jiemi))
//}

func TestAse(t *testing.T) {
	data := "YRntSq16j1LrubuEi5+hFQ=="
	rets := PhoneDecrypt(data)
	fmt.Println("1=====:", rets)
}

func TestRsaEncryptBlock(t *testing.T) {
	src := `{"customer":{"id":10,"created_at":"2022-12-27 17:22:45","updated_at":"2022-12-29 16:53:45","phone":"WAaQEoAwjRxtTdpweN+eTA==","point":4,"rank_id":66,"rank_validity_time":1703762003,"rank_promoted_type":1,"validity_amount":12,"validity_rate":0,"current_rate":3,"company_id":1,"shop_invite_code":"9","staff_invite_code":"","remark":"","platform":1,"rank_version":1672226003},"point_detail":{"id":0,"created_at":null,"updated_at":null,"point":1,"customer_id":10,"event":1,"expire_time":1672502399,"order_id":131,"order_no":"1672304014101075","rank_benefit_id":114,"state":0,"remain_point":1,"reason":"订单积分","source_type":1,"source_id":"1672304014101075"},"wx_open_id":"oBeQn5NmfZ34GUbp0qek-EdXLoLA","wx_union_id":"","ali_user_id":"","phone":"WAaQEoAwjRxtTdpweN+eTA==","platform":1}`
	bytesEncrypt, err := RsaEncryptBlock([]byte(src), []byte(`-----BEGIN rsa public key-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDy1/FLpQqRvtZLqNcFjXhlWrSL
5gL0Vm4kGsdqtG6WIzqFxIw1gP/SAEMUIaF0uNDgUcMbo4fRmHpQaKsOW81YY1QI
bxk2S6tU+9ELQYlOZFuFZ5jIip74vS6Lx/J/IiXbMnRvpKc0YDI9xzPx7aRrjHlG
ARJDnvZhCVnb9HHZFwIDAQAB
-----END rsa public key-----`))
	t.Error(err)
	t.Error(string(bytesEncrypt))

	text, err := RsaDecryptBlock(bytesEncrypt, []byte(`-----BEGIN rsa private key-----
MIICXAIBAAKBgQDy1/FLpQqRvtZLqNcFjXhlWrSL5gL0Vm4kGsdqtG6WIzqFxIw1
gP/SAEMUIaF0uNDgUcMbo4fRmHpQaKsOW81YY1QIbxk2S6tU+9ELQYlOZFuFZ5jI
ip74vS6Lx/J/IiXbMnRvpKc0YDI9xzPx7aRrjHlGARJDnvZhCVnb9HHZFwIDAQAB
AoGBAL0qn6kIQCp2GOJI/G4z3IQ/WwLbQpPou9VeEtc5BCfp+012ZK3M9fo1AAuv
guC0kukaZ7yg70zC1QzL6+u8cUCTcq7POxZdSzARRKcFE/Mc3vrFXwetmeNse7aP
y4U62jcQVrw+Y6IkBSFVfMYqeNGPOCT3Umy8vq2bhxo8zOchAkEA+lmpr1HD+lUe
BAuCcXfTp76mJLiVrFGzfVrXqZu4KdYY26GMHniGldcq3atJatUAgc1Tri0+3wsH
ys9ti9z3JwJBAPhS6VfILkNEoHFJYbXXlHpCXEnP3GTAEcheXzV4i52JAEqrXvAL
yOBCFZobbVZpPZlolalsqbqcn2wo8uJwxJECQAeGfnVIre1udZKFjgw/H9ug/XmJ
GuatJgoUmvr8NVL8no6rknyv/suuRhmXtoNBl9xPAb7wmT03JarRBWf44m0CQHQ/
nC0T8VRcVB+0kqFmAoQZfMqxHCOuJqT+SOPnQrTE5fYOs6r8WVVimmpCXLUPH18p
rqZZ0DskBx3DLbEDyYECQAcCfn7XPCLALkhrSQRrTVcD0wrAselCIwXSnBWFkUrT
1HOam25yHeppBB0AxCIH7S82yCCcpgcV5MAHbhNSMA8=
-----END rsa private key-----
`))
	t.Error(err)
	t.Error(string(text))
	if err != nil {
		return
	}
}
