package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

/*
7.array slice []bytes string
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
		"req1":req1,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}
// base64原理
// http://www.ruanyifeng.com/blog/2008/06/base64.html
func req1()  {
	verCodeUrl := "https://safety.yjgl.sz.gov.cn/ytsafe/tooken/tooken.shtml?r=" + fmt.Sprint(rand.Float64())
	fmt.Println(verCodeUrl)


	//42 + 62 = ?
	verCode  := "d4fe98e5c14147ee90426f410986d9a91583761368950_219134115154"
	randCode := "13531714e0954da2b3bd3845daad5d47"
	v := base64.StdEncoding.EncodeToString([]byte(verCode))
	r := base64.StdEncoding.EncodeToString([]byte(randCode))
	fmt.Println(v)
	fmt.Println(r)
}

func httpGet(url string) ([]byte,error)  {
	resp,err :=	http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpReq(method,url,body string,header map[string]string) ([]byte,error) {
	client := &http.Client{}
	req,err := http.NewRequest(method,url,strings.NewReader(body))
	if err != nil {
		return nil,err
	}
	for k,v := range header {
		req.Header.Add(k,v)
	}
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	return ioutil.ReadAll(resp.Body)
}