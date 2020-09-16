package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.0.md
协程,通道

协程:一个应用程序运行在机器上的一个进程,进程是一个运行在自己内存地址空间里的独立执行体
1 进程 = 多个 线程(操作系统级)
多个线程共享同一个内存地址空间,一起工作.

竞态:使用多线程的应用难以做到准确,最主要的问题是内存中的数据共享,他们会被多线程无法预知的方式进行
操作,导致一些无法重现或者随机的结果

解决方法:同步不同的线程,对数据加锁,这样多线程中只有一个线程可以变更数据

程序级的协程 != 操作系统级线程

协程是根据一个/多个线程的可用性,执行于线程之上的,协程的涉及隐藏了许多线程创建和管理方面的复杂工作
协程是轻量级,比线程更轻量,使用更少的内存和资源

https://www.jianshu.com/p/aad2b27992eb
应用程序层:
栈:维护函数调用的上下文,在用户更高的地址空间处分配(简单来说就是放置函数的地方)
堆:容纳应用程序动态分配的内存区域,当程序使用new或malloc时,就是得到来自堆中的内存,在栈的下方,堆分配的内存比栈大一点
栈向低地址增长；堆向高地址增长
"segment fault"错误的来源:
在Linux或者是win内存中,有些地址是始终不能读写的,例如0地址,当指针指向这些地址的时候,就会出现“段错误(segment fault)"
1.程序员将指针初始化为NULL,但是没有赋予合理的初值就开始使用.
2.程序员没有初始化栈上的指针,指针的值一般是随机数,之后就开始使用该指针.

协程工作在相同的地址空间中,共享内存的方式是同步的,可以用sync包来实现,推荐使用channels来同步协程
协程可以运行在多个操作系统线程之间,也可以运行在线程之内,可以使用很小内存占用处理大量任务
由于操作系统线程上的协程时间片,可以使用少量线程就能拥有多个提供服务的协程,go运行时可以检测到
哪些协程被阻塞了,展示搁置并处理其他协程

并发方式:确定性的(明确定义排序),和非确定性的(加锁/互斥从而未定义排序,抢占式调度),go的协程和通道支持确定性的
任何go程序有main()函数都可以看做一个协程,尽管没用通过关键字go来启动,协程可以在程序初始化(init())的过程中运行

并发和并行的差异
go的并发原语提供了良好的并发设计基础,表达程序结构一边表示独立的执行的动作,所以go的重点不在于并行的首要位置
并发程序可能是并行的,也可能不是.并行是一种通过使用多处理器以提高速度的能力,但往往是,一个涉及良好的并发程序
在并行方面的表现也非常出色

环境变量 GOMAXPROCS
=1 时,所有的协程都会共享一个线程
>1 时,会有一个线程池管理许多的线程
假设n为机器核心数量/处理器的数量,若设置环境变量 GOMAXPROCS >= n,或者在代码中执行runtime.GOMAXPROCS(n),
那么协程会被分割(分散)到n个处理器上,更多的处理器 != 性能的线性提升,
有这样子的一个经验法则,对于n个核心的情况,
设置GOMAXPROCS = n-1 可以获得最佳性能,同时满足 协程的数量 > 1 + GOMAXPROCS > 1

