package main

func main() {

}

type L struct {
	c chan int
}

func (l L)Lock() bool {
	isLocked := false
	select {
	case <-l.c:
		isLocked = true
	default:
		isLocked = false
	}
	return isLocked
}
func (l L)Unlock()  {
	l.c <- 1
}
func NewL() L {
	lo := L{
		c : make(chan int),
	}
	lo.c <- 1
	return lo
}