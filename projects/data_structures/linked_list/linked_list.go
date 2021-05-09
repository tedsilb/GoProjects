/* Linked list implementation. */
package linked_list

import (
	"fmt"
)

/* A node in a linked list. */
type Node struct {
	value string
	next  *Node
}

/* Implementation of a singly-linked list. */
type LinkedList struct {
	first  *Node
	last   *Node
	Length uint32
}

/* Add a value to the end of the list. */
func (l *LinkedList) Add(value string) {
	node := &Node{value: value}

	// First element
	if l.first == nil {
		l.first = node
		l.last = node
		l.Length = 1
		return
	}

	l.last.next = node
	l.last = node
	l.Length++
}

/* Add a value to the front of the list. */
func (l *LinkedList) AddFront(value string) {
	node := &Node{value: value}

	// First element
	if l.first == nil {
		l.first = node
		l.last = node
		l.Length = 1
		return
	}

	node.next = l.first
	l.first = node
	l.Length++
}

/* Get the value at the provided index. */
func (l *LinkedList) Get(index int) (r *string, e error) {
	// TODO: make negative indexing work
	if index == 0 {
		return &l.first.value, nil
	}

	prev := l.first
	for i := 0; i < index; i++ {
		if prev.next == nil {
			return nil, fmt.Errorf("element %d does not exist in the list", i)
		}
		prev = prev.next
	}

	return &prev.value, nil
}
