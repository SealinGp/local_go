package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/cgi"
	"os"
	"strings"
	"time"
)

const (
	timeLayOut = "2006/01/02 15:04:05"
)
/*
go environment explain (such as $GOXXX)
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/02.2.md

how to run webserver.go
one way:
[/]          cd ~/local_go/src             #your git clone directory
[~/local_go/src] pwd | xargs go run test/webserver.go

another way:
[~/local_go/src]  local_go_pwd=$(pwd)
[~/local_go/src]  go run test/webserver.go $local_go_pwd $GOBIN/go >>../debugger.log 2>&1 &

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
	Env  string[] //环境变量配置,[]string{"key=value"}
}
*/
func handleFunc(res http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "/favicon.ico" {
		return
	}
	handler := new(cgi.Handler)

	envs := map[string]string{
		"HOME"    : os.Getenv("HOME"),
		"GOCACHE" : os.Getenv("GOCACHE"),
	}
	for key,val := range envs {
		en := strings.Join([]string{key,val},"=")
		handler.Env = append(handler.Env,en)
	}

	//path to the CGI executable
	handler.Path = "/usr/local/go/bin/go"

	//go运行的文件所在的目录
	osArgs := os.Args
	var pwdDir string
	var goSrciptPath string
	if len(osArgs) > 1 {
		pwdDir       = osArgs[1]
	}
	if len(osArgs) > 2 {
		goSrciptPath = osArgs[2]
	}
	if pwdDir == "" {
		fmt.Println("lack work directory")
		os.Exit(1)
	}
	if goSrciptPath != "" {
		handler.Path = goSrciptPath
	}

	//~/local_go/src
	handler.Dir = pwdDir

	//~/local_go/src/practice/array.go
	script := handler.Dir + req.URL.Path
	args := []string{"run", script}

	//get 参数获取后,在命令行后添加
	params := req.URL.Query()
	param, ok := params["func"]
	if ok {
		args = append(args, param[0])
	}
	handler.Args = append(handler.Args, args...)

	fmt.Println("--------------------------------------")
	fmt.Println(time.Now().Format(timeLayOut))
	fmt.Println(handler.Path, handler.Args)
	handler.ServeHTTP(res, req)
}
