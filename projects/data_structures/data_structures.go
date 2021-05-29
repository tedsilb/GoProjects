package main

import (
	"fmt"

	"github.com/tedsilb/GoProjects/projects/data_structures/doubly_linked_list"
	"github.com/tedsilb/GoProjects/projects/data_structures/linked_list"
)

func main() {
	l := linked_list.LinkedList{}
	fmt.Printf("Created linked list: %s\n", l)

	dl := doubly_linked_list.DoublyLinkedList{}
	fmt.Printf("Created doubly linked list: %s\n", dl)
}
