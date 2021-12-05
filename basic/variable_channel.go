package main

import (
	"fmt"
	"strconv"
	"time"
)

var channelFuns = map[string]func(){
	"channel1":    channel1,
	"channel2":    channel2,
	"channel3":    channel3,
	"channel4":    channel4,
	"channel5":    channel5,
	"channel6":    channel6,
	"channel7":    channel7,
	"channel7_1":  channel7_1,
	"channel7_2":  channel7_2,
	"channel8":    channel8,
	"channel8_1":  channel8_1,
	"channel8_2":  channel8_2,
	"channel9":    channel9,
	"channel10":   channel10,
	"channel11":   channel11,
	"channel12":   channel12,
	"channel13":   channel13,
	"channel13_1": channel13_1,
	"channel13_2": channel13_2,
}

//refurl: https://colobu.com/2016/04/14/Golang-Channels/

/*
 作用:协程之间通信的方式
 定义方式
    双向管道:
	var ch1 chan int;    				//双向管道
	单向管道:主要用于函数的输入参数声明,说明该函数对管道的操作
	var ch1 chan<- int;  				//单向写
	var ch1 <-chan int;  				//单向读

	ch2 := make(chan int,capacity int); //capacity 容量/缓存
	容量设置的时候,当容量满的时候才会发生blocking(阻塞)
*/

//buffered channels 缓存管道,可以避免阻塞
func channel1() {
	ch := make(chan int, 100)

	for i := 3; i < 5; i++ {
		ch <- i
	}

	v1 := make([]int, 0)
	v2 := make([]int, 2)
	for i := 0; i < 2; i++ {
		cValue := <-ch
		fmt.Println(cValue)

		v1 = append(v1, cValue)
		v2[i] = cValue + 1
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(append(v2, v1...))
}

//管道必须在写入方关闭
func channel2() {
	c := make(chan int)

	//创建协程并调用
	go func() {
		c <- 3 + 4
		close(c)
	}()

	i := <-c
	fmt.Println(i)
}

//blocking 阻塞(平行goroutine(协程) 的执行顺序是不可预测的)
func channel3() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	//入栈7+2+8=17
	go channle3_sum(s[0:len(s)/2], c, 0)
	// go channle3_sum(s[:len(s)/2],c);

	//入栈-9+4+0=-5
	go channle3_sum(s[len(s)/2:len(s)], c, 1)
	// go channle3_sum(s[len(s)/2:],c);

	//出栈 -5,17 等待管道写入数据后(等待协程执行结束),读取出来
	x, y := <-c, <-c
	fmt.Println(x, y)
}
func channle3_sum(a []int, c chan int, index int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(index)
	c <- sum
}

//range 处理channel
func channel4() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()

	c := make(chan int)

	go func() {
		//若在管道写入完毕后不关闭管道,则程序会一直阻塞在for..range
		defer close(c)

		for i := 0; i < 10; i++ {
			c <- i
		}

		// //若在管道写入完毕后不关闭管道,则程序会一直阻塞在for..range
		// close(c);
	}()

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Finished")
}

//select 处理channel
func channel5() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//c读
			fmt.Println(<-c)
		}
		//quit写
		quit <- 0
	}()

	channel5_fibonacci(c, quit)
}

func channel5_fibonacci(c, quit chan int) {
	x, y := 0, 1

	//死循环等待协程写入quit数据后读取跳出循环
	for {
		select {
		//i=0,x=0,y=1    0
		//i=1,x=1,y=1    1
		//i=2,x=1,y=2    1
		//i=3,x=2,y=3    2
		//i=4,x=3,y=5    3
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*

带超时机制
timeout
time.After(t int) 在时间t后返回一个单向可读的channel
*/
func channel6() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case t1 := <-time.After(time.Second * 1):
		fmt.Println("timeout 1", t1)
	}
}

/*
超时管道
time.NewTimer和ticker
time.NewTimer(t int):定时器,在t时间后返回一个单向读时间channel
*/
func channel7() {
	//2s后返回单向可读的时间channel
	timer1 := time.NewTimer(time.Second * 2)

	fmt.Println("now ", time.Now().Format("2006-01-02 15:04:05"))

	//阻塞2s
	time2 := <-timer1.C

	fmt.Println("2 second later ", time2.Format("2006-01-02 15:04:05"))
}

/*
超时管道
提前停止
time.NewTimer(t int)
*/
func channel7_1() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("now ", time.Now().Format("2006-01-02 15:04:05"))

	//开启一个协程
	go func() {
		time1 := <-timer1.C
		fmt.Println("2s later timer1 ", time1.Format("2006-01-02 15:04:05"))
	}()

	//未等待协程/定时器完毕就中途停止了
	stop1 := timer1.Stop()
	if stop1 {
		fmt.Println("timer1 stoped ", time.Now().Format("2006-01-02 15:04:05"))
	}
}

