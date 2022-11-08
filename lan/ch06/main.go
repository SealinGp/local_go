package main

import (
	"log"
	"sync"
)

func main() {

	ch := make(chan struct{})

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Close(ch)
		}()
	}

	wg.Wait()
}

func Close(ch chan struct{}) {
	select {
	case <-ch:
		log.Printf("[E] closed")
		return
	default:
		close(ch)
	}
}
