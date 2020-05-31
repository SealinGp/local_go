package timewheel

import (
	"algorithm/timewheel/delayqueue"
	"errors"
	"time"
	"unsafe"
)

type TimingWheel struct {
	tick int64       //milliseconds 转动1个槽需要的时间
	wheelSize int64  //时间轮有多少个槽

	interval  int64   //时间轮转1圈需要的时间(ms)
	currentTime int64 //milliseconds 当前是第几个槽
	buckets []*bucket //所有的槽
	queue   *delayqueue.DelayQueue //包含定时任务的延迟队列,靠这个来驱动时间轮转动

	//高等级时间轮 type *TimingWheel
	//这个字段可能会通过Add函数被不断的更新和读取
	overflowWheel unsafe.Pointer

	exitC chan struct{}
	waitGroup waitGroupWrapper
}

//精度 1ms
func NewTimingWheel(tick time.Duration,wheelSize int64) *TimingWheel {
	tickMs := int64(tick / time.Millisecond)

	//不足1ms报错
	if tickMs < 0 {
		panic(errors.New("required: tick >= 1ms"))
	}

	startMs := timeToMs(time.Now().UTC())

	return newTimingWheel(
		tickMs,
		wheelSize,
		startMs,
		delayqueue.New(int(wheelSize)),
	)
}

func newTimingWheel(tickMs int64, wheelSize int64, startMs int64,queue *delayqueue.DelayQueue) *TimingWheel {
	buckets := make([]*bucket,wheelSize)
	for i := range buckets {
		buckets[i] = newBucket()
	}
	return &TimingWheel{
		tick:tickMs,
		wheelSize:wheelSize,
		currentTime:startMs,
		interval: tickMs * wheelSize,
		buckets:buckets,
		queue:queue,
		exitC:make(chan struct{}),
	}
}