/*
time.NewTicker(t int):计时器,每过t时间,就向channel发送一个时间,
channel的接受者可以以固定的时间间隔从channel中读取事件
*/
func channel7_2() {
	//每500ms发送一个时间到channel
	ticker := time.NewTicker(time.Millisecond * 500)
	now("start at ", false)

	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()

	now("ended at ", false)
}

//管道关闭后,可读取数据(缓冲区读取完毕后一直读0),但不能写入数据(报panic错误),
func channel8() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

//通过循环读取已关闭的管道,缓冲区读取完毕后,会跳出循环
func channel8_1() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	for i := range c {
		fmt.Println(i)
	}
}

//通过i,ok := <-c 查询channel的状态,判断值是0值还是正常读取的值
func channel8_2() {
	c := make(chan int, 10)
	close(c)

	i, ok := <-c
	fmt.Println(i, ok)
}

//channel 可以在gorountine(协程)之间同步(控制协程)
func channel9() {
	defer now("end ", false)
	now("start ", false)

	ch := make(chan bool, 1)

	go channel9_work(ch, 1)
	go channel9_work(ch, 2)

	//等待协程将管道写入完毕
	for c := range ch {
		fmt.Println(c)
	}
}
func channel9_work(ch1 chan bool, index int) {
	//协程开始任务
	fmt.Println(index)
	time.Sleep(time.Second)

	//完成任务,管道写入数据
	ch1 <- true
}

var formatTime string = "2006-01-02 15:04:05"

func now(msg string, delTime bool) {
	if !delTime {
		msg = msg + time.Now().Format(formatTime)
	}
	fmt.Println(msg)
}

/*
golang中的并发限制跟超时控制
refUrl:https://juejin.im/entry/5a7aaac26fb9a0634a38fce2
*/

/*
无缓冲管道(通过内存共享)控制并发简单的例子
任务无序完成(不会按照创建任务的顺序完成)
*/
func channel10() {
	input := []int{3, 2, 1}
	ch := make(chan string)
	startTime := time.Now()
	now("multi tasks start,totalNum:"+fmt.Sprintf("%d", len(input))+"\n", false)

	for taskId, sleeptime := range input {
		go channel10_run(taskId, sleeptime, ch)
	}

	//管道读取
	for range input {
		fmt.Println(<-ch)
	}

	now("muti tasks ended. processTime:"+fmt.Sprintf("%s", time.Since(startTime))+"\n", false)
}

func channel10_run(task_id, sleepTime int, ch chan<- string) {
	//任务开始f
	time.Sleep(time.Duration(sleepTime) * time.Second)

	//任务结束,写入管道
	ch <- fmt.Sprintf("task id %d , sleep %d s", task_id, sleepTime)
}

/*
有缓冲管道(通过内存共享)控制并发简单的例子
控制任务有序完成
*/
func channel11() {
	tasks := []int{3, 2, 1}

	//带长度为5的数组中有3个管道
	chs := make([]chan string, len(tasks))
	startTime := time.Now()
	len := strconv.Itoa(len(tasks))
	now("multi tasks start,totalNum:"+len+"\n", false)

	for taskId, sleeptime := range tasks {
		chs[taskId] = make(chan string)
		go channel11_run(taskId, sleeptime, chs[taskId])
	}

	//管道读取
	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	processTime := time.Since(startTime).String()
	now("muti tasks ended. processTime:"+processTime+"\n", false)
}
func channel11_run(task_id, sleepTime int, ch chan<- string) {
	//任务开始f
	time.Sleep(time.Duration(sleepTime) * time.Second)

	//任务结束,写入管道
	ch <- fmt.Sprintf("task id %d , sleep %d s", task_id, sleepTime)
}

