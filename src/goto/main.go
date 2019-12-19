package main

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/19.3.md
url缩短练习版本1
目的: 长url -> 短url
访问 短url 的时候,将其重定向到 长url所在页面
*/
import (
	"fmt"
	"net/http"
)


const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
var store = NewURLStore()

func main() {
	http.HandleFunc("/",Redirect)
	http.HandleFunc("/add",Add)
	http.ListenAndServe(":8988",nil)
}

func Redirect(rw http.ResponseWriter,rq *http.Request)  {
	key := rq.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(rw,rq)
		return
	}
	http.Redirect(rw,rq,url,http.StatusFound)
}
func Add(rw http.ResponseWriter,rq *http.Request)  {
	url := rq.FormValue("url")
	if url == "" {
		rw.Header().Set("Content-type","text/html")
		fmt.Fprint(rw,AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(rw,"http://localhost:8988/%s",key)
}