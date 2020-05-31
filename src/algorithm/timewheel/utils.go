package timewheel

import (
	"sync"
	"time"
)

type waitGroupWrapper struct {
	sync.WaitGroup
}

func truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}
func timeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
func msToTime(t int64) time.Time {
	return time.Unix(0,t*int64(time.Millisecond)).UTC()
}

func (w *waitGroupWrapper)Wrap(cb func())  {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}