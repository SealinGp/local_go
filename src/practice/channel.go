package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"testing"
)

/*
通道,超时,计时器
1 s  = 1000 ms(毫秒)
1 ms = 1000 us(微妙)
1 us = 1000 ns(纳秒)
像圆周率π一样,e为数学中的常数, e = 2.71828182856904523536...
*/
//func init() {
//	fmt.Println("Content-Type:text/plain;charset=utf-8\n\n")
//}
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
		"cha1"   : cha1,
		"cha2"   : cha2,
		"cha3"   : cha3,
		"cha4"   : cha4,
		"cha5"   : cha5,
		"cha6"   : cha6,
		"cha7"   : cha7,
		"cha8"   : cha8,
		"cha9"   : cha9,
		"cha10"  : cha10,
		"cha11"  : cha11,
		"cha12"  : cha12,
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.6.md
func cha1()  {
	/**
		协程和恢复(recover)
		func server(workChan <-chan *Work) {
			for work := range workChan {
				go safelyDo(work)
			}
		}
		func safelyDo(work *Work) {
			defer func(){
				if err := recover(); err != nil {
					log.Println(err.Error(),work)
				}
			}()
			do(work)
		}
	上边的代码 如果do(work)发生panic ,错误会被记录且当前协程会退出释放,而其他协程不会受影响.
	因为 recover 总是返回nil, 除非直接在defer修饰的函数中调用,defer修饰的代码可以调用那些
	自身可以使用panic和recover 避免失败的库例程
	 */
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.7.md

/**
任务和worker
假设我们需要处理很多任务,一个worker处理一项任务,一个任务可以被定为一个结构体
*/
type Task struct {
	a string
	//属性...
}

/**
旧模式,共享内存,由各个任务组成的任务池共享内存,避免资源竞争,需要对任务池加锁(java,C++,C#)
 */
type TaskPool struct {
	Mu sync.Mutex  `description:"是互斥锁,用来在代码中保护临界区资源,同一时间只有一个go协程可以进入该区域"`
	Tasks []*Task
}
func cha2()  {
	TP      := new(TaskPool)
	taskNum := 3
	for i := 0; i < taskNum; i++ {
		TP.Tasks = append(TP.Tasks,new(Task))
	}
	/**
	这些worker有很多可以并发执行,他们可以在go协程中启动,一个worker先将pool锁定
	从Pool中获取第一个任务,然后解锁跟处理任务,加锁保证了同一时间只有一个go协程可以进入到Pool
	中,一个任务有且只能被赋予一个worker,加锁实现同步的方式在工作协程比较少的情况下可以工作的很好
	但是当协程数量很多时,处理效率会因为频繁的加锁/解锁开销而降低 ,当工作协程数增加到一个阈值时,
	程序效率急剧下降,成为了瓶颈
	 */
	go Worker(TP)
	go Worker(TP)
}
func Worker(pool *TaskPool)  {
	for {
		//加锁
		pool.Mu.Lock()

		//取出当前任务
		task      := pool.Tasks[0]
		//更新任务池
		pool.Tasks = pool.Tasks[1:]

		//解锁
		pool.Mu.Unlock()
		//执行任务
		task.process()
	}
}
//任务执行的具体内容
func (t *Task)process()  {
	t.a = "abc"
}

/**
其他模式:使用通道
一个通道接受需要处理的任务(队列),一个通道接受处理完成的任务
worker在协程中启动,其数量N应该根据任务数量进行调整
主线程扮演master节点角色
 */
func cha3()  {
	N            := 5
	pending,done := make(chan *Task,N),make(chan *Task)
	finished     := make(chan bool)

	//老板分发任务,同时分发N个 给5人,5人同时工作
	//send work
	go func() {
		for N > 0  {
			N--
			pending <- new(Task)
		}
		close(pending)
	}()
	//exec work
	go func() {
		for pendWork := range pending {
			pendWork.process()
			done <- pendWork
		}
		close(done)
	}()

	//老板接收任务结果,一个个接收
	//received work
	go func() {
		for doneWork := range done  {
			fmt.Println(doneWork.a)
		}
		finished<-true
	}()

	fin := <-finished
	fmt.Println("finished:",fin)
}
/**
总结
使用锁的情景:
	访问共享数据结构中的缓存信息
	保存应用程序上下文和状态信息数据

使用通道的情景:
	与异步操作的结果进行交互
	分发任务
	传递数据所有权
 */


/**
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.8.md
惰性生成器的实现(在需要时求值)
	通道+协程
 */

//惰性求值 例子1
func cha4()  {
	ints = yield()
	fmt.Println(generate())
	fmt.Println(generate())
	fmt.Println(generate())
}
var ints chan int
func yield() chan int  {
	ch    := make(chan int)
	count := 0
	go func() {
		for {
			fmt.Println("gorountine")
			ch <- count
			count++
		}
	}()
	fmt.Println(count,"yield end")
	return ch
}
func generate() int {
	return <-ints
}

//惰性求值 例子2,一次获取一个偶数
func cha5()  {
	//1.具体计算过程
	evenFunc := func(state Any) (Any,Any) {
		OS,ok := state.(int)
		ns    := 0
		if ok {
			ns = OS + 2
		}
		return OS,ns
	}
	event := cha5_2(evenFunc,0)

	//4.消费者拿取计算结果
	for i := 0; i < 10; i++ {
		fmt.Println(event())
	}
}
type Any interface {}
type EvalFunc func(any Any) (Any,Any)
//3.生产者开协程具体计算
func cha5_1(e EvalFunc,a Any) func() Any {
	ch       := make(chan Any)

	//开启协程通过函数计算出需要的东西,写入管道,函数e为具体计算过程
	go func() {
		as  := a
		var retVal Any
		for {
			retVal,as = e(as)
			ch <- retVal
		}
	}()

	//将管道中的内容,通过返回函数来返回(需要的时候再调用)
	return func() Any {
		return <-ch
	}
}
//2.
func cha5_2(evalFunc EvalFunc,initState Any) func() int  {
	ef := cha5_1(evalFunc,initState)
	return func() int {
		return ef().(int)
	}
}

//惰性求值 练习1
func cha6()  {
	pro := func() uint64 {
		var ap uint64
		ap    = 1
		le   := len(fibArr)
		if le >= 2 {
			ap  = fibArr[le-1] + fibArr[le-2]
		}
		fibArr = append(fibArr,ap)
		return fibArr[len(fibArr)-1]
	}
	g := getFib(pro)
	for i := 0; i < 10; i++ {
		fmt.Println(g())
	}
}
var fibArr []uint64
func generateFib(pro func()uint64) func() uint64  {
	fibch := make(chan uint64)

	go func() {
		for {
			fib := pro()
			fibch <- fib
		}
	}()

	return func() uint64 {
		return <-fibch
	}
}
func getFib(pro func()uint64) func() uint64  {
	fib := generateFib(pro)
	return func() uint64 {
		return fib()
	}
}

/**
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.9.md
Futures模式:
	使用某一个值之先对其进行计算,这种情况下,你可以在另一个处理器上计算该值,
	需要使用时,就已经计算好了
闭包+通道 可实现,类似于生成器
 */
func cha7()  {
	/**
	func InverseProduct(a Matrix, b Matrix) {
	    a_inv_future := InverseFuture(a)   // start as a goroutine
	    b_inv_future := InverseFuture(b)   // start as a goroutine
	    a_inv := <-a_inv_future
	    b_inv := <-b_inv_future
	    return Product(a_inv, b_inv)
	}

	func InverseFuture(a Matrix) chan Matrix {
	    future := make(chan Matrix)
	    go func() {
	        future <- Inverse(a)
	    }()
	    return future
	}
	 */
}

/**
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.10.md
复用

c/s 模式 客户端/服务端
 */
func cha8()  {
	//1.服务端具体的处理逻辑
	var op binOp = func(a,b int) int {
		return a+b
	}

	//5.服务端等待客户端请求
	reqChan,quitChan  := startServer(op)

	//6.客户端开始请求,模拟100个客户端请求
	const N = 100
	var reqs [N]Request
	for i := 0; i < N ; i++ {
		//客户端请求的参数
		req       := &reqs[i]
		req.a      = i
		req.b      = i + N
		req.replyc = make(chan int)

		//客户端请求开始
		reqChan <- req
	}

	//8.服务端处理完成后返回客户端
	for i := N-1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("failed at",i)
		} else {
			fmt.Println("Request",i,"ok")
		}
	}

	//14.10.2 通过信号通道关闭服务器 关闭已处理完的协程
	quitChan<-true

	fmt.Println("done")
}
//客户端的请求,自身中包含一个通道,而服务器向这个通道发送响应(告知其服务端是否处理完成)
type Request struct {
	a,b int
	replyc chan int
}
type binOp func(a,b int) int
//2.服务端调用具体的实现方法,并返回到请求的管道中
func run(op binOp,req *Request)  {
	req.replyc <- op(req.a,req.b)
}
//3.服务端开启协程的操作
func server(op binOp,service chan *Request,quit <-chan bool)  {
	for {
		select {
		case req := <-service:
			go run(op,req)
		case <-quit:
			return
		}
		//一个客户端请求一个协程
		req := <-service
		go run(op,req)
	}
}
//4.服务端启动分发的协程
func startServer(op binOp) (chan *Request,chan bool)  {
	reqChan := make(chan *Request)
	quitChan := make(chan bool)
	go server(op,reqChan,quitChan)
	return reqChan,quitChan
}

