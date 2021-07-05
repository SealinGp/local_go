package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"os"
	"strings"
)

/*
生成公钥私钥方式
rsa 非对称加密解密方式
*/
func init() {
	fmt.Println("Content-Type:text/plain;charset=gbk2312\n\n")
}
func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("lack param ?func=xxx")
		return
	}

	execute(args[1])
}

func execute(n string) {
	funs := map[string]func(){
		"rsa1": rsa1,
		"rsa2": rsa2,
		"rsa3": rsa3,
		"rsa4": rsa4,
		"rsa5": rsa5,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

//-----example----- start
//PKCS1
func rsa1() {
	//秘钥生成
	pub, pri, err := RsaGenKey(1024, "PKCS1", "PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pub)
	fmt.Println("---")
	fmt.Println(pri)

	//公钥加密
	token, err := RSAEncrypt([]byte("sea"), pub, "PKCS1")

	//私钥解密
	msg, err := RSADecrypt(token, pri, "PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(msg))
}

//私钥 PKCS8
func rsa2() {
	Puk, Prk, err := RsaGenKey(1024, "PKCS8", "PKCS1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg := []byte("sea")
	msgEn, err := RSAEncrypt(msg, Puk, "PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Puk)
	fmt.Println(Prk)

	msgDe, err := RSADecrypt(msgEn, Prk, "PKCS8")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(msgDe))
}

//公钥 PKIX
func rsa3() {
	Puk, Prk, err := RsaGenKey(1024, "PKCS1", "PKIX")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg := []byte("sea")
	msgEn, err := RSAEncrypt(msg, Puk, "PKIX")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Puk)
	fmt.Println(Prk)

	msgDe, err := RSADecrypt(msgEn, Prk, "PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(msgDe))
}

//私钥签名,公钥验签
func rsa4() {
	Puk, Prk, err := RsaGenKey(1024, "PKCS1", "PKCS1")
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := "sea"
	enMsg, err := sign(msg, Prk, "PKCS1", "SHA1")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = verfiySign(msg, enMsg, Puk, "PKCS1", "SHA1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(true)
}
func rsa5() {
	PrivateKey := "-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIICXQIBAAKBgQDl5UrPMf4qQ3UaLII0XLETs+XRzHVCgqwjtskFHvLXwK/Z1Vr4\n" +
		"RSr7mcrjfWVrCSINhiRjY4Xis5I8rgcgGw7RGmS0PWS2TLaspTMfxqQ3T7hbcsil\n" +
		"2Ss+pH2jRwtkn2iOVhgk9vifb7sBJrLsVCrjzeiaeZaNc4KcGlUXaVcNTQIDAQAB\n" +
		"AoGBAIesk0LGOT6OAw0IWWs3jNWY5Le1FzrCTX7iP65C/oQv1lgTXxWIFH7Z22/4\n" +
		"MCNEB5G9qbnyITCSU2p2NgRPk6TbY3hfNWkhoGyudtINJ314r0/X/8XjgOd292M8\n" +
		"uvy9pg8GuGT8TGoxCs+F/vXQxDw6HmNHsAUHiTRPPeek7YqtAkEA5vzjt0ZC0L77\n" +
		"PH1s01S7LpSOXxq3hmRlfmjI6VDlcG46o/g2Li0YRsL4TrqrxBHOkph7mUUGwgXt\n" +
		"3es/IZNylwJBAP7KIICgbmTgiAc8eFcN1ZOOUo7/wuvXNd4RjjAs9GZ/qR79biXN\n" +
		"LD8Ch0Xd+fDYrk4c/lE41uXQjn0a+Xtgj7sCQBpXqNiT6LbJsPk7DJglR5uOUZZD\n" +
		"A78N4A1EgfUpxqDF0WY1vmgRuH0Jayv/WetoZHiPbzkRiC3EY1Y1p+N6X00CQE13\n" +
		"e0ZggPAe7Hz2v8gIJsXEYgmkbclzF6e7QrYXFQANFIidmV3Y8fj+dc6iXRoDZ4vM\n" +
		"eO6ND5m0PX6AMxZ2F30CQQDVkvKY4zRE8mCzc3dSQSK52eWkKKLahCo95Zpgu1vq\n" +
		"W56bL/6ztCkqPytHdqtbql0GI0Kxd4NRffIpH5oGKwGG\n" +
		"-----END RSA PRIVATE KEY-----"
	enData := "G30AIrZ5piuBA0P30k6bara1HhKGzKT4mwcpXEBbkQtAq6kcZPpmUms1gJP8J" +
		"PTkIVLayQCU1+vFUmO64Y30rmziaCrh0WiCfFlEkJqUGsJkCCSSlUzEI6yWTiVDj9i8" +
		"aG/ZVsqnY/Gp2E4UlicEgYkp1XA1MZnSYjl0n9QwRkc="

	deData, err := RSADecrypt([]byte(enData), PrivateKey, "PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(deData))
}

//-----example----- end

//-----实现方法----- start
/**
https://www.jianshu.com/p/60fe90594583 (有一处有问题,已提交评论)
*/
//生成公钥,私钥(公钥加密,私钥解密|私钥签名,公钥验签)
//bits 位数,一般为1024/2048
//PriEnType 私钥加密算法 PKCS1 | PKCS8
//PubEnType 公钥加密算法 PKCS1 | PKIX
//PubKey 生成的公钥
//PriKey 生成的私钥
func RsaGenKey(bits int, PriEnType string, PubEnType string) (PubKey string, PriKey string, Err error) {
	//1.生成私钥
	//GenerateKey函数使用随机数据生成器random生成一对具有指定位数的RSA秘钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		Err = err
		return
	}
	//2.MarshalPKCS1PrivateKey 将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	var derPrivateStream []byte
	if PriEnType == "PKCS1" {
		derPrivateStream = x509.MarshalPKCS1PrivateKey(privateKey)
	}

	if PriEnType == "PKCS8" {
		derPrivateStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			Err = err
			return
		}
	}
	if derPrivateStream == nil {
		Err = errors.New("privateKe enType:PKCS1 | PKCS8")
		return
	}

	//3.Block代表PEM编码的结构,对其配置
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derPrivateStream,
	}
	//4.写入缓冲中
	buffPrivate := &bytes.Buffer{}
	err = pem.Encode(buffPrivate, &block)
	if err != nil {
		Err = err
		return
	}
	PriKey = buffPrivate.String()

	//1.生成公钥
	var derPublicStream []byte
	publicKey := privateKey.PublicKey
	if PubEnType == "PKCS1" {
		derPublicStream = x509.MarshalPKCS1PublicKey(&publicKey)
	}
	if PubEnType == "PKIX" {
		derPublicStream, err = x509.MarshalPKIXPublicKey(&publicKey)
		if err != nil {
			Err = err
			return
		}
	}
	if derPublicStream == nil {
		Err = errors.New("PublicKey enType:PKCS1 | PKIX")
		return
	}

	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPublicStream,
	}
	buffPublic := &bytes.Buffer{}
	err = pem.Encode(buffPublic, &block)
	if err != nil {
		Err = err
		return
	}
	PubKey = buffPublic.String()
	return
}

//公钥加密
//enMsg 需要加密的数据
//publicKey 公钥
//PubEnType 公钥加密算法 PKCS1 | PKIX
//enByte    加密后的数据
func RSAEncrypt(enMsg []byte, publicKey string, PubEnType string) (enByte []byte, Err error) {
	//1.从公钥中找出block和pubKey
	pubKey, err := rsaParsePubKey(publicKey, PubEnType)
	if err != nil {
		Err = err
		return
	}

	encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, enMsg)
	if err != nil {
		Err = err
		return
	}
	encryptedStr := base64.StdEncoding.EncodeToString(encryptedBytes)
	enByte = []byte(encryptedStr)
	return
}

//私钥解密
//deMsg 公钥加密后的数据
//privateKey 私钥
//PriEnType  私钥加密算法 PKCS1 | PKCS8
//deByte     解密后的数据
func RSADecrypt(deMsg []byte, privateKey string, PriEnType string) (deByte []byte, Err error) {
	dst, err := base64.StdEncoding.DecodeString(string(deMsg))
	if err == nil {
		deMsg = dst
	}

	priKey, err := rsaParsePriKey(privateKey, PriEnType)
	if err != nil {
		Err = err
		return
	}
	decryptedStr, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, deMsg)
	if err != nil {
		Err = err
		return
	}
	deByte = decryptedStr
	return
}

//公钥解析
//publicKey 公钥
//PubEnType 公钥解析算法 PKIX|PKCS1
func rsaParsePubKey(publicKey string, PubEnType string) (Pub *rsa.PublicKey, Err error) {
	//1.从公钥中找出block和pubKey
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		Err = errors.New("error publicKey")
		return
	}

	if PubEnType == "PKIX" {
		pu, err := x509.ParsePKIXPublicKey(block.Bytes)
		Pub = pu.(*rsa.PublicKey)
		Err = err
		return
	}
	if PubEnType == "PKCS1" {
		return x509.ParsePKCS1PublicKey(block.Bytes)
	}
	Err = errors.New("error enType")
	return
}

