package datastructure

type cellStack struct {
	element int
	next    *cellStack
}

type Stack struct {
	init *cellStack
}

func NewStack() *Stack {
	return &Stack{init: nil}
}

func (s *Stack) Push(n int) {
	c := &cellStack{element: n, next: s.init}
	s.init = c
}

func (s *Stack) Pop() int {
	popped := s.init.element
	s.init = s.init.next
	return popped
}