GOMAXPROCS = 9 适用于1颗cpu的笔记本电脑
GOMAXPROCS = 8 适用于32核的机器上,更高的数值无法提升性能
总结: GOMAXPROCS 等同于(并发的)线程数量,在一台核心数>1的机器上,会尽可能有等同于核心数的线程在并行运行
*/
var (
	timezone string
	port     string
	numCores int
	fun      string
	NewYork  string
	Tokyo    string
	London   string
)
func init()  {
	flag.StringVar(&fun,"fun","","func you need to run")

	flag.StringVar(&timezone,"timezone","Asia/Shanghai","usage")
	flag.StringVar(&port,"port","","port")
	flag.StringVar(&NewYork,"NewYork","","address")
	flag.StringVar(&Tokyo,"Tokyo","","address")
	flag.StringVar(&London,"London","","address")
	flag.IntVar(&numCores,"numCores",2,"core")
}
func main() {
	flag.Parse()

	if fun == "" {
		flag.Usage()
		return
	}
	execute(fun)
}
func execute(n string) {
	funs := map[string]func(){
		"gor1" : gor1,
		"gor2" : gor2,
		"gor3" : gor3,
		"gor4" : gor4,
		"gor5" : gor5,
		"gor6" : gor6,
		"gor7" : gor7,
		"gor8" : gor8,
		"gor9" : gor9,
		"gor10" : gor10,
		"gor11" : gor11,
		"gor12" : gor12,
		"gor13" : gor13,
		"gor14" : gor14,
		"gor15" : gor15,
		"gor16" : gor16,
		"gor17" : gor17,
		"gor18" : gor18,
		"gor19" : gor19,
		"gor20" : gor20,
		"gor21" : gor21,
		"gor22" : gor22,
		"gor23" : gor23,
		"gor24" : gor24,
		"gor25" : gor25,
		"gor_test" : gor_test,
	}
	defer func() {
		if e := recover();e != nil {
			log.Println(e)
		}
	}()
	funs[n]()
}
func gor1()  {
	log.Println(numCores)
	//通过命令行指定使用的核心数量
	runtime.GOMAXPROCS(numCores)
}

//并行只用了10s,若不使用go关键字,串行会耗时10+5+2 = 17s
func gor2()  {
	fmt.Println("In main")
	go longWait()
	go shortWait()
	time.Sleep(10*time.Second)
	fmt.Println("end main")
}
func longWait()  {
	fmt.Println("Beginning long wait")
	time.Sleep(5*time.Second)
	fmt.Println("end long wait")
}
func shortWait()  {
	fmt.Println("Beginning shrot wait")
	time.Sleep(2*time.Second)
	fmt.Println("end short wait")
}

/*
使用channel在协程之间通信 = 通过通信来共享内存
通道实际上是类型化消息的队列,使数据得以传输,先进先出的结构(排队结构),
并且也是引用类型,需使用Make来给他分配内存
所有数据类型都可以用来声明管道,包括interface{},slice,array,map,func(),struct
*/
func gor3()  {
	ch := make(chan string)
	/*
	执行顺序为:
	1.goroutine gor3_1 : c <- "a"
	2.goruntine gor3_2 : i = <-c

	3.goroutine gor3_1 : fmt.Println("abc")

	4.goruntine gor3_1 : c <- "b"
	5.goruntine gor3_2 : i = <- c
	...
	*/
	go gor3_1(ch)
	go gor3_2(ch)

	time.Sleep(time.Second)
}
func gor3_1(c chan<- string)  {
	c <- "a"
	fmt.Println("abc")
	c <- "b"
	c <- "c"
	c <- "d"
	c <- "e"
	c <- "f"
}
func gor3_2(c <-chan string)  {
	var i string
	for {
		i = <-c
		fmt.Println(i)
	}
}
func gor4()  {
	t := time.Now()
	a := make(chan string)

	go gor4_1(a)
	go gor4_2(a)

	//从通道接收,等待直到管道a中有内容
	b := <-a
	c := <-a

	fmt.Println(b,c)
	fmt.Println("耗时",time.Since(t).String())
}
//发送至通道
func gor4_1(c chan<- string)  {
	//模拟任务处理时间
	time.Sleep(3*time.Second)

	c <- "func gor3_1"
}
func gor4_2(c chan<- string)  {
	//模拟任务处理时间
	time.Sleep(2*time.Second)

	c <- "func gor3_2"
}
func gor5()  {
	p := make(chan int)
	go gor5_1(p)
	p <- 2
	fmt.Println("end gor5")
}
func gor5_1(c <-chan int)  {
	c1 := <-c
	fmt.Println(c1)
}

