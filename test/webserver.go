package main;
import (   
  "os"
  "fmt"
  "net/http"
  "net/http/cgi"
)
/*
how to run webserver.go
one way:
[/]          cd ~/local_go             #your git clone directory
[~/local_go] pwd | xargs go run test/webserver.go

another way:
[~/local_go]  local_go_pwd=$(pwd)
[~/local_go]  go run test/webserver.go $local_go_pwd >>basic/output.txt 2>&1 &

another way:
change the variable ` pwdDir = "/root/www/local_go" ` to your relevant directory(~/local_go)

*/
func main() {	
	http.HandleFunc("/", handleFunc);
	http.ListenAndServe(":8989",nil);
	select{};
}

func handleFunc(res http.ResponseWriter, req *http.Request) {
	handler      := new(cgi.Handler);

	//path to the CGI executable
	handler.Path  = "/usr/local/go/bin/go";

	//go运行的文件所在的目录
	osArgs          := os.Args;
	var pwdDir string;
	if len(osArgs) > 1 {
        pwdDir = osArgs[1];
	}
	if pwdDir == "" {
		pwdDir = "/root/www/local_go";
		fmt.Println("using default handle directory...");
	}

	handler.Dir   = pwdDir;
	fmt.Println(handler.Path);

	script       := handler.Dir + req.URL.Path;
	args         := []string{"run", script};
    //get 参数获取后,在命令行后添加
	params       := req.URL.Query();
	param,ok     := params["func"];
	if ok {
		args = append(args,param[0]);
	}	

	handler.Args = append(handler.Args, args...);

	fmt.Println(handler.Args);
	handler.ServeHTTP(res, req);
}