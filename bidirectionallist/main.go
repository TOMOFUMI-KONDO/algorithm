package main

type Cell struct {
	val  int
	next *Cell
	prev *Cell
}

func NewCell(val int) Cell {
	return Cell{val: val}
}

func (c *Cell) Val() int {
	return c.val
}

func (c *Cell) Next() *Cell {
	return c.next
}

func (c *Cell) Prev() *Cell {
	return c.prev
}

func (c *Cell) InsertNext(val int) {
	newCell := NewCell(val)
	newCell.next = c.next
	newCell.prev = c
	c.next = &newCell
}

func (c *Cell) InsertPrev(val int) {
	newCell := NewCell(val)
	newCell.prev = c.prev
	newCell.next = c
	c.prev = &newCell
}

func (c *Cell) DeleteNext() {
	if c.next == nil {
		return
	}
	c.next = c.next.next
	c.next.prev = c
}

func (c *Cell) DeletePrev() {
	if c.prev == nil {
		return
	}
	c.prev = c.prev.prev
	c.prev.next = c
}
