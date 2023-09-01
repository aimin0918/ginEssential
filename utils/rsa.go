package utils

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// GenRSAKey - 生成RSA密钥对
func GenRSAKey(size int) (privateKeyBytes, publicKeyBytes []byte, err error) {
	//生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return
	}
	privateKeyBytes = x509.MarshalPKCS1PrivateKey(privateKey)
	publicKeyBytes = x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	return
}

///**
//公钥加密
//*/
//func RsaEncrypt(src, publicKeyByte []byte) (bytes []byte, err error) {
//	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyByte)
//	if err != nil {
//		return
//	}
//	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, src)
//}

// RsaEncryptBlock - 公钥加密
func RsaEncryptBlock(src, publicKeyByte []byte) (bytesEncrypt []byte, err error) {
	// 将字符串pem解码
	block, _ := pem.Decode(publicKeyByte)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	pubKey := publicKey.(*rsa.PublicKey)
	keySize, srcSize := pubKey.Size(), len(src)
	// 单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt = buffer.Bytes()
	return
}

///**
//私钥解密
//*/
//func RsaDecrypt(src, privateKeyBytes []byte) (bytesDecrypt []byte, err error) {
//	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
//	if err != nil {
//		return
//	}
//	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)
//}

// RsaDecryptBlock -  私钥解密
func RsaDecryptBlock(src, privateKeyBytes []byte) (bytesDecrypt []byte, err error) {
	//将字符串进行pem解码
	block, _ := pem.Decode(privateKeyBytes)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	keySize := privateKey.Size()
	srcSize := len(src)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt = buffer.Bytes()
	return
}

//  RsaSign - Rsa签名
// plainText 明文
// filePath 私钥文件路径
// 返回签名后的数据 错误
func RsaSign(src, privateKeyByte []byte) ([]byte, error) {
	priKey, err := x509.ParsePKCS8PrivateKey(privateKeyByte)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := priKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid pkcs8 key ")
	}

	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(src))
	hashed := h.Sum(nil)

	signText, err := rsa.SignPKCS1v15(rand.Reader, rsaKey, crypto.SHA1, hashed)
	if err != nil {
		return nil, err
	}
	return signText, nil
}

// Rsa签名验证
// plainText 明文
// filePath 公钥文件路径
// 返回签名后的数据 错误
func RsaVerify(src, publicKeyByte []byte, signText []byte) error {
	pubInter, err := x509.ParsePKIXPublicKey(publicKeyByte)
	if err != nil {
		return err
	}
	pubKey := pubInter.(*rsa.PublicKey)
	hashText := sha512.Sum512([]byte(src))
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA512, hashText[:], signText)
	if err != nil {
		return err
	}
	return nil
}
