package datastructure

import (
	"fmt"
	"testing"
)

func TestBidirectionalListInsert(t *testing.T) {
	list := NewBidirectionalList(0)

	actual := list.Get(0)
	if actual != 0 {
		t.Errorf("list[0] = %d; want 0", actual)
	}

	list.Insert(1, 1)
	fmt.Printf("%#v\n", list)
	actual = list.Get(1)
	if actual != 1 {
		t.Errorf("list[1] = %d; want 1", actual)
	}
	actual = list.Get(0)
	if actual != 0 {
		t.Errorf("list[0] = %d; want 0", actual)
	}

	list.Insert(2, 2)
	actual = list.Get(2)
	if actual != 2 {
		t.Errorf("list[2] = %d; want 2", actual)
	}

	list.Insert(-1, 0)
	actual = list.Get(0)
	if actual != -1 {
		t.Errorf("list[-1] = %d; want -1", actual)
	}
	actual = list.Get(1)
	if actual != 0 {
		t.Errorf("list[0] = %d; want 0", actual)
	}

	list.Insert(-2, 0)
	actual = list.Get(0)
	if actual != -2 {
		t.Errorf("list[-2] = %d; want -2", actual)
	}
}

func TestBidirectionalListDelete(t *testing.T) {
	cell1 := &cellBiDirList{0, nil, nil}
	cell2 := &cellBiDirList{1, nil, nil}
	cell3 := &cellBiDirList{2, nil, nil}
	cell4 := &cellBiDirList{4, nil, nil}
	cell5 := &cellBiDirList{5, nil, nil}
	cell1.next = cell2
	cell1.prev = cell5
	cell2.next = cell3
	cell2.prev = cell1
	cell3.next = cell4
	cell3.prev = cell2
	cell4.next = cell5
	cell4.prev = cell3
	cell5.next = cell1
	cell5.prev = cell4
	list := &BidirectionalList{cell1, cell5, 5}

	list.Delete(1)
	actual := list.Get(1)
	if actual != 2 {
		t.Errorf("list[1] = %d; want 2", actual)
	}
	actual = list.Get(0)
	if actual != 0 {
		t.Errorf("list[0] = %d; want 0", actual)
	}

	list.Delete(2)
	actual = list.Get(3)
	if actual != 5 {
		t.Errorf("list[-1] = %d; want -2", actual)
	}
	actual = list.Get(1)
	if actual != 2 {
		t.Errorf("list[0] = %d; want 0", actual)
	}
}
