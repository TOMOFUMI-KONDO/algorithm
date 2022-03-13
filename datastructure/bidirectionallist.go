package datastructure

import "fmt"

type BidirectionalList struct {
	head *cellBiDirList
	end  *cellBiDirList
	len  int
}

func NewBidirectionalList(v int) *BidirectionalList {
	cell := &cellBiDirList{v, nil, nil}
	cell.next = cell
	cell.prev = cell
	return &BidirectionalList{cell, cell, 1}
}

type cellBiDirList struct {
	val  int
	next *cellBiDirList
	prev *cellBiDirList
}

func (b *BidirectionalList) Get(n int) int {
	return b.get(n).val
}

func (b *BidirectionalList) Insert(v int, n int) {
	if n > b.len {
		panic(fmt.Errorf("n must be lower or equal than list length(%d), but got %d", b.len, n))
	}
	defer func() { b.len++ }()

	newCell := &cellBiDirList{v, nil, nil}

	if n == 0 {
		newCell.next = b.head
		newCell.prev = b.end
		b.head.prev = newCell
		b.end.next = newCell
		b.head = newCell
		return
	}

	if n == b.len {
		newCell.next = b.head
		newCell.prev = b.end
		b.head.prev = newCell
		b.end.next = newCell
		b.end = newCell
		return
	}

	inserted := b.get(n)
	newCell.next = inserted
	newCell.prev = inserted.prev
	inserted.prev.next = newCell
	inserted.prev = newCell
}

func (b *BidirectionalList) Delete(n int) {
	if n > b.len-1 {
		return
	}

	deleted := b.get(n)
	deleted.prev.next = deleted.next
	deleted.next.prev = deleted.prev

	if n == 0 {
		b.head = deleted.next
	}
	if n == b.len-1 {
		b.end = deleted.prev
	}

	b.len--
}

func (b *BidirectionalList) get(n int) *cellBiDirList {
	if n < b.len/2 {
		cell := b.head
		for i := 0; i < n; i++ {
			cell = cell.next
		}
		return cell
	} else {
		cell := b.end
		for i := 0; i < b.len-n-1; i++ {
			cell = cell.prev
		}
		return cell
	}
}
