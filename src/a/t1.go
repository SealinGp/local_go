package a

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
func main() {

}
