package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/cgi"
	"os"
)

/*
how to run webserver.go
one way:
[/]          cd ~/local_go/src             #your git clone directory
[~/local_go/src] pwd | xargs go run test/webserver.go

another way:
[~/local_go/src]  local_go_pwd=$(pwd)
[~/local_go/src]  go run test/webserver.go $local_go_pwd >>../output.txt 2>&1 &

another way:
change the variable ` pwdDir = "/root/www/local_go" ` to your relevant directory(~/local_go)

*/
func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(net.JoinHostPort("0.0.0.0", "8989"), nil)
	select {}
}

/*
type ResponseWriter
type Handler stuct {
	Path string //path to the CGI executable
	Root string //root URI prefix of handler or empty for "/"
	Dir  string //CGI working directory
}
*/
func handleFunc(res http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/favicon.ico" {
		return
	}
	handler := new(cgi.Handler)

	//path to the CGI executable
	handler.Path = "/usr/local/go/bin/go"

	//go运行的文件所在的目录
	osArgs := os.Args
	var pwdDir string
	if len(osArgs) > 1 {
		pwdDir = osArgs[1]
	}
	if pwdDir == "" {
		panic("lack work directory")
	}
	handler.Dir = pwdDir

	script := handler.Dir + req.URL.Path
	args := []string{"run", script}

	//get 参数获取后,在命令行后添加
	params := req.URL.Query()
	param, ok := params["func"]
	if ok {
		args = append(args, param[0])
	}
	handler.Args = append(handler.Args, args...)

	fmt.Println(handler.Path, handler.Args)
	fmt.Println(os.Getenv("GOPATH"), "after GOPATH")
	fmt.Println("---------------------------------------------")
	handler.ServeHTTP(res, req)
}
