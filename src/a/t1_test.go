package a

import (
	"testing"
)


func BenchmarkA1(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		A1()
	}
}