package main;
import(
	"fmt"
	"sync"
	"runtime"
	"os"
);

/*

进程(process):
 (面向进程设计的)UNIX/Linux version<Linux2.4: 进程是程序的实体
 (面向线程设计的)win/mac os/Linux version>=Linux2.6: 进程是线程的容器
 程序本身只是指令,数据,以及组织形式的描述,进程是程序的真正实例
 多个进程 可能 与同一个程序有关,每个进程可 同步(顺序) 或 异步(平行) 的方式运行
 现代计算机系统可在同一段时间内以进程的形式将多个程序加载到存储器中,并借由时间共享(时分复用),
 以在一个cpu(处理器)上表现出同时运行的感觉,
线程(thread):
 运算调度的最小单位,一个进程可以有多条线程执行不同的任务(一个线程可以并发多个线程),
 多条线程共享该进程中的资源(如虚拟地址空间,文件描述符?).
 文件描述符:索引值,当程序打开一个现有文件或者创建一个新文件时,内核向进程返回一个文件描述符
协程:
 程序组件(子例程也是),协程比子例程更灵活,源于Simula和Modula-2语言,应用:多任务,迭代器,无限列表,管道pipe 
 轻量级的线程(用户级线程,子程序,函数),对内核透明(系统并不知道有协程的存在),完全由用户自己的程序进行调度,
 因为是由用户程序自己控制,那么就很难像抢占式调度那样做到强制的 CPU 控制权切换到其他进程/线程,通常只能进
 行协作式调度,需要协程自己主动把控制权转让出去之后,其他协程才能被执行到
总结:
 1.进程,线程系统级别,协程语言级
 2.每个进程至少有一个线程,线程是真正的运行单位
 3.线程进程都是同步机制,协程是异步
 4.IO密集型一般使用多线程或者多进程,CPU密集型一般使用多进程,非阻塞异步并发一般使用协程

go 并发
本质上,goroutine 就是协程.不同的是,Golang 在 runtime.系统调用等多方面对 goroutine 调
度进行了封装和处理,当遇到长时间执行或者进行系统调用时,会主动把当前 goroutine 的CPU (P)
转让出去,让其他 goroutine 能被调度并执行,也就是 Golang 从语言层面支持了协程.Golang 的
一大特色就是从语言层面原生支持协程,在函数或者方法前面加 go关键字就可创建一个协程.

线程和协程比较
 内存消耗:
  每个 goroutine(协程)默认占用内存远比Java,C的线程少
  goroutine：2KB 
　线程：8MB
 切换调度开销方面:
  goroutine远比线程小
  线程：涉及模式切换(从用户态切换到内核态),16个寄存器,PC,SP...等寄存器的刷新等
  goroutine：只有三个寄存器的值修改 - PC / SP / DX.
在网络编程中,我们可以理解为 Golang 的协程本质上其实就是对IO事件的封装,并且通过语言级的支持让异步的代码看上去像同步执行的一样
*/ 
func main() {
	args := os.Args;	
	execute(args[1]);
}

func execute(n string) {
	funs := map[string]func() {
		"gorun"  : gorun,
		"gorun2" : gorun2,
		"gorun3" : gorun3,
		"gorun4" : gorun4,
		"gorun5" : gorun5,
	};	
	funs[n]();		
}


/*
  main函数启动了5个goroutine,然后返回,
  这时程序退出,被启动的say()的协程还没
  来得及执行,就退出了,因此不会打印任何东西.
  我们需要控制main函数等待所有goroutine退出后再返回,
  所以就涉及到了通信的问题
*/
func gorun() {
	for i := 0; i < 5; i++ {
	  go say("hello ");		
	}
}
func say(msg string) {
	fmt.Println(msg);
}

/*
 常见的两种并发通信模型:共享内存和消息
 共享内存:使用锁变量来同步协程,golang主要使用channel作为通信模型
 锁:共享读锁,独占写锁
*/

//锁变量实现同步协程
var counter int = 0;
func Count(lock *sync.Mutex) {
	lock.Lock(); //写锁
	counter++;
	fmt.Println("counter=",counter);
	lock.Unlock();//写解锁
}
func gorun2() {
	//建立互斥锁实例 一个互斥锁只能同时被一个 goroutine 锁定,其它 goroutine 将阻塞直到互斥锁被解锁(重新争抢对互斥锁的锁定)
	lock := &sync.Mutex{};
	//并行10个协程
	for i := 0; i < 10; i++ {
		go Count(lock);
	}

	//无限循环,控制等待上面10个协程执行结束后,释放写锁后,变量为10,跳出
	for {
		lock.Lock();
		c := counter;
		lock.Unlock();

		//这个函数的作用是让当前goroutine让出CPU,好让其它的goroutine获得执行的机会
		runtime.Gosched();
		if c >= 10 {
			break;
		}
	}
}
//runtime.Gosched作用:停留给协程执行的一个机会/最后一个协程优先执行完毕 9 0~8
func gorun3() {
	for i := 0; i < 10; i++ {
        go show(i);
    }

    runtime.Gosched();
    fmt.Println("end");
}
func show(i int) {
	fmt.Println(i);
}

/*
channel实现同步协程
channel是进程内的通信方式
默认情况下channel的接收[写入]和发送[读取]都是阻塞的,除非另一端已写/读
var channelName chan Type
	//使用
	var ch1 chan int;
	ch2 := make(chan int);

	//写入数据到channel,会阻塞,直到有其他协程从这个channel中读取数据
	ch1 <- 1;

	//从channel中读取数据,如果channel中没有写入数据,则阻塞,直到channel被写入数据为止
	val := <-ch1;
*/
func gorun4() { 
	len1 := 100;
	//创建一个数组channel 长度(缓冲)为10个,在缓冲区被写完之前不会阻塞
	chs := make([] chan int,len1);

	//同时写入10个channel,写完后然后开始阻塞,等待读解除阻塞
	for i := 0; i < len1; i++ {		
		chs[i] = make(chan int);
		go Count1(chs[i],i);
	}

	//按顺序读出,解除阻塞
	for _,ch := range chs {
		<-ch;
	}
}
func Count1(chIndex chan int,index int) {
	chIndex <- 1;
	fmt.Println("counting index=",index);
}

/*
select
 Unix中,select()函数用于监控一组描述符,该机制常被用于实现高并发的socket服务器程序
 Golang中,用于处理异步IO问题,select默认是阻塞的,只有当监听的channel中有发送或接收可以进行时才会运行,
 当多个channel都准备好的时候,select是随机的选择一个执行,
 使用方式
 	select {
 		//成功读取数据
 		case <- chan1:

 		//成功写入数据
 		case chan1 <- 1:

 		//默认
 		default:
 	}
*/
 func gorun5() {
 	for i := 0; i < 5; i++ {
 		go test(i);
 	}

 	select{};
 }
 func test(i int) {
 	fmt.Println(i);
 }