//同步通道,带缓冲的通道
func gor6()  {
	/*
	有缓冲管道
	通道可以同时容纳的元素(这里是指string)的个数
	在缓冲100全部被使用之前,该管道不会阻塞
	总结:同时允许多少个协程同时对管道进行操作(协程并行数量限制)

	无缓冲管道:
	对于同一个通道,发送操作（协程或者函数中的）,在接收者准备好之前是阻塞的
	简单来说: 接收操作->发送操作 的顺序
	buf = 0时,
	执行到这一句ch1 <- "a" 导致panic
	*/
	buf := 4
	ch1 := make(chan string,buf)

	for i := 0; i< 5; i++ {
		go gor6_1(ch1,i)
	}

	for i := 0; i< 5; i++ {
		fmt.Println(<-ch1,i)
	}
	fmt.Println("end")
}
func gor6_1(c chan<- string,i int)  {
	c <- "gor6_1 " + string(i)
	fmt.Println(i)
}
/*
信号量模式
*/
func gor7()  {
	N    := 5
	data := make([]float64,N)
	res  := make([]float64,N)
	sem  := make(chan float64,N)

	for i,xi := range data  {
		go func(i int, xi float64) {
			res[i] = float64(i)
			sem <- res[i]
		}(i,xi)
	}

	for i := 0; i < N ; i++ {
		fmt.Println(<-sem)
	}
}

/*
信号量是实现互斥锁的常见同步机制,限制对资源的访问,解决读写问题
通过信号量来实现互斥锁的例子
互斥锁:防止多条线程对同一个变量进行读写的机制
*/
type semaphore chan interface{}
//write
func (s semaphore)w(n int)  {
	e := new(interface{})
	for i := 0; i < n ; i++ {
		s <- e
	}
}
//read
func (s semaphore)r(n int)  {
	for i := 0; i < n ; i++ {
		<-s
	}
}
func (s semaphore)Lock()  {
	s.w(1)
}
func (s semaphore)Unlock()  {
	s.r(1)
}
func (s semaphore)Wait(n int)  {
	s.w(n)
}
func (s semaphore)Signal()  {
	s.r(1)
}
func gor_test()  {
	c := container{
		items:[]item{"a1","a2","a3"},
	}
	ifdone := make(chan bool)

	go c.consume(ifdone)
	<-ifdone
}
func (c container)product()  <-chan item {
	ch := make(chan item)
	go func() {
		for _,it := range c.items {
			ch <- it
		}
		close(ch)
	}()
	return ch
}

func (c container)consume(done chan<- bool)  {
	for it := range c.product() {
		log.Println(it)
	}
	done <- true
}



func gor8()  {
	c    := make(chan int)
	done := make(chan bool)

	go gor8_1(c,10,10)
	go gor8_2(c,done)

	<-done
}
func gor8_1(c chan<- int,num,step int)  {
	ns := make([]int,num)
	for i := range ns {
		c <- i*step
	}

	/*
	https://blog.csdn.net/butterfly5211314/article/details/81842519
	close 函数是一个内建函数,用来关闭channel,这个channel必须为发送者
	当最后一个发送的值,被接受者从关闭的channel中接收时,接下来所有接收的值都会非阻塞直接成功,返回
	channel元素的0值
	*/
	close(c)
}
func gor8_2(c <-chan int,done chan<- bool)  {
	/*
      它从指定通道中读取数据直到通道关闭，才继续执行下边的代码
	  使用for-range语句来读取通道是更好的办法,因为会自动检测通道是否关闭
	*/
	for n := range c {
		fmt.Println(n)
	}

	//当c所有的值都被接收了(即通道关闭了),则 ok = false
	k,ok := <-c
	fmt.Println(k,ok)
	done <- true
}

//通道工厂模式
func gor9()  {
	ch := make(chan bool)

	stream := gor9_1(3)
	go gor9_2(stream,ch)

	//go gor9_2(gor9_1(3),ch)

	<-ch
}
func gor9_1(n int) chan int  {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++  {
			fmt.Println("write",i,"start")
			ch <- i
			fmt.Println("write",i,"end")
		}

		//一次向通道内写入超过通道缓存的数量,那么在写入完后需要关闭通道
		close(ch)
	}()
	return ch
}
func gor9_2(ch chan int,c chan<- bool)  {
	for {
		fmt.Println("read 1")
		r,ok := <-ch
		if !ok {
			c <- true
			break
		}
		fmt.Println(r)
		fmt.Println("read 2")
	}
}

