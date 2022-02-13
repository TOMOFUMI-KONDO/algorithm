package datastructure

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)

	if s.init.element != 2 {
		t.Errorf("s.init.element == %d; want 2", s.init.element)
	}
	if s.init.next.element != 1 {
		t.Errorf("s.init.next.element != %d; want 1", s.init.next.element)
	}

	popped := s.Pop()

	if popped != 2 {
		t.Errorf("popped == %d; want 2", popped)
	}

	popped = s.Pop()

	if popped != 1 {
		t.Errorf("popped == %d; want 1", popped)
	}
}
