package main

import (
	"fmt"
	"strconv"
	"sync"
)
func main() {

}

//goroutine 交替打印
func switchPrint()  {
	c1 := make(chan string)
	c2 := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 3; i++ {
			c1 <- strconv.Itoa(i)

			s := <-c2
			fmt.Print(s)
		}
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 'a'; i < 'd'; i++ {
			v := <-c1

			fmt.Print(v)

			c2 <- string(i)
		}
	}()
	wg.Wait()
	fmt.Println("")
}