/**
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.11.md
限制同时处理的请求数量: 带缓冲的管道
充分利用有限的资源
 */
func cha9()  {
	service := make(chan *Request)
	go server1(service)
}
const maxReqs = 50
var  sem = make(chan bool,maxReqs)
func process(r *Request)  {
	r.a++
	r.b++
}
func handle(r *Request)  {
	//处理一次,计数一次,直到计数管道缓冲区满,然后阻塞
	sem <- true
	process(r)
	<-sem
}
func server1(service chan *Request)  {
	for {
		request := <-service
		go handle(request)
	}
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.12.md
链式协程 ? 暂未弄清楚
 */
var ngoroutine = flag.Int("n",5,"how many goroutines")
func cha10()  {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil,leftmost

	for i := 0; i < *ngoroutine ; i++ {
		left, right = right, make(chan int)
		go cha10_1(left,right)
	}

	right <- 0

	x := <-leftmost
	fmt.Println(x)
}
//left = 1+0
func cha10_1(left,right chan int)  {
	left <- 1 + <-right
}

/**
14.13
在多核机器上并行计算
假设 NCPU 个核心 对应一个四核处理器 将计算量分成NCPU个部分,每一个部分和其他部分并行
 */
const NCPU = 4
func cha11()  {
	runtime.GOMAXPROCS(NCPU)
	cha11_1()
}
func cha11_1()  {
	sem   := make(chan int,NCPU)
	for i := 0; i < NCPU ; i++ {
		go func(s chan<- int) {
			s <- 1
		}(sem)
	}

	for i := 0; i < NCPU ; i++ {
		<-sem
	}
}

/**
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.15.md
漏桶算法
客户端协程执行一个无限循环从某个源头,也许是网络接收数据,数据读取到Buffer类型的缓冲区,为了避免分配过多的
缓冲区以及释放缓冲区,他暴露了一份空闲缓冲区列表,并且使用一个缓冲管道来表示这个列表
var freeList = make(chan *Buffer,100)
这个可重用的缓冲区队列(freeList)是与服务器共享的,当接收数据时,客户端尝试从freeList获取缓冲区,
但如果此时通道为空,则会分配新的缓冲区,一旦消息被加载后,它将被发送到服务器上的serverChan管道 var serverChan = make(chan *Buffer)
 */

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.16.md
func cha12()  {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}
func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}