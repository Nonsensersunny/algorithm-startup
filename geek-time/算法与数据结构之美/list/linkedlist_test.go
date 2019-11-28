package list

import "testing"

func TestAdd(t *testing.T) {
	li := NewLinkedList()
	li.Add(1)
	li.Add(2)
	li.Add(3)
	li.PrintList()
}

func TestDelete(t *testing.T) {
	li := NewLinkedList()
	li.Add(1)
	li.Add(2)
	li.Add(3)
	li.PrintList()
	if err := li.Delete(1); err != nil {
		t.FailNow()
	}
	li.PrintList()
	_ = li.Delete(1)
	li.PrintList()
}