//通道迭代模式 = 生产者-消费者模式
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md
type item string
type container struct {
	items []item
}
func gor10()  {
	//确定是否读完
	done := make(chan bool)
	con := &container{
		items:[]item{"abc","def"},
	}
	go gor10_2(con,done)

	<-done
}
//生产者
func (c *container)gor10_1() <- chan item  {
	ch := make(chan item)
	go func() {
		for i := 0; i < len(c.items) ; i++ {
			ch <- c.items[i]
		}
		close(ch)
	}()
	return ch
}
//消费者
func gor10_2(con *container,d chan<- bool)  {
	for x := range con.gor10_1()  {
		fmt.Println(x)
	}

	d <- true
}

/*
管道和选择器模式
例子:筛选
*/
func gor11()  {
	//main goroutine
	ch := make(chan int)
	t := time.Now()

	//write goroutine1
	go gor11_write(ch)
	for {
		//1: 停住
		//3: 继续 ch=2 prime=2

		//4: 停住
		prime := <- ch

		//fmt.Println(prime)plugplug
		ch1 := make(chan int)

		//3: 开启一个 filter goroutine ch=0 ch1=0 prime=2
		go gor11_filter(ch,ch1,prime)

		//3: ch1=0 ch=0
		ch = ch1

		//3: false
		//周期后结束
		if time.Since(t) > time.Millisecond*1 {
			return
		}
	}

}
func gor11_write(ch chan<- int)  {
	for i:= 2; ; i++ {
		//2: i=2 ch=2
		//5: i=3 ch=3
		ch <- i
	}

}
func gor11_filter(ch <-chan int,ch1 chan<- int,prime int)  {
	for {
		//4:停住 ch=0 ch1=0 prime=2
		//6: 继续 ch=3 i=3

		//7: 停住 ch = 0
		i := <-ch

		//6: prime=2 i=3
		fmt.Println(prime,i)

		//6: 3%2 = 1 ch1 = 3
		if i % prime != 0 {
			ch1 <- i
		}
	}
}

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.3.md
协程同步,管道关闭,一般使用defer关闭

https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.4.md
使用select切换协程
select监听进入通道的数据,也可以是用通道发送值的时候
fallthrough是不允许的

select作用:
选择处理列出的多个通信情况中的一个
如果都阻塞了,会等待知道其中一个可以处理
如果没阻塞,那么随机选择一个
如果没有通道操作可以处理了(都关闭了),并且写了default语句,那么执行default语句
若在select中使用发送操作(chan1 <- xxx),有default可以确保发送不被阻塞,如果没有default,select就会一直阻塞
*/

//生产者-消费者,select例子
func gor12() {
	ch1,ch2  := make(chan int),make(chan int)
	go gor12_write(ch1)
	go gor12_write1(ch2)
	go gor12_read(ch1,ch2)
	time.Sleep(time.Second*1)
}
func gor12_write(ch chan<- int)  {
    for i := 0; ; i++ {
       ch <- i * 10
	}
}
func gor12_write1(ch chan<- int)  {
	for i := 0; ; i++ {
		ch <- i * 100
	}
}
func gor12_read(ch1,ch2 <-chan int)  {
    for {
		select {
		case v := <-ch1:
			fmt.Println(v,"ch1")
		case v := <-ch2:
			fmt.Println(v,"ch2")
	   }
	}
}

func gor13()  {
	//练习5.4
	n := 15
	for i := 0;i < n;i++ {
		fmt.Println(i)
	}

	i := 0
	FOR:
		if i < n {
			fmt.Println(i)
			i++
			goto FOR
		}
}
func gor14()  {
	done := make(chan bool)
	ch := make(chan int)
	go tel(ch,5,done)

	//方法1
	//for a := range ch  {
	//	fmt.Println(a)
	//}

	//方法2
	//for {
	//	i,ok := <- ch
	//	if !ok {
	//		break;
	//	}
	//	fmt.Println(i)
	//}

	//方法3
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-done:
			fmt.Println(v)
			return
		}
	}
}
func tel(ch chan<- int,num int,d chan<- bool)  {
	for i := 0; i < num; i++ {
		ch <- i
	}
	d <- true

	//方法2|方法1 需要打开
	//close(ch)
}

