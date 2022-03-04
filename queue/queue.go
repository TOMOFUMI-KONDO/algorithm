package queue

type Queue struct {
	first *element
	last  *element
}

type element struct {
	value int
	next  *element
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val int) {
	if q.first == nil {
		q.first = &element{value: val}
		q.last = q.first
	} else {
		q.last.next = &element{value: val}
	}
}

func (q *Queue) Pop() int {
	popped := q.first.value
	q.first = q.first.next
	return popped
}
