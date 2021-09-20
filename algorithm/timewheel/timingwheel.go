package timewheel

import (
	"errors"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/SealinGp/local_go/algorithm/timewheel/delayqueue"
)

type TimingWheel struct {
	tick      int64 //milliseconds 转动1个槽需要的时间
	wheelSize int64 //时间轮有多少个槽

	interval    int64                  //时间轮转1圈需要的时间(ms)
	currentTime int64                  //milliseconds 当前是第几个槽
	buckets     []*bucket              //所有的槽
	queue       *delayqueue.DelayQueue //包含定时任务的延迟队列,靠这个来驱动时间轮转动

	//高等级时间轮 type *TimingWheel
	//这个字段可能会通过Add函数被不断的更新和读取
	overflowWheel unsafe.Pointer

	exitC     chan struct{}
	waitGroup waitGroupWrapper
}

//精度 1ms
func NewTimingWheel(tick time.Duration, wheelSize int64) *TimingWheel {
	tickMs := int64(tick / time.Millisecond)

	//不足1ms报错
	if tickMs <= 0 {
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

func newTimingWheel(tickMs int64, wheelSize int64, startMs int64, queue *delayqueue.DelayQueue) *TimingWheel {
	buckets := make([]*bucket, wheelSize)
	for i := range buckets {
		buckets[i] = newBucket()
	}
	return &TimingWheel{
		tick:        tickMs,
		wheelSize:   wheelSize,
		currentTime: truncate(startMs, tickMs),
		interval:    tickMs * wheelSize,
		buckets:     buckets,
		queue:       queue,
		exitC:       make(chan struct{}),
	}
}

//添加定时器到当前时间轮的槽中去
func (tw *TimingWheel) add(t *Timer) bool {
	currentTime := atomic.LoadInt64(&tw.currentTime)

	//过期时间不足1ms则添加失败,也就是说添加任务的过期时间至少要 >= 1ms
	if t.expiration < currentTime+tw.tick {
		return false

		//过期时间小于时间轮1圈
	} else if t.expiration < currentTime+tw.interval {
		//散列放置任务到对应的槽上
		virtualID := t.expiration / tw.tick //过期时间需要转动多少次
		b := tw.buckets[virtualID%tw.wheelSize]
		b.Add(t)

		//把这个槽的延迟队列过期时间更新
		if b.SetExpiration(virtualID * tw.tick) {
			tw.queue.Offer(b, b.expiration)
		}
		return true

		//过期时间超出当前时间轮的总时间(interval)则新建一个大刻度的时间轮
	} else {
		overflowWheel := atomic.LoadPointer(&tw.overflowWheel)
		if overflowWheel == nil {
			atomic.CompareAndSwapPointer(
				&tw.overflowWheel,
				nil,
				unsafe.Pointer(newTimingWheel(
					tw.interval,
					tw.wheelSize,
					currentTime,
					tw.queue,
				)),
			)
			overflowWheel = atomic.LoadPointer(&tw.overflowWheel)
		}

		return (*TimingWheel)(overflowWheel).add(t)
	}
}

func (tw *TimingWheel) addOrRun(t *Timer) {
	//当添加的定时器超时时间不足 时间轮的最小精度的时候,则直接执行(< tw.tick)
	if !tw.add(t) {
		go t.task()
	}
}

//高等级时间轮
func (tw *TimingWheel) advanceClock(expiration int64) {
	currentTime := atomic.LoadInt64(&tw.currentTime)

	//如果该定时器还未过期
	if expiration >= currentTime+tw.tick {
		currentTime = truncate(expiration, tw.tick)
		atomic.StoreInt64(&tw.currentTime, currentTime)

		overflowWheel := atomic.LoadPointer(&tw.overflowWheel)
		if overflowWheel != nil {
			(*TimingWheel)(overflowWheel).advanceClock(currentTime)
		}
	}
}

func (tw *TimingWheel) Start() {
	tw.waitGroup.Wrap(func() {
		tw.queue.Poll(tw.exitC, func() int64 {
			return timeToMs(time.Now().UTC())
		})
	})
	tw.waitGroup.Wrap(func() {
		for {
			select {
			case ele := <-tw.queue.C:
				b := ele.(*bucket)
				tw.advanceClock(b.expiration)
				//随着时间的迁移,定时任务越来越近,当定时任务的精度到达最小时间轮的精度的时候,则将其从大时间轮移动到小时间轮中去
				b.Flush(tw.addOrRun)
			case <-tw.exitC:
				return
			}
		}
	})
}

// 停止当前时间轮
//
// If there is any timer's task being running in its own goroutine, Stop does
// not wait for the task to complete before returning. If the caller needs to
// know whether the task is completed, it must coordinate with the task explicitly.
func (tw *TimingWheel) Stop() {
	close(tw.exitC)
	tw.waitGroup.Wait()
}

func (tw *TimingWheel) AfterFunc(d time.Duration, f func()) *Timer {
	t := &Timer{
		expiration: timeToMs(time.Now().UTC().Add(d)),
		task:       f,
	}
	tw.addOrRun(t)
	return t
}

type Scheduler interface {
	Next(time.Time) time.Time
}

func (tw *TimingWheel) ScheduleFunc(s Scheduler, f func()) (t *Timer) {
	expiration := s.Next(time.Now().UTC())
	if expiration.IsZero() {
		return
	}

	t = &Timer{
		expiration: timeToMs(expiration),
		task: func() {
			expiration := s.Next(msToTime(t.expiration))
			if !expiration.IsZero() {
				t.expiration = timeToMs(expiration)
				tw.addOrRun(t)
			}
			f()
		},
	}
	tw.addOrRun(t)

	return
}
