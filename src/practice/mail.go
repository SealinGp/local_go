package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/smtp"
)
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.12.md
func main()  {
	addr := net.JoinHostPort("smtp.qq.com","25")
	client,err := smtp.Dial(addr)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(1)

	auth := smtp.PlainAuth(
		"",
		"464485940@qq.com",
		"uavgxasmyopebjga",
		"smtp.qq.com",
	)
	err = client.Auth(auth)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(2)

	err = client.Mail("464485940@qq.com")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(3)

	err = client.Rcpt("sealingp@163.com")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(4)

	wc,err := client.Data()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(5)

	defer wc.Close()
	buf := bytes.NewBufferString("发送内容111")
	_,err = buf.WriteTo(wc)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(6)

	//auth := smtp.PlainAuth(
	//	"",
	//	"464485940@qq.com",
	//	"qwerSEA123",
	//	"smtp.qq.com",
	//)
	//smtp.SendMail(addr,auth,"")
}
