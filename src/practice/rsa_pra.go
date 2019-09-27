package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"hash"
	"strings"

	//"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
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
		"rsa1" : rsa1,
		"rsa2" : rsa2,
		"rsa3" : rsa3,
		"rsa4" : rsa4,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//PKCS1
func rsa1()  {
	//秘钥生成
	pub,pri,err := RsaGenKey(1024,"PKCS1","PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pub)
	fmt.Println("---")
	fmt.Println(pri)

	//公钥加密
	token,err := RSAEncrypt([]byte("sea"),pub,"PKCS1")

	//私钥解密
	msg,err := RSADecrypt(token,pri,"PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(msg))
}
//私钥 PKCS8
func rsa2()  {
	Puk,Prk,err := RsaGenKey(1024,"PKCS8","PKCS1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg       := []byte("sea")
	msgEn,err := RSAEncrypt(msg,Puk,"PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Puk)
	fmt.Println(Prk)

	msgDe,err := RSADecrypt(msgEn,Prk,"PKCS8")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(msgDe))
}
//公钥 PKIX
func rsa3()  {
	Puk,Prk,err := RsaGenKey(1024,"PKCS1","PKIX")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg       := []byte("sea")
	msgEn,err := RSAEncrypt(msg,Puk,"PKIX")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Puk)
	fmt.Println(Prk)

	msgDe,err := RSADecrypt(msgEn,Prk,"PKCS1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(msgDe))
}

//私钥签名,公钥验签
func rsa4()  {
	Puk,Prk,err := RsaGenKey(1024,"PKCS1","PKCS1")
	if err != nil {
		fmt.Println(err.Error())
	}
	msg       := "sea";
	enMsg,err := sign(msg,Prk,"PKCS1","SHA1")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = verfiySign(msg,enMsg,Puk,"PKCS1","SHA1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(true)
}
/**
https://www.jianshu.com/p/60fe90594583 (有一处有问题,已提交评论)
 */
//生成公钥,私钥,公钥加密,私钥解密
func RsaGenKey(bits int,PriEnType string,PubEnType string) (PubKey string,PriKey string,Err error)  {
	//1.生成私钥
	//GenerateKey函数使用随机数据生成器random生成一对具有指定位数的RSA秘钥
	privateKey,err := rsa.GenerateKey(rand.Reader,bits)
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
		derPrivateStream,err = x509.MarshalPKCS8PrivateKey(privateKey)
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
		Type  : "RSA PRIVATE KEY",
		Bytes : derPrivateStream,
	}
	//4.写入缓冲中
	buffPrivate := &bytes.Buffer{}
	err          = pem.Encode(buffPrivate,&block)
	if err != nil {
		Err    = err
		return
	}
	PriKey = buffPrivate.String()

	//1.生成公钥
	var derPublicStream []byte
	publicKey       := privateKey.PublicKey
	if PubEnType == "PKCS1" {
		derPublicStream = x509.MarshalPKCS1PublicKey(&publicKey)
	}
	if PubEnType == "PKIX" {
		derPublicStream,err = x509.MarshalPKIXPublicKey(&publicKey)
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
		Type : "RSA PUBLIC KEY",
		Bytes: derPublicStream,
	}
	buffPublic := &bytes.Buffer{}
	err         = pem.Encode(buffPublic,&block)
	if err != nil {
		Err = err
		return
	}
	PubKey = buffPublic.String()
	return
}

//公钥加密
func RSAEncrypt(enMsg []byte,publicKey string,enType string) (enByte []byte,Err error)  {
	//1.从公钥中找出block和pubKey
	pubKey,err := RSAParsePubKey(publicKey,enType)
	if err != nil {
		Err = err
		return
	}

	encryptedStr,err := rsa.EncryptPKCS1v15(rand.Reader,pubKey,enMsg)
	if err != nil {
		Err = err
		return
	}
	enByte          = encryptedStr
	return
}
//私钥解密
func RSADecrypt(deMsg []byte,privateKey string,enType string) (deByte []byte,Err error) {
	priKey,err := RSAParsePriKey(privateKey,enType)
	if err != nil {
		Err = err
		return
	}


	decryptedStr,err := rsa.DecryptPKCS1v15(rand.Reader,priKey,deMsg)
	if err != nil {
		Err = err
		return
	}
	deByte = decryptedStr
	return
}
//公钥解析
//publicKey 公钥
//enType 公钥解析算法 PKIX|PKCS1
func RSAParsePubKey(publicKey string,enType string) (Pub *rsa.PublicKey,Err error) {
	//1.从公钥中找出block和pubKey
	block,_ := pem.Decode([]byte(publicKey))
	if block == nil {
		Err = errors.New("error publicKey")
		return
	}

	if enType == "PKIX" {
		pu,err := x509.ParsePKIXPublicKey(block.Bytes)
		Pub = pu.(*rsa.PublicKey)
		Err = err
		return
	}
	if enType == "PKCS1" {
		return x509.ParsePKCS1PublicKey(block.Bytes)
	}
	Err = errors.New("error enType")
	return
}

//私钥解析
//privateKey 私钥
//enType 私钥解密算法 PKCS1|PKCS8
func RSAParsePriKey(privateKey string,enType string) (PriKey *rsa.PrivateKey,Err error) {
	//从私钥中找出block和priKey
	block,_ := pem.Decode([]byte(privateKey))
	if block == nil {
		Err = errors.New("error privateKey")
		return
	}
	if enType == "PKCS1" {
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	}
	if enType == "PKCS8" {
		PriKey1,err := x509.ParsePKCS8PrivateKey(block.Bytes)
		PriKey = PriKey1.(*rsa.PrivateKey)
		Err   = err
		return
	}
	return
}
/*
* rsa 数字签名 私钥签名,公钥验签
 */

//私钥签名
//msg 需要加密的数据
//privateKey 私钥
//alg 算法 sha1|sha256
//PriEnType 私钥解密算法 见func RSAParsePriKey()
func sign(msg,privateKey,PriEnType,alg string) (enSign string, Err error) {
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

	priv,err := RSAParsePriKey(privateKey,PriEnType)
	if err != nil {
		Err = err
		return
	}

	hash1.Write([]byte(msg))
	encryptedData,err := rsa.SignPKCS1v15(rand.Reader,priv,algr,hash1.Sum(nil))
	if err != nil {
		Err = err
		return
	}

	enSign = hex.EncodeToString(encryptedData)
	return
}
//公钥验签
func verfiySign(msg,sig,publicKey,pubEnType,alg string) (Err error) {
	var hash1 hash.Hash
	var algr crypto.Hash
	switch strings.ToLower(alg) {
	case "sha1":
		hash1 = sha1.New()
		algr  = crypto.SHA1
	case "sha256":
		hash1 = sha256.New()
		algr  = crypto.SHA256
	default:
		Err = errors.New("alg error")
		return
	}

	pub,err := RSAParsePubKey(publicKey,pubEnType)
	if err != nil {
		Err = err
		return
	}

	var sig1 []byte
	sig1,err = hex.DecodeString(sig)
	if err != nil {
		Err = err
		return
	}
	hash1.Write([]byte(msg))
	return rsa.VerifyPKCS1v15(pub,algr, hash1.Sum(nil), sig1)
}
