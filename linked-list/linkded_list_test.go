package main

import (
	"testing"
)

func TestLinkedListInsert(t *testing.T) {
	list := &Cell{element: 0, next: nil}

	Insert(1, list)
	if list.next.element != 1 {
		t.Errorf("list[1] = %d; want 1", list.next.element)
	}

	Insert(2, list)
	if list.next.element != 2 {
		t.Errorf("list[1] = %d; want 2", list.next.element)
	}
}

func TestLinkedListDelete(t *testing.T) {
	list := &Cell{element: 0, next: &Cell{element: 1, next: &Cell{element: 2, next: nil}}}

	Delete(list)
	if list.next.element != 2 {
		t.Errorf("list[1] = %d; want 2", list.next.element)
	}
}
