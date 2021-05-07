package linked_list

import (
	"testing"
)

func TestGetIndex0(t *testing.T) {
	l := LinkedList{}
	l.Add("item")

	actual, err := l.Get(0)
	if err != nil {
		t.Errorf("error in getting index 0 from list: %e", err)
	}

	if actual != "item" {
		t.Errorf(`Index 0 = %v, want "item"`, actual)
	}
}

func TestGetIndex3(t *testing.T) {
	l := LinkedList{}
	l.Add("item0")
	l.Add("item1")
	l.Add("item2")
	l.Add("item3")

	actual, err := l.Get(3)
	if err != nil {
		t.Errorf("error in getting index 3 from list: %e", err)
	}

	if actual != "item3" {
		t.Errorf(`Index 0 = %v, want "item3"`, actual)
	}
}