//练习14.8 fib使用通道channel,速度最快
func gor15()  {
	term := 25
	i    := 0
	c    := make(chan int)

	go gor15_1(term,c)
	start := time.Now()
	for {
		if result, ok := <-c ; ok {
			fmt.Println(i,result)
			i++
		} else {
			fmt.Println(time.Since(start))
			return
		}
	}
}
func gor15_1(l int,c chan int)  {
	f := gor15_2()
	for i := 0; i < l; i++ {
		a := i
		if i <= 2 {
			a = 1
		}
		c <- f(a)
	}
	close(c)
}
func gor15_2() (func(int) int) {
	var red []int
	return func(i int) int {
		if l := len(red); l >= 2 {
			i = red[l-1] + red[l-2]
		}
		red = append(red,i)
		return red[len(red) - 1]
	}
}

//练习14.8的改良版(两者速度差不多,但内存肯定是这个用的少,因为少用了一个缓冲数组)
func gor16()  {
	start := time.Now()
	le    := 25
	done  := make(chan int)
	go gor16_1(le,done)
	i := 0
	for d := range done  {
		fmt.Println(i,d)
		i++
	}

	fmt.Println(time.Since(start))
}
func gor16_1(n int,c chan int)  {
	x,y := 1,1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//练习14.8.1 fib使用通道channel 和 select 让通道退出
func gor17()  {
	start := time.Now()
	ch := make(chan int)
	done := make(chan bool)
	le := 25

	go gor17_1(le,ch,done)

	i := 0
	for {
		select {
		case v := <- ch:
			fmt.Println(i,v)
			i++
		case v := <- done:
			fmt.Println(v)
			fmt.Println(time.Since(start))
			return
		}
	}
}
func gor17_1(n int,c chan int,done chan bool)  {
	x, y := 1,1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	done <- true
}

//2个协程????
func gor18()  {
	start := time.Now()
	x     := gor18_1()
	le    := 25
	for i := 0; i < le; i++ {
		fmt.Println(<-x)
	}

	fmt.Println(time.Since(start))
}
func gor18_1() (<-chan int) {
	d := make(chan int,2)
	a,b,out := gor18_2(d)
	go func() {
		d<- 0
		d<- 1

		<-a
		for {
			d <- (<-a + <-b)
		}
	}()
	<-out
	return out
}
func gor18_2(in <-chan int) (<-chan int,<-chan int,<-chan int) {
	a, b, c := make(chan int,2),make(chan int,2),make(chan int,2)
	go func() {
		for {
			o := <-in // x = 0
			a <- o
			b <- o
			c <- o
		}
	}()
	return a,b,c
}

//提供总数量为num的 随机 0|1 的协程
func gor19()  {
	start    := time.Now()
	num := 25
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c," ")
		}
	}()

	for {
		if num <= 0 {
			fmt.Println(time.Since(start))
			return
		}
		num--
		select {
		case c<-0:
		case c<-1:
		}
	}
}

//练习14.10 直角坐标系又叫笛卡尔坐标系 channel1 极坐标,channel 笛卡尔坐标

//极坐标
type JI struct {
	le     float64 "OP长度,极径"
	corner float64 "OP与OX的夹角,极角.单位:度数"
}
type ZJ struct {
	x float64  "X轴长度"
	y float64  "Y轴长度"
}
func gor20()  {
	ji        := new(JI)
	ji.le     = 3
	ji.corner = 60
	channel1 := make(chan *JI) //极坐标
	channel2 := make(chan *ZJ) //笛卡尔坐标 [x,y]
	go func() {
		//读取极坐标
		ji1  := <-channel1

		//计算笛卡尔坐标
		zji1 := new(ZJ)
		zji1.calculate(ji1)
		channel2<-zji1
	}()
	//写入极坐标
	channel1<- ji

	//读取笛卡尔坐标
	zj := <-channel2
	fmt.Println(zj.x,zj.y)
}
func (z *ZJ)calculate(J *JI) {
	z.y = J.le * math.Sin(J.corner)
	z.x = math.Sqrt(J.le*J.le - z.y*z.y)
}

