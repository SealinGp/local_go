package main;
import (   
  "fmt"
  "net/http"
  "net/http/cgi"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){		
		handler := new(cgi.Handler);
		handler.Path = "/usr/local/go/bin/go";
		handler.Dir = "/root/www/sea/local_go";
		fmt.Println(handler.Path);

		script := "/root/www/sea/local_go" + req.URL.Path;
		args := []string{"run", script};
		handler.Args = append(handler.Args, args...);
		fmt.Println(handler.Args);

		handler.ServeHTTP(res, req);
	});
	http.ListenAndServe(":8989",nil);
	select{};
}