package main

type cell struct {
	element int
	next    *cell
}

type Stack struct {
	init *cell
}

func NewStack() *Stack {
	return &Stack{init: nil}
}

func (s *Stack) Push(n int) {
	c := &cell{element: n, next: s.init}
	s.init = c
}

func (s *Stack) Pop() int {
	popped := s.init.element
	s.init = s.init.next
	return popped
}
