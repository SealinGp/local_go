package main

import "fmt"

func main() {
	c1 := make(chan bool,1)
	c2 := make(chan bool)
	d  := make(chan bool)

	c1 <- true

	go func() {
		for i := 0; i < 3; i++ {
			<-c1
			fmt.Print(i)
			c2<-true
		}
	}()
	go func() {
		for i := 'a'; i < 'd'; i++ {
			<-c2
			fmt.Print(string(i))
			c1<-true
		}
		d<-true
	}()


	<-d
}