//生产者-消费者模型
//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-06-goroutine.md
func gor21()  {
	done := make(chan bool)
	taskCh := make(chan int,10)

	go gor21_producer(taskCh,2)
	go gor21_consumer(taskCh,done)

	<-done
}
func gor21_producer(task chan<- int,beishu int)  {
	defer close(task)
	for i := 0; i<10; i++ {
		task <- i*beishu
	}
}
func gor21_consumer(task <-chan int,done chan<- bool)  {
	for t := range task  {
		index := t
		fmt.Println(index)
	}
	done <- true
}

/**
发布-订阅模型(published-subscribe = pub-sub)
https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-06-goroutine.md
生产者 = 发布者
消费者 = 订阅者
传统生产者-消费者模型中,将消息发送到一个队列中,发布者-订阅者,将消息发布给一个主题
 */
type (
	sub       chan interface{}         //订阅者
	topicFunc func(v interface{}) bool //主题

	Publisher struct {                //发布者
		m sync.RWMutex                //读写锁
		buffer int                    //订阅队列的缓存大小
		timeout time.Duration         //发布超时时间
		subscribers map[sub]topicFunc //订阅者信息
	}
)
func gor22()  {
	p     := NewPublisher(100*time.Millisecond,10)
	done  := make(chan bool)
	done1 := make(chan bool)

	all       := p.Subscribe()
	golangTop := p.SubscribeTopic(func(v interface{}) bool {
		if s,ok := v.(string);ok {
			return strings.Contains(s,"golang")
		}
		return false
	})

	p.Publish("hello,world!")
	p.Publish("hello,golang!")
	p.Close()

	go func() {
		for msg := range all {
			fmt.Println("all:",msg)
		}
		done <- true

	}()

	go func() {
		for msg := range golangTop {
			fmt.Println("golang:",msg)
		}
		done1<- true
	}()

	<-done1
	<-done
}
func NewPublisher(timeOut time.Duration,buf int) *Publisher {
	return &Publisher{
		buffer:buf,
		timeout:timeOut,
		subscribers:make(map[sub]topicFunc),
	}
}

//添加一个订阅者,订阅指定主题
func (p *Publisher)SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{},p.buffer)

	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[ch] = topic
	return ch
}
//添加一个订阅者,订阅所有主题
func (p *Publisher)Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

//退出订阅
func (p *Publisher)Evict(sub chan interface{})  {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers,sub)
	close(sub)
}

//发布一个主题,允许一定时间内的超时
func (p *Publisher)sendTopic(su sub,topic topicFunc,v interface{}, wg *sync.WaitGroup)  {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case su <- v:
		fmt.Println("sendTopic success!")
	case <-time.After(p.timeout):
		fmt.Println("sendTopic time out!")
	}
}
func (p *Publisher)Publish(v interface{})  {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	for sub,top := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub,top,v,&wg)
	}

	wg.Wait()
}

//关闭发布者,同时关闭订阅者管道
func (p *Publisher)Close()  {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers,sub)
		close(sub)
	}
}

// http://dev.api.com/_book/ch8/ch8-02.html
func gor23()  {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\r Fibonacci(%d) = %d\n",n,fibN)
}
func fib(x int) int  {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func spinner(d time.Duration)  {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c",r)
			time.Sleep(d)
		}
	}
}

// http://dev.api.com/_book/ch8/ch8-02.html
//clock2
func gor24()  {
	clock2()
}
func clock2()  {
	listener, err := net.Listen("tcp","localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}
	lo,_  := time.LoadLocation(timezone)
	fmt.Println("in:",timezone,"listen:",port,"now:",time.Now().In(lo).Format("15:04:05\n"))

	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			for {
				_, err := io.WriteString(c,time.Now().In(lo).Format("15:04:05\n"))
				if err != nil {
					return
				}
				time.Sleep(time.Second)
			}
		}(conn)
	}
}

