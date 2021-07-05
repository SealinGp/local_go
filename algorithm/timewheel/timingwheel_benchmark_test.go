package timewheel_test

import (
	"algorithm/timewheel"
	"testing"
	"time"
)

var cases = []struct {
	name string
	N    int
}{
	{"N-1m", 1000000},   //100w
	{"N-5m", 5000000},   //500w
	{"N-10m", 10000000}, //1k w
}

func genD(i int) time.Duration {
	return time.Duration(i%10000) * time.Millisecond
}
func BenchmarkTimingWheel_StartStop(b *testing.B) {
	tw := timewheel.NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			base := make([]*timewheel.Timer, c.N)
			for i := 0; i < len(base); i++ {
				base[i] = tw.AfterFunc(genD(i), func() {})
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				tw.AfterFunc(time.Second, func() {}).Stop()
			}

			b.StopTimer()
			for i := 0; i < len(base); i++ {
				base[i].Stop()
			}
		})
	}
}

func BenchmarkStandardTimer_StartStop(b *testing.B) {
	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			base := make([]*time.Timer, c.N)
			for i := 0; i < len(base); i++ {
				base[i] = time.AfterFunc(genD(i), func() {})
			}
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				time.AfterFunc(time.Second, func() {}).Stop()
			}

			b.StopTimer()
			for i := 0; i < len(base); i++ {
				base[i].Stop()
			}
		})
	}
}
