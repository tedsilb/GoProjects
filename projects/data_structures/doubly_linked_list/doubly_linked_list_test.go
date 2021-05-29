package doubly_linked_list

import (
	"testing"
)

func TestLength_0(t *testing.T) {
	l := &DoublyLinkedList{}

	actual := l.Length

	if actual != 0 {
		t.Errorf(`l.Length = %d, want 0`, actual)
	}
}

func TestLength_1_Add(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")

	actual := l.Length

	if actual != 1 {
		t.Errorf(`l.Length = %d, want 1`, actual)
	}
}

func TestLength_1_AddFront(t *testing.T) {
	l := &DoublyLinkedList{}
	l.AddFront("item0")

	actual := l.Length

	if actual != 1 {
		t.Errorf(`l.Length = %d, want 1`, actual)
	}
}

func TestLength_5(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")
	l.AddFront("item1")
	l.Add("item2")
	l.AddFront("item3")
	l.AddFront("item4")

	actual := l.Length

	if actual != 5 {
		t.Errorf(`l.Length = %d, want 5`, actual)
	}
}

func TestAdd_GetIndex0(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item")

	actual, err := l.Get(0)

	if err != nil {
		t.Errorf("error getting index 0 from list: %e", err)
	}
	if *actual != "item" {
		t.Errorf(`Index 0 = %v, want "item"`, actual)
	}
}

func TestAdd_GetIndex3(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")
	l.Add("item1")
	l.Add("item2")
	l.Add("item3")

	actual, err := l.Get(3)

	if err != nil {
		t.Errorf("error getting index 3 from list: %e", err)
	}
	if *actual != "item3" {
		t.Errorf(`Index 3 = %v, want "item3"`, actual)
	}
}

func TestAddFront_GetIndex0(t *testing.T) {
	l := &DoublyLinkedList{}
	l.AddFront("item")

	actual, err := l.Get(0)

	if err != nil {
		t.Errorf("error getting index 0 from list: %e", err)
	}
	if *actual != "item" {
		t.Errorf(`Index 0 = %v, want "item"`, actual)
	}
}

func TestAddFront_GetIndex3(t *testing.T) {
	l := &DoublyLinkedList{}
	l.AddFront("item0")
	l.AddFront("item1")
	l.AddFront("item2")
	l.AddFront("item3")

	actual, err := l.Get(3)

	if err != nil {
		t.Errorf("error getting index 3 from list: %e", err)
	}
	if *actual != "item0" {
		t.Errorf(`Index 3 = %v, want "item0"`, actual)
	}
}

func TestDelete_FirstElement(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")
	l.Add("item1")
	l.Add("item2")
	l.Add("item3")

	err := l.Delete(0)

	if err != nil {
		t.Errorf("error deleting index 0 from list: %e", err)
	}
	elem0, err := l.Get(0)
	if err != nil {
		t.Errorf("error getting index 0 from list: %e", err)
	}
	if *elem0 != "item1" {
		t.Errorf(`index 0 = %v, want "item1"`, elem0)
	}
	if l.first.value != "item1" {
		t.Errorf(`l.first.value = %v, want "item1"`, l.first.value)
	}
}

func TestDelete_OnlyElementOfList(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")

	err := l.Delete(0)

	if err != nil {
		t.Errorf("error deleting index 0 from list: %e", err)
	}
	if l.first != nil {
		t.Errorf(`f.last = %v, want nil`, l.last)
	}
	if l.last != nil {
		t.Errorf(`f.last = %v, want nil`, l.last)
	}
}

func TestDelete_MiddleElement(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")
	l.Add("item1")
	l.Add("item2")
	l.Add("item3")

	err := l.Delete(1)

	if err != nil {
		t.Errorf("error deleting index 1 from list: %e", err)
	}
	actual, err := l.Get(1)
	if err != nil {
		t.Errorf("error getting index 1 from list: %e", err)
	}
	if *actual != "item2" {
		t.Errorf(`index 1 = %v, want "item2"`, l.last)
	}
}

func TestDelete_LastElement(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Add("item0")
	l.Add("item1")
	l.Add("item2")

	err := l.Delete(2)

	if err != nil {
		t.Errorf("error deleting index 2 from list: %e", err)
	}
	if l.last.value != "item1" {
		t.Errorf(`l.last = %v, want "item1"`, l.last.value)
	}
}

func BenchmarkAdd(b *testing.B) {
	l := &DoublyLinkedList{}

	for n := 0; n < b.N; n++ {
		l.Add("item")
	}
}

func BenchmarkAddFront(b *testing.B) {
	l := &DoublyLinkedList{}

	for n := 0; n < b.N; n++ {
		l.AddFront("item")
	}
}

func BenchmarkGet(b *testing.B) {
	l := &DoublyLinkedList{}

	for n := 0; n < b.N; n++ {
		l.Add("item")
	}

	l.Get(b.N - 1)
}
