package main;
import (   
  "fmt"
  "net/http"
  "net/http/cgi"
)

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
	handler.Dir   = "/root/www/sea/local_go";
	fmt.Println(handler.Path);

	script       := "/root/www/sea/local_go" + req.URL.Path;
	args         := []string{"run", script};
    //get 参数获取后,在命令行后添加
	params       := req.URL.Query();
	param,ok     := params["func"];
	if ok {
		args = append(args,param[0]);
	}
	//将错误重定向到webserver中
	// args = append(args,"2>&1");

	handler.Args = append(handler.Args, args...);

	fmt.Println(handler.Args);
	handler.ServeHTTP(res, req);
}