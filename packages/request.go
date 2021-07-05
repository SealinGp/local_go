package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("func required")
	}
	fuc := map[string]func(){
		"req1": req1,
	}
	fuc[os.Args[1]]()
}

func req1() {
	req, err := http.NewRequest("GET", "https://safety.yjgl.sz.gov.cn/ytsafe/rand/randCode.shtml", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	dst := []byte{}
	base64.StdEncoding.Decode(dst, body)
	base64.StdEncoding.EncodeToString()
	fmt.Println(string(dst))
}
