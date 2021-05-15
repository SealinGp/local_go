package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

/**
wrk:http压测工具,底层封装了epoll和kqueue
1.使用编写lua脚本来写请求信息(设置body,header...)
2.或者使用简单的GET方式作为请求

graphviz : 把二进制的pprof文件转换为svg

pprof    : golang性能能分析包
命令     : go tool pprof

参数说明:
heap    : 堆内存分析
profile : cpu占用时间分析
allocs  : 分配内存次数分析

数据采样方式(生成pprof文件)
web服务   : 导入 _ net/http/pprof 包并新增一个http.ListenAndServe(address,nil), 访问 /debug/pprof查看信息(包含profile)
非web服务 : 手动写 runtime/pprof runtime.StartCPUProfile在代码中调用生成pprof文件 或者 调用第三方封装好的pprof包来生成文件并使用graphviz分析

生成的pprof文件如何分析:
golang二进制文件 = 你当前运行的golang二进制文件
go tool pprof golang二进制文件 pprof文件
go tool pprof http://url

命令行pprof分析的参数说明:

命令top  : 按指标大小列出前10个函数,比如内存是按内存占用前10,cpu是按执行时间耗时最长前10
参数flat : 本函数占用的内存量 flag%:本函数内存占用内存总量百分比 sum%:前面每行flag%的和 cum:累积量 cum%:累积量占用总量百分比

命令list   : 查看某个函数的代码以及该函数每行代码的指标信息,如果函数名不正确,会进行模糊匹配
命令traces : 打印所有调用栈,以及调用栈的指标信息

内存泄漏相关:https://www.lbbniu.com/7449.html
1.goroutine内部某个函数每次调用消耗内存巨大,并且调用完毕内存没有释放,导致内存泄漏,这种属于goroutine内部代码实现有问题导致的内存泄漏
2.goroutine内部某个函数每次调用消耗内存不大,但是调用goroutine的次数非常多,并且由于某些原因导致这些goroutine无法退出,占用的内存不会释放,导致内存泄漏

goroutine内存泄漏(如果不知道何时停止一个goroutine,那么这个goroutine可能导致潜在的内存泄漏)
1.goroutine本身的占用的栈空间造成内存泄漏
2.goroutine中的变量逃逸到堆内存中,导致堆内存泄漏

如何判断是goroutine泄漏引发的内存泄漏?
隔一段时间获取goroutine的数量(go tool pprof http://xxx/debug/pprof/goroutine),如果后面获取的那次,某些goroutine比前一次多,多获取几次,是持续增长的那么久可能是goroutine内存泄漏

定位goroutine泄漏的2中办法
1.web访问
访问http://xxx/debug/pprof/goroutine?debug=1
根据goroutine阻塞数量最多的位置定位
访问http://xxx/debug/pprof/goroutine?debug=2
根据goroutine阻塞时长最长的地方找到对应的调用位置

2.命令行交互

*/

func main() {
	HeapMem()
}

//1.goroutine内部实现变量没有释放造成内存泄漏
func HeapMem() {
	go func() {
		if err := http.ListenAndServe(":9876", nil); err != nil {
			println(err)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}

//2.goroutine内存泄漏,每次调用goroutine后没有正确的退出造成goroutine栈内存占用不释放,造成goroutine内存泄漏
func GoRoutineMem() {
	go func() {
		if err := http.ListenAndServe(":9876", nil); err != nil {
			println(err)
			os.Exit(1)
		}
	}()
	outCh := make(chan int)
	go func() {
		if false {
			<-outCh
		}
		select {}
	}()

	//开启100个goroutine/秒,并且goroutine内阻塞不退出
	tick := time.Tick(time.Second / 100)
	i := 0
	for range tick {
		i++
		fmt.Println(i)
		alloc1(outCh)
	}
}

//goroutine内阻塞不退出
func alloc1(outCh chan<- int) {
	go func() {
		defer fmt.Println("alloc-fm exit")

		//分配内存,假装用一下
		buf := make([]byte, (1<<20)*10)
		_ = len(buf)
		fmt.Println("alloc done")
		outCh <- 0
	}()
}
