package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head  *node
	lenth int
}

func (l *linkedList) add(value int) {
	n := &node{value, nil}
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.lenth++
}

func (l *linkedList) prepend(n *node) {
	n.next = l.head
	l.head = n
}

func (l *linkedList) print() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (l *linkedList) remove(value int) {
	current := l.head
	if current.value == value {
		l.head = current.next
		l.lenth--
		return
	}
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.lenth--
			return
		}
		current = current.next
	}
}

func main() {
	l := &linkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	l.add(4)
	l.add(5)

	l.prepend(&node{0, nil})

	l.remove(4)

	l.print()
}
