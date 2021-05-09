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

func (n Node) String() string {
	return fmt.Sprintf("Node(%s)", n.value)
}

/* Implementation of a singly-linked list. */
type LinkedList struct {
	first  *Node
	last   *Node
	Length uint32
}

func (l LinkedList) String() string {
	return fmt.Sprintf("LinkedList(Length=%d)", l.Length)
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
	node, err := l.get(index)
	if err != nil {
		return nil, err
	}
	return &node.value, nil
}

/* Get the node at the provided index. */
func (l *LinkedList) get(index int) (r *Node, e error) {
	if index == 0 {
		return l.first, nil
	}

	curr := l.first
	for i := 0; i < index; i++ {
		if curr.next == nil {
			return nil, fmt.Errorf("element %d does not exist in the list", i)
		}
		curr = curr.next
	}

	return curr, nil
}

/* Delete the node at the provided index. */
func (l *LinkedList) Delete(index int) error {
	if l.first == nil || l.last == nil {
		return fmt.Errorf("element %d does not exist in the list (list is empty)", index)
	}

	if index == 0 {
		if l.first.next == nil {
			if *l.last == *l.first {
				l.last = nil
			}
			l.first = nil
		} else {
			l.first = l.first.next
		}

	} else if index == int(l.Length-1) {
		secondLast, err := l.get(int(l.Length - 2))
		if err != nil {
			return fmt.Errorf("unable to retrieve element %d from the list "+
				"(this should never happen)", l.Length-2)
		}

		l.last = secondLast
		secondLast.next = nil

	} else {
		var prev *Node
		curr := l.first
		for i := 0; i < index; i++ {
			if curr.next == nil {
				return fmt.Errorf("element %d does not exist in the list", i)
			}
			prev = curr
			curr = curr.next
		}

		prev.next = curr.next
	}

	l.Length--
	return nil
}
