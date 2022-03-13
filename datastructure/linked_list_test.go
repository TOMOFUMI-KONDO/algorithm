package datastructure

import (
	"testing"
)

func TestLinkedListInsert(t *testing.T) {
	list := NewLinkedList(0)

	list.Insert(1, 0)
	actual := list.Get(0)
	if actual != 1 {
		t.Errorf("list[1] = %d; want 1", actual)
	}

	list.Insert(2, 2)
	actual = list.Get(2)
	if actual != 2 {
		t.Errorf("list[1] = %d; want 2", actual)
	}
}

func TestLinkedListDelete(t *testing.T) {
	list := LinkedList{
		base: &cellLinkedList{
			val: 0, next: &cellLinkedList{
				val: 1, next: &cellLinkedList{
					val: 2, next: nil,
				},
			},
		},
		len: 3,
	}

	list.Delete(1)

	var actual int
	actual = list.Get(0)
	if actual != 0 {
		t.Errorf("list[0] = %d; want 2", actual)
	}
	actual = list.Get(1)
	if actual != 2 {
		t.Errorf("list[1] = %d; want 2", actual)
	}
}
