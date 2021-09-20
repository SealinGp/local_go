package timewheel_test

import (
	"fmt"
	"time"
)

func Example_startTimer() {
	tw := NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	exitC := make(chan time.Time, 1)
	tw.AfterFunc(time.Second, func() {
		fmt.Println("The timer fires")
		exitC <- time.Now()
	})

	<-exitC
}

func Example_stopTimer() {
	tw := NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()
	t := tw.AfterFunc(time.Second, func() {
		fmt.Println("The timer fires")
	})

	<-time.After(900 * time.Millisecond)
	t.Stop()
}