/*
超时控制
若某个goruntine(协程)运行时间太长,会拖累其他goruntine(协程),
因此需要超时控制
*/
func channel12() {

	tasks := []int{3, 2, 1} //协程运行时间3s,2s,1s,因此会有1/2个超时的协程
	timeout := 2
	chs := make([]chan string, len(tasks))
	len := strconv.Itoa(len(tasks))
	startTime := time.Now()
	now("multi tasks start,totalNum:"+len+"\n", false)

	for task_id, sleeptime := range tasks {
		chs[task_id] = make(chan string)
		go channel12_Run(task_id, sleeptime, timeout, chs[task_id])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	processTime := time.Since(startTime).String()
	now("muti tasks ended. processTime:"+processTime+"\n", false)
}

/*
通过select 和 time.After来判断超时(是指协程在2s后仍然没有完成任务)
sleeptime 协程耗时长
timeout   设定的超时时间
*/
func channel12_Run(task_id, sleeptime, timeout int, ch chan<- string) {
	ch_run := make(<-chan string)
	go channel12_run(task_id, sleeptime, ch)

	timeOut := time.After(time.Duration(timeout) * time.Second)
	select {
	case re := <-ch_run:
		ch <- re
	case <-timeOut:
		re := fmt.Sprintf("task id %d, timeout", task_id)
		ch <- re
	}
}

//协程执行耗时
func channel12_run(task_id, sleeptime int, ch chan<- string) {
	//假设执行时间
	time.Sleep(time.Duration(sleeptime) * time.Second)

	ch <- fmt.Sprintf("task id %d, sleep %d s", task_id, sleeptime)
	return
}

/*
并发限制(通过管道缓冲限制)
*/

//无缓冲管道,报错
func channel13() {
	ch := make(chan string)

	//这个goruntine被阻塞了,后面的程序没有执行
	ch <- "123"

	now(<-ch, true)
}

//有缓冲管道
func channel13_1() {
	ch := make(chan string, 1)

	ch <- "123"

	now(<-ch, true)
}

/*
缓冲管道限制并发数量(控制同时运行的协程数量)
控制执行顺序
*/
func channel13_2() {
	tasks := []int{3, 2, 1}
	tasksLen := len(tasks)
	timeout := 2
	chs := make([]chan string, tasksLen)
	startTime := time.Now()
	/*
		创建一个缓冲管道(chs),作为并发限制的计数器
		1为并发限制,表示限制并发数为1,

		则任务0 超时2s ,运行任务1 超时2s,任务3,耗时1s,
		总耗时 2+2+1= 5s

		并发限制为2时
		任务0和任务1并发运行,超时2s,任务2 耗时1s
		总耗时 2+1 = 3s
	*/

	chLimit := make(chan bool, 2)
	limitFunc := func(chLimit <-chan bool, ch chan string, task_id, sleeptime, timeout int) {
		//阻塞等待
		channel13_2_Run(task_id, sleeptime, timeout, ch)

		//在缓冲为1的chLimit,等待1个任务channel13_2_Run执行完毕
		//在缓冲为2的chLimit,等待2个任务channel13_2_RUn执行完毕
		<-chLimit
	}
	now("multi tasks start,totalNum:"+strconv.Itoa(tasksLen)+"\n", false)

	for i, sleeptime := range tasks {
		chs[i] = make(chan string, 1)
		//在前面的一个未读出时阻塞
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleeptime, timeout)
	}

	for _, ch := range chs {
		now(<-ch, true)
	}

	processTime := time.Since(startTime).String()
	now("muti tasks ended. processTime:"+processTime+"\n", false)
}

func channel13_2_Run(task_id, sleeptime, timeout int, ch chan<- string) {
	ch_run := make(chan string)

	go channel13_2_run(task_id, sleeptime, ch_run)

	//任务channel13_2_run超时限制处理,等待任务执行完毕/超时 返回
	select {
	//任务执行完毕
	case re := <-ch_run:
		ch <- re

		//任务超时处理
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d,time out", task_id)
		ch <- re
	}
}

func channel13_2_run(task_id, sleeptime int, ch chan<- string) {
	time.Sleep(time.Duration(sleeptime) * time.Second)

	ch <- fmt.Sprintf("task id %d, sleep %d second", task_id, sleeptime)
	return
}