//私钥解析
//privateKey 私钥
//PriEnType 私钥解密算法 PKCS1|PKCS8
func rsaParsePriKey(privateKey string, PriEnType string) (PriKey *rsa.PrivateKey, Err error) {
	//从私钥中找出block和priKey
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		Err = errors.New("error privateKey")
		return
	}
	if PriEnType == "PKCS1" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	if PriEnType == "PKCS8" {
		PriKey1, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		PriKey = PriKey1.(*rsa.PrivateKey)
		Err = err
		return
	}
	Err = errors.New("PriEnType error")
	return
}

/*
* rsa 数字签名 私钥签名,公钥验签
 */

//私钥签名
//msg 根据此消息生成签名
//privateKey 私钥
//alg 算法 sha1|sha256
//PriEnType 私钥解密算法 PKCS1|PKCS8
func sign(msg, privateKey, PriEnType, alg string) (enSign string, Err error) {
	var hash1 hash.Hash
	var algr crypto.Hash
	switch strings.ToLower(alg) {
	case "sha1":
		hash1 = sha1.New()
		algr = crypto.SHA1
	case "sha256":
		hash1 = sha256.New()
		algr = crypto.SHA256
	default:
		Err = errors.New("alg error")
		return
	}

	priv, err := rsaParsePriKey(privateKey, PriEnType)
	if err != nil {
		Err = err
		return
	}

	hash1.Write([]byte(msg))
	encryptedData, err := rsa.SignPKCS1v15(rand.Reader, priv, algr, hash1.Sum(nil))
	if err != nil {
		Err = err
		return
	}

	enSign = hex.EncodeToString(encryptedData)
	return
}

//公钥验签
//msg 根据此消息验证签名
//sig 签名信息
//publicKey 公钥
//pubEnType 公钥解析算法 PKIX|PKCS1
func verfiySign(msg, sig, publicKey, pubEnType, alg string) (Err error) {
	var hash1 hash.Hash
	var algr crypto.Hash
	switch strings.ToLower(alg) {
	case "sha1":
		hash1 = sha1.New()
		algr = crypto.SHA1
	case "sha256":
		hash1 = sha256.New()
		algr = crypto.SHA256
	default:
		Err = errors.New("alg error")
		return
	}

	pub, err := rsaParsePubKey(publicKey, pubEnType)
	if err != nil {
		Err = err
		return
	}

	var sig1 []byte
	sig1, err = hex.DecodeString(sig)
	if err != nil {
		Err = err
		return
	}
	hash1.Write([]byte(msg))
	return rsa.VerifyPKCS1v15(pub, algr, hash1.Sum(nil), sig1)
}

//-----实现方法----- end
