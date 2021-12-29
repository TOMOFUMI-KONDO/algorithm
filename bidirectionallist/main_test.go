package main

import "testing"

func TestBidirectionalListInsert(t *testing.T) {
	list := &Cell{val: 0}

	v := list.Val()
	if v != 0 {
		t.Errorf("list[0] = %d; want 0", v)
	}

	list.InsertNext(1)
	v = list.Next().Val()
	if v != 1 {
		t.Errorf("list[1] = %d; want 1", v)
	}
	v = list.Next().Prev().Val()
	if v != 0 {
		t.Errorf("list[0] = %d; want 0", v)
	}

	list.Next().InsertNext(2)
	v = list.Next().Next().Val()
	if v != 2 {
		t.Errorf("list[2] = %d; want 2", v)
	}

	list.InsertPrev(-1)
	v = list.Prev().Val()
	if v != -1 {
		t.Errorf("list[-1] = %d; want -1", v)
	}
	v = list.Prev().Next().Val()
	if v != 0 {
		t.Errorf("list[0] = %d; want 0", v)
	}

	list.Prev().InsertPrev(-2)
	v = list.Prev().Prev().Val()
	if v != -2 {
		t.Errorf("list[-2] = %d; want -2", v)
	}
}

func TestBidirectionalListDelete(t *testing.T) {
	list := &Cell{
		val: 0,
		next: &Cell{
			val:  1,
			next: &Cell{val: 2},
		},
		prev: &Cell{
			val:  -1,
			prev: &Cell{val: -2},
		},
	}

	list.DeleteNext()
	v := list.Next().Val()
	if v != 2 {
		t.Errorf("list[1] = %d; want 2", v)
	}
	v = list.Next().Prev().Val()
	if v != 0 {
		t.Errorf("list[0] = %d; want 0", v)
	}

	list.DeletePrev()
	v = list.Prev().Val()
	if v != -2 {
		t.Errorf("list[-1] = %d; want -2", v)
	}
	v = list.Prev().Next().Val()
	if v != 0 {
		t.Errorf("list[0] = %d; want 0", v)
	}
}
