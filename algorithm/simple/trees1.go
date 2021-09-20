package simple

import "container/list"

type queue1 struct {
	l *list.List
}

func NewQueue1() *queue1 {
	q := &queue1{
		l: list.New(),
	}
	return q
}

func (q *queue1) InQueue(val interface{}) {
	q.l.PushBack(val)
}

func (q *queue1) OutQueue() interface{} {
	e := q.l.Front()
	if e != nil {
		q.l.Remove(e)
		return e.Value
	}

	return nil
}
