package linked_list

import (
	"testing"
)

func TestAdd_GetIndex0(t *testing.T) {
	l := LinkedList{}
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
	l := LinkedList{}
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
	l := LinkedList{}
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
	l := LinkedList{}
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

func TestLength_0(t *testing.T) {
	l := LinkedList{}

	actual := l.Length

	if actual != 0 {
		t.Errorf(`l.Length = %d, want 0`, actual)
	}
}

func TestLength_1_Add(t *testing.T) {
	l := LinkedList{}
	l.Add("item0")

	actual := l.Length

	if actual != 1 {
		t.Errorf(`l.Length = %d, want 1`, actual)
	}
}

func TestLength_1_AddFront(t *testing.T) {
	l := LinkedList{}
	l.AddFront("item0")

	actual := l.Length

	if actual != 1 {
		t.Errorf(`l.Length = %d, want 1`, actual)
	}
}

func TestLength_5(t *testing.T) {
	l := LinkedList{}
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

func BenchmarkAdd(b *testing.B) {
	l := LinkedList{}

	for n := 0; n < b.N; n++ {
		l.Add("item")
	}
}

func BenchmarkAddFront(b *testing.B) {
	l := LinkedList{}

	for n := 0; n < b.N; n++ {
		l.AddFront("item")
	}
}

func BenchmarkGet(b *testing.B) {
	l := LinkedList{}

	for n := 0; n < b.N; n++ {
		l.Add("item")
	}

	l.Get(b.N - 1)
}
