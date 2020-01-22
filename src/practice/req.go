package main

import (
	"fmt"
	"io/ioutil"
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
func req1()  {

}

func httpSimpleGet(url string) ([]byte,error)  {
	resp,err :=	http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpSend(method,url,body string,header map[string]string) ([]byte,error) {
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