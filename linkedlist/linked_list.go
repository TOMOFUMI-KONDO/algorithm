package main

import "fmt"

type Cell struct {
	element int
	next    *Cell
}

func Insert(el int, p *Cell) {
	c := &Cell{element: el, next: p.next}
	p.next = c
}

func Delete(p *Cell) {
	fmt.Println(p.element)
	p.next = p.next.next
}
