package main

import (
	"log"
	"os"
	"sync"
	"syscall"
)

//https://github.com/chai2010/advanced-go-programming-book/blob/master/ch6-cloud/ch6-02-lock.md
func main() {
	m := map[string]func(){
		"lock1":lock1,
		"lock2":lock2,
		"lock3":lock3,
	}
	m[os.Args[1]]()
}

var counter int
func lock1()  {
	var wg sync.WaitGroup
	for i := 0 ; i < 1000 ; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	log.Println(counter)
}

func lock2()  {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	log.Println(counter)
}

//秒杀|抢红包的原理,抢完了就没了
type LockS struct {
	c chan struct{}
}

//满不可塞,空不可取
func (l LockS)Lock() bool {
	lockResult := false
	select {
		case <-l.c:
			lockResult = true
	default:
	}
	return lockResult
}
func (l LockS)Unlock()  {
	l.c <- struct{}{}
}
func NewLockS() LockS {
	l :=  LockS{}
	l.c = make(chan struct{},2)
	l.c <- struct{}{}
	l.c <- struct{}{}
	return l
}
func lock3()  {
	ls := NewLockS()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(is int) {
			defer wg.Done()
			if !ls.Lock() {
				log.Println("goroutine ",is,"lock failed")
				return
			}
			counter++
			log.Println("goroutine ",is,"lock success!")
			log.Println("cur counter",counter)
			ls.Unlock()
		}(i)
	}
	wg.Wait()
}

func lock4()  {
	syscall.Socket()
}