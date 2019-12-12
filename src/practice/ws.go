package main

import (
	"fmt"
)

//连接池的简单实现
type pool struct {
	workFun    chan func()
	maxRoutine chan struct{}
}
func main() {
	res := make(chan string)
	ts := func() {
		res <- "finished"
	}
	pools := NewPool(2)

	pools.PoolControl(ts)
	pools.PoolControl(ts)
	pools.PoolControl(ts)
	fmt.Println(<-res)

	//for r := range res  {
	//	fmt.Println(r)
	//}
}
func NewPool(maxRoutineNum int) *pool {
	return &pool{
		workFun:    make(chan func()),
		maxRoutine: make(chan struct{},maxRoutineNum),
	}
}

func (p *pool)PoolControl(task func()) {
	select {
	//这里等待
	case p.workFun <- task:
		fmt.Println("?")

		//先进这里,当p.maxRoutine里面的缓存满了以后,这里堵塞
	case p.maxRoutine <- struct{}{}:
		fmt.Println("?1")
		go p.work(task)
	}
}

func (p *pool)work(task func())  {
	defer func() {
		<-p.maxRoutine
	}()
	for {
		//task 里面做完了就发给p.workFun
		task()
		task = <-p.workFun
	}
}