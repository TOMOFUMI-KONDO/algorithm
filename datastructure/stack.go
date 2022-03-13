package datastructure

type element struct {
	value int
	next  *element
}

type Stack struct {
	top *element
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (s *Stack) Push(val int) {
	s.top = &element{value: val, next: s.top}
}

func (s *Stack) Pop() int {
	popped := s.top.value
	s.top = s.top.next
	return popped
}
