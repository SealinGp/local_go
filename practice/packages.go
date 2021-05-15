package main

import (
	"fmt"
	"os"
	"sync"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.1.md
9.packages
简单描述不同包的一些功能

archive/tar,zip-compress : 压缩,解压缩文件的功能

*/
type Info struct {
	mu  sync.Mutex
	str string
}

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
		"pack1": pack1,
	}
	if nil == funs[n] {
		fmt.Println("func", n, "unregistered")
		return
	}
	funs[n]()
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/09.3.md
锁,sync包
map类型不存在锁的机制,因此map类型为非线程安全的,当并行访问一个共享的Map类型时,map数据会出错

sync包中有一个RWMutex锁,能通过RLock 多线程共享读,一个线程独占写锁
*/
func pack1() {
	i := Info{
		mu:  sync.Mutex{},
		str: "abc",
	}
	Update(&i, "abc")
	fmt.Println(i.str)
}
func Update(info *Info, str string) {
	info.mu.Lock()
	info.str = str
	info.mu.Unlock()
}

func pack2() {
	/*
		 import . "xxx"
		当使用.作为包别名的时候,可以不通过包名使用包中的函数 xx()

		import _ "xxx"
		只执行xxx包中的init函数,初始化全局变量

		导入外安装包(如github.com)
		1.命令行 go get github.com/xxx/xxx
		2.命令行 go install github.com/xxx/xxx
		2.main主程序代码 import "guthub.com/xxx/xxx"

	*/
}
