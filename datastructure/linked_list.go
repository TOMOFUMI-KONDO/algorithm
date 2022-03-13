package datastructure

import "fmt"

type LinkedList struct {
	base *cellLinkedList
	len  int
}

func NewLinkedList(v int) *LinkedList {
	return &LinkedList{&cellLinkedList{v, nil}, 1}
}

type cellLinkedList struct {
	val  int
	next *cellLinkedList
}

func (l *LinkedList) Get(n int) int {
	return l.get(n).val
}

func (l *LinkedList) Insert(val int, n int) {
	if n > l.len {
		panic(fmt.Errorf("n must be lower or equal than list length(%d), but got %d", l.len, n))
	}
	defer func() { l.len++ }()

	if n == 0 {
		l.base = &cellLinkedList{val, l.base}
		return
	}

	prevOfInserted := l.get(n - 1)
	prevOfInserted.next = &cellLinkedList{val, prevOfInserted.next}
}

func (l *LinkedList) Delete(n int) {
	if n > l.len-1 {
		return
	}
	defer func() { l.len-- }()

	if n == 0 {
		l.base = l.base.next
		return
	}

	prevOfDeleted := l.get(n - 1)
	prevOfDeleted.next = prevOfDeleted.next.next
}

func (l *LinkedList) get(n int) *cellLinkedList {
	cell := l.base
	for i := 0; i < n; i++ {
		cell = cell.next
	}
	return cell
}
