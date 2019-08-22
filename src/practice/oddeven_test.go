package practice

import "testing"

func Even(i int) bool  {
	return i%2 == 0
}
func Odd(i int) bool  {
	return i%2 != 0
}

func TestEven(t *testing.T)  {
	if !Even(10) {
		t.Log("10 is even")
		t.Fail()
	}
	if Even(7) {
		t.Log("7 is not even")
		t.Fail()
	}
}
func TestOdd(t *testing.T)  {
	if !Odd(11) {
		t.Log("11 is odd")
		t.Fail()
	}
	if Odd(10) {
		t.Log("10 is not odd")
		t.Fail()
	}
}