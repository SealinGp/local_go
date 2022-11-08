package simple

type Stack struct {
	arr []int
}

func NewStack1() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *Stack) Pop() int {
	if len(s.arr) == 0 {
		return 0
	}

	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return v
}

func (s *Stack) Len() int {
	return len(s.arr)
}

type CQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func ConstructorCQ() CQueue {
	return CQueue{
		stack1: NewStack1(),
		stack2: NewStack1(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() <= 0 {
		for this.stack1.Len() > 0 {
			this.stack2.Push(this.stack1.Pop())
		}
	}

	if this.stack2.Len() > 0 {
		return this.stack2.Pop()
	}

	return 0
}
