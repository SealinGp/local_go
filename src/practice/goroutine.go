package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
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
	}
	if nil == funs[n] {
		fmt.Println("func",n,"unregistered")
		return
	}
	funs[n]()
}
var numCores = flag.Int("n",2,"usage")
func gor1()  {
	flag.Parse()
	//通过命令行指定使用的核心数量
	runtime.GOMAXPROCS(*numCores)
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
			res[i] = xi + 1
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

		//fmt.Println(prime)
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