//clockwall
type wallStr struct {
	ch chan []byte
}

var ws  = NewWallStr(3)
func gor25()  {
	var wg sync.WaitGroup

	if NewYork != "" {
		wg.Add(1)
		go clockwall(NewYork,wg,[]byte("NewYork:["))
	}
	if Tokyo != "" {
		wg.Add(1)
		go clockwall(Tokyo,wg,[]byte("Tokyo:["))
	}
	if London != "" {
		wg.Add(1)
		go clockwall(London,wg,[]byte("London:["))
	}

	go func() {
		i := 0
		for {
			i++
			a := ws.Read()
			if len(a) > 0 {
				a = append(a,' ')
				if i == 3 {
					a = append(a,'\r')
					i = 0
				}
				os.Stdout.Write(a)
			}
		}
	}()

	wg.Wait()
}
func clockwall(address string,wg sync.WaitGroup,tz []byte)  {
	defer wg.Done()

	c,e := net.Dial("tcp",address)
	if e != nil {
		log.Fatal(e)
	}
	defer  c.Close()


	mustCopy(ws,c,tz)
}
func mustCopy(w io.Writer,r io.Reader,tz []byte)  {
	//if _,e := io.Copy(w,r); e != nil {
	//	log.Fatal(e)
	//}

	p := make([]byte,1024)
	for {
		i,e := r.Read(p)

		if e != nil {
			break
		}

		all := tz
		all = append(all,p[0:i-1]...)
		all = append(all,']')
		if _,e = w.Write(all);e != nil {
			break
		}

	}
}
func NewWallStr(num int) *wallStr {
	return &wallStr{ch:make(chan []byte,num)}
}
func (ws wallStr)Read() []byte {
	var res []byte
	select {
	case res = <-ws.ch:
	case <-time.Tick(time.Second*5):
	}
	return res
}
func (ws wallStr)Write(s []byte) (int,error) {
	ws.ch <- s
	return len(s),nil
}

/**
https://gocn.vip/topics/1569
实参,形参
实参: 实参其实就是将参数的值复制到函数里来(参数在函数内外,地址是不同的)
形参: 形参其实就是将参数的地址传递到函数里来(参数在函数内外,地址是相同的)
切片和映射只有实参,并且不用加*号

https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
1.指针变量存储的值是变量地址,自身也有一个地址
2.函数调用的时候,需要在当前协程上面的堆栈分配新的内存,2个函数帧之间会有一个转换(值传递还是指针传递(贡献变量地址),函数帧内逃逸分析,协程内有效内存与无效内存)
*/
func gor26()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(12)
	//将变量放在堆还是栈上,是由go的逃逸分析机制决定的

	//形参
	//1.会将变量移动到heap中,因为两个goroutine都要用到该变量
	//2.并发安全性,外部变量的改变会影响内部变量的改变
	for i := 0;  i < 6; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("T1:",i)
		}()
	}

	//实参
	//1.会将i放到groutine的stack上
	//2.外部变量的改变不会影响内部变量的改变
	for i := 0;  i < 6; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("T2:",i)
		}(i)
	}
	wg.Wait()
}

