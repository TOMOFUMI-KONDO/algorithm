package datastructure

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)

	popped := s.Pop()
	if popped != 2 {
		t.Errorf("popped == %d; want 2", popped)
	}

	popped = s.Pop()
	if popped != 1 {
		t.Errorf("popped == %d; want 1", popped)
	}
}
