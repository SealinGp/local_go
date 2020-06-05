package timewheel_test

import (
	"algorithm/timewheel"
	"strconv"
	"testing"
	"time"
)

//测试延迟任务是否成功
func TestTimingWheel_AfterFunc(t *testing.T) {
	tw := timewheel.NewTimingWheel(time.Millisecond,20)
	tw.Start()
	defer tw.Stop()

	durations := []time.Duration{
		1 * time.Millisecond,
		5 * time.Millisecond,
		10 * time.Millisecond,
		50 * time.Millisecond,
		100 * time.Millisecond,
		500 * time.Millisecond,
		1 * time.Second,
	}

	for i, d := range durations {
		t.Run("now:"+d.String(), func(t *testing.T) {
			exitC := make(chan time.Time)

			start := time.Now().UTC()
			tw.AfterFunc(d, func() {
				t.Log("now" + strconv.Itoa(i))
				exitC <- time.Now().UTC()
			})
			got := (<-exitC).Truncate(time.Millisecond)
			min := start.Add(d).Truncate(time.Millisecond)

			err := 5 * time.Millisecond
			if got.Before(min) || got.After(min.Add(err)) {
				t.Errorf("Timer(%s) expiration: want[%s, %s], got %s",d,min,min.Add(err),got)
			}
		})
	}
}

type scheduler struct {
	intervals []time.Duration
	current   int
}

func (s *scheduler)Next(prev time.Time) time.Time {
	if s.current >= len(s.intervals) {
		return time.Time{}
	}
	next := prev.Add(s.intervals[s.current])
	s.current++
	return next
}

func TestTimingWheel_ScheduleFunc(t *testing.T)  {
	tw := timewheel.NewTimingWheel(time.Millisecond,20)
	tw.Start()
	defer tw.Stop()

	s := &scheduler{
		intervals:[]time.Duration{
			1 * time.Millisecond, //start + 1ms
			4 * time.Millisecond, //start + 5ms
			5 * time.Millisecond, //start + 10ms
			40 * time.Millisecond, //start + 50ms
			50 * time.Millisecond, //start + 100ms
			400 * time.Millisecond, //start + 500ms
			500 * time.Millisecond, //start + 1000ms
		},
	}
	exitC := make(chan time.Time, len(s.intervals))

	start := time.Now().UTC()
	tw.ScheduleFunc(s, func() {
		exitC <- time.Now().UTC()
	})

	accum := time.Duration(0)
	for _, d := range s.intervals {
		got := (<-exitC).Truncate(time.Millisecond)
		accum += d
		min := start.Add(accum).Truncate(time.Millisecond)
		err := 5 * time.Millisecond

		//精度相差 +- 5ms则报错
		if got.Before(min) || got.After(min.Add(err)) {
			t.Errorf("Timer(%s) expiration: want [%s, %s] got %s",accum,min,min.Add(err),got)
		}
	}
}