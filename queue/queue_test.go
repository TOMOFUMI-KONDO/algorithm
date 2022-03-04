package queue

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)

	popped := q.Pop()
	if popped != 1 {
		t.Fatalf("popped = %d; want 1", popped)
	}

	popped = q.Pop()
	if popped != 2 {
		t.Fatalf("popped = %d; want 2", popped)
	}
}
