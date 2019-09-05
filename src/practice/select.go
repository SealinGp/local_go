package main

import (
	"fmt"
	"os"
	//"runtime"
	"sync"
	"time"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.5.md
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
		"sel1"  : sel1,
		"sel2"  : sel2,
		"sel3"  : sel3,
		"sel4"  : sel4,
		"sel5"  : sel5,
		"sel6"  : sel6,
		"sel7"  : sel7,
		"sel8"  : sel8,
		"sel9"  : sel9,
		"sel10" : sel10,

	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}

func sel1()  {
	/*
	 * 以d为周期给返回的通道发送时间,d单位纳秒,作用通常用于限制请求频率
	   chan := time.Tick(d)
	   rate_per_sec := 10
	   var dur Duration = 1e9
	   chRate := time.Tick(dur)
	   for req := range requests {
	       <- chRate // rate limit our Service.Method RPC calls,每 1e9 纳秒阻塞一次,限制调用
	       go client.Call("Service.Method", req, ...)
	   }

	* 若周期请求数量暴增,可使用带缓冲的通道和计时器对象来承载


	* 定时器Timer 跟 计时器Ticker 结构体很像,但Timer 在规定时间后,只发送一次时间
	* time.After(d duration)<-chan Time 在时间d 后,当前时间被发到返回的通道中, 因此它和NewTimer(d).C 是等价的
	* 下列例子(任务超时),也同样说明了select中default的作用
	*/
	start := time.Now()

	tick := time.Tick(time.Millisecond*200) //200ms 执行一次(1s内执行1000/200=5,5-1 = 4次,实际上只执行了4次,最后一次还没执行就结束了)
	boom := time.After(time.Second*1)       //1s    后结束(1000ms)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
			//任务结束的信号
		case <-boom:
			fmt.Println("BOOM!")
			fmt.Println(time.Since(start))
			return
			//在每200ms周期内都执行这个,满足200ms就执行一次其他任务
		default:
			fmt.Println("default")
			time.Sleep(5e7)
		}
	}
}

//1.简单超时模式,从通道ch中接收数据,但是最多等待1s(1e9 = 1s,约等于)
func sel2()  {
	start   := time.Now()
	timeout := make(chan bool,1) //注意:超时管道设置缓冲大小,避免协程死锁,确保超时的通道可以被垃圾回收
	ch      := make(chan int)
	//超时设置
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	//接收数据的协程任务
	go func() {
		time.Sleep(time.Second*2)
		ch <- 1
	}()

	//没有执行其他的default任务,直接等待至超时
	select{
	case v := <-ch:
		fmt.Println("get ch finished:",v)
		break
	case <-timeout:
		fmt.Println("time out")
		break
	}

	//有执行其他的default任务,执行其他的default任务至超时,并且超时时间不一定是规定的时间,当default跟timeout都符合时,select会随机选择一个case
	/*f : for {
		select{
		case v := <-ch:
			fmt.Println("get ch finished:",v)
			break f
		case <-timeout:
			fmt.Println("time out")
			break f
		default:
			fmt.Println("default task")
		}
	}*/
	fmt.Println(time.Since(start))
}

//2.取耗时很长的同步调用
/**
注意缓冲大小设置为 1 是必要的，可以避免协程死锁以及确保超时的通道可以被垃圾回收。
此外，需要注意在有多个 case 符合条件时， select 对 case 的选择是伪随机的，
如果上面的代码稍作修改如下，则 select 语句可能不会在定时器超时信号到来时立刻选中
time.After(timeoutNs) 对应的 case，因此协程可能不会严格按照定时器设置的时间结束
 */
func sel3()  {
	ch := make(chan int,1)

	//执行任务
	go func() {
		time.Sleep(time.Second*2)
		ch<-1
	}()

	select {
	case v := <-ch:
		fmt.Println("task finished",v)
	case <-time.After(time.Second*1):
		fmt.Println("after 1s time out")
		break
	}

/*L:
	for {
		select {
		case <-ch:
			// do something
		case <-time.After(timeoutNs):
			// call timed out
			break L
		}
	}*/
}

//第三种形式
func sel4()  {
	/**
	假设程序从多个复制的数据库同时读取,只需要一个答案,需要接收首先到达的答案,Query 函数获取数据的连接切片
	并请求,并行请求每一个数据库并返回收到的第一个响应
	func Query(conns []Conn,query string) Result {
		ch := make(chan Result,1)
		for _, conn := range conns {
			go func(c Conn) {
				select {
					case ch <- c.DoQuery(query):
					default:
				}
	 		}(conn)
		}

		//结果通道必须带缓冲,保证第一个发送进来的数据有位置可以存放,确保放入的首个数据总会成功
		//所以第一个到达的值会被获取跟执行的顺序无关,正在执行的协程总是可以使用runtime.Goexit()来停止
		return <-ch
	}

	在应用中缓存数据:
	应用程序中用到了来自数据库/常见的数据存储 的数据时,经常会把数据缓存到内存中,因为从数据库中获取数据的操作代价很高
	如果数据库中的值不发生变化,就没有问题,如果值有变化,我们需要一个机制来周期性的从数据库重新读取这些值:缓存的值
	就不可用/过期了
	 */
}

//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.6.md
func sel5()  {
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
	Mu sync.Mutex  "是互斥锁,用来在代码中保护临界区资源,同一时间只有一个go协程可以进入该区域"
	Tasks []*Task
}
func sel6()  {
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
	fmt.Println(t.a)
}

/**
其他模式:使用通道
一个通道接受需要处理的任务(队列),一个通道接受处理完成的任务
worker在协程中启动,其数量N应该根据任务数量进行调整
主线程扮演master节点角色
 */
func sel7()  {
	N := 5
	pending,done := make(chan *Task),make(chan *Task)
	go sendWork(pending,N)

	//开启多个协程同步(顺序接受任务,处理任务,接已完成的任务)
	for i := 0; i < N; i++ {
		go Worker2(pending,done)
	}

	for doneTask := range done  {
		fmt.Println(doneTask.a)
	}
}
func Worker2(in ,out chan *Task)  {
	for {
		//接受需要处理的任务
		t := <-in

		//开始完成任务
		t.process()

		//接受处理完成的任务
		out <- t
	}
}
func sendWork(t chan<- *Task,taskNum int)  {
	for i:= 0; i < taskNum; i++ {
		t <- new(Task)
	}
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
var ints chan int
func sel8()  {
	ints = yield()
	fmt.Println(generate())
	fmt.Println(generate())
	fmt.Println(generate())
}
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
type Any interface {}
type EvalFunc func(any Any) (Any,Any)
func sel9()  {
	//具体计算过程
	evenFunc := func(state Any) (Any,Any) {
		OS,ok := state.(int)
		ns    := 0
		if ok {
			ns = OS + 2
		}
		return OS,ns
	}
	event := sel9_2(evenFunc,0)

	//消费者
	for i := 0; i < 10; i++ {
		fmt.Println(event())
	}
}

//生产者
func sel9_1(e EvalFunc,a Any) func() Any {
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
func sel9_2(evalFunc EvalFunc,initState Any) func() int  {
	ef := sel9_1(evalFunc,initState)
	return func() int {
		return ef().(int)
	}
}

//惰性求值 练习1
var fibArr []uint64
func sel10()  {
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