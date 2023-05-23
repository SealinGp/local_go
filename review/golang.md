goroutine:
  单线程调度器 0.x
  多线程调度器 1.x
  任务窃取调度器 1.1 GMP 模型

  - G: go routine 用户态的线程
  - M: 系统级线程
  - P: 处理器
    抢占调度 1.2~1.13
    基于信号的抢占式调度器 1.14~ now
    - 垃圾回收在扫描栈时会触发抢占式调度

编译过程: 
  1.词法/语法分析 
  2.类型检查 
  3.中间码生成
  4.机器码生成

数据类型:
  - 字符串:
  - 数组:
  - 切片: 

  if newCap > oldCap*2 {
    wantCap = newCap
  }
  - map

  ```go
  type hmap struct {
    count int
    B int
    buckets unintptr
    oldBuckets uintptr
    extra *mapextra
  }

  type mapextra struct {

  }

  - 扩容:
    原容量: x 
    需要的容量: y
    新容量: z
    y > 2*x => z = y
    y < 1024 => z = 2*x
    y > 1024 => z = x*(1+0.25)x  
  ```

函数调用:
  
