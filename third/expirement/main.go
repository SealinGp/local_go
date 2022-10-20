package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

var delta int32

func main() {
	rand.Seed(time.Now().Unix())

	eg := new(errgroup.Group)

	size := rand.Intn(10)
	for i := 0; i < size; i++ {
		tmp := i
		eg.Go(func() error {
			return TestXxx(tmp)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Printf("[E] size:%v, err:%v, delta:%v", size, err, delta)
		return
	}

	log.Printf("[I] done. size:%v, delta:%v", size, delta)
}

//none block
func TestXxx(i int) error {
	atomic.AddInt32(&delta, 1)

	r := rand.Intn(10)
	if r <= 5 {
		err := fmt.Errorf("rand err:%vs", r)
		log.Printf("[%v] err:%v", i, err)
		time.Sleep(time.Second * time.Duration(r))
		return err
	}

	return nil
}
