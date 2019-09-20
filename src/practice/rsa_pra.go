package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
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
	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
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
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}
func rsa1()  {
	pub,pri,err := RsaGenKey(2048)
	if err != nil {
		fmt.Println(err)
	}

	token,err := RSAEncrypt([]byte("hi"),pub)

	fmt.Println(string(token))

	msg,_ := RSADecrypt(token,pri)

	fmt.Println(string(msg))
}
/**
https://www.jianshu.com/p/60fe90594583 (有一处有问题,已提交评论)
 */
//生成公钥,私钥
func RsaGenKey(bits int) (PubKey string,PriKey string,Err error)  {
	//1.生成私钥
	//GenerateKey函数使用随机数据生成器random生成一对具有指定位数的RSA秘钥
	privateKey,err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		Err = err
		return
	}
	//2.MarshalPKCS1PrivateKey 将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derPrivateStream := x509.MarshalPKCS1PrivateKey(privateKey)
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
	publicKey       := privateKey.PublicKey
	derPublicStream := x509.MarshalPKCS1PublicKey(&publicKey)
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
func RSAEncrypt(enMsg []byte,publicKey string) (enByte []byte,Err error)  {
	//1.从公钥中找出block和pubKey
	block,_ := pem.Decode([]byte(publicKey))
	if block == nil {
		Err = errors.New("error publicKey")
		return
	}

	pubKey,err := x509.ParsePKCS1PublicKey(block.Bytes)
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
func RSADecrypt(deMsg []byte,privateKey string) (deByte []byte,Err error) {
	//从私钥中找出block和priKey
	block,_ := pem.Decode([]byte(privateKey))
	if block == nil {
		Err = errors.New("error privateKey")
		return
	}
	priKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
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