/**
https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/
linux调度器并不区分线程和进程,在大部分操作系统的实现中,线程属于进程
多个线程
1MB = 1024 KB
1KB = 1024 Byte
1 Byte = 8 bit

go调度器的进化过程
1.单线程调度器 0.x 版本
程序中只能存在一个活跃线程,由GM模型组成
c语言实现
->获取调度器全局锁
->调用gosave保存栈寄存器和程序计数器
->调用nextgandunlock获取下一个需要运行的goroutine并解锁调度器
->修改全局线程m上要执行的goroutine
->调用gogo函数运行最新的goroutine

2.多线程调度器 1.0
允许运行多线程的程序
全局锁导致竞争严重

3.任务窃取调度器 1.1
引入了处理器P,构成了G-M-P模型
在处理器P的基础上实现了基于工作窃取的调度器
在某些情况下,goroutine不会让出线程,造成饥饿问题
时间过长的垃圾回收(Stop-the-word,STW) 会导致程序长时间无法工作


4.抢占式调度器 1.2~now
  基于协作的抢占式调度器 1.2~1.13
     通过编译器在函数调用时插入抢占检查指令,在函数调用时当前goroutine是否发起了抢占请求,实现基于协作的抢占式调度
     goroutine可能会因为垃圾回收和循环长时间占用资源导致程序暂停
  基于信号的抢占式调度器 1.14~now
     实现基于信号的真抢占式调度
     垃圾回收在扫描栈时会触发抢占式调度
     抢占的时间点不够多,还不能覆盖全部的边缘情况
->编译器会在调用函数钱插入runtime.morestack
->go 语言运行时会在垃圾回收暂停程序,系统监控发现goroutine运行超过10ms时发出抢占请求StackPreempt
->当发生函数调用时,可能会执行编译器插入的runtime.morestack函数,他调用的runtime.newstack会检查goroutine的stackguard0字段是否为StackPreempt
->如果stackguard0是stackPreempt,就会触发抢占让出当前线程

5.非均匀存储访问调度器
  对运行时的各种资源进行分区
  实现非常复杂,到今天还没有提上日程



G(goroutine)-M(thread)-P(cpu) 模型
G: Go语言调度器中待执行的任务,在运行时在调度器中的地位 = 线程在操作系统中的地位,占用了更小的内存空间
//当前只显示部分成员
type g struct {
   stack stack            //栈内存范围
   stackguard0  uintptr   //用于调度器抢占式调度
   _panic       *panic    //存储panic的链表
   _defer       *_defer   //存储defer函数的链表
   m            *m        //当前goroutine占用的线程
   atomicstatus uint32    //goroutine状态
   sched        gobuf     //存储goroutine调度相关的数据
   goid         int64     //goroutine的id,该字段对开发者不可见
   ...
}

//这些内容会在调度器保存或恢复上下文的时候用到,其中栈指针和程序计数器会用来存储或者恢复寄存器中的值,改变程序即将执行的代码
type gobuf struct {
    sp uintptr  //栈指针
    pc uintprt  //程序计数器
    g  guintptr //持有gobuf的goroutine
    ret sys.Uintreg //系统调用的返回值
    ...
}

goroutine状态(常见的):
_Grunnable 没有执行代码,没有栈的所有权,存储在运行队列中
_Grunning  可以执行代码,有栈的所有权,被赋予了内核线程M和处理器P
_Gsyscall  正在执行系统调用,拥有栈的所有权,没有执行用户代码,被赋予了内核线程M但不在运行队列上
_Gwaiting  由于运行时被阻塞,没有执行用户代码并且不在运行队列上,但是可能存在于Channel的等待队列上
_Gpreempted 由于抢占而被阻塞,没有执行用户代码并且不在运行队列上,等待唤醒

M:操作系统线程,调度器最多可以创建1w个线程,但是其中大多数线程都不会执行用户代码(可能陷入系统调用),最多只会有GOMAXPROCS个活跃线程能够正常运行
默认情况下,一个4核机器上会创建4个活跃的操作系统线程,每一个线程都对应一个运行时中的runtime.m结构体
type m struct {
   g0   *g //持有调度栈的goroutine,它会深度参与运行时的调度过程,包括goroutine的创建,大内存分配和CGO函数的执行
   curg *g //当前线程上运行的用户goroutine
   p    puintptr //正在运行代码的处理器p
   nextp puintptr //暂存的处理器nextp
   oldp puintptr //执行系统调用之前的使用线程的处理器oldp
   ...
}

P:线程和goroutine的中间层,提供线程需要的上下文环境,负责调度线程上的等待队列,通过处理器P的调度,每一个内核线程都能执行多个goroutine
它能在goroutine进行一些I/O操作时及时切换,提高线程利用率
type p struct {
   m muintptr //

   runqhead uint32 //处理器持有的运行队列
   runqtail uint32
   runq     [256]guintptr

   runnext  guintptr //线程下一个需要执行的goroutine
   ...
}



*/