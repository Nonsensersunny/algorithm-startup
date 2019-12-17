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

	_ = li.Delete(1)
	li.PrintList()

	_ = li.Delete(1)
	li.PrintList()
}

func TestGetNth(t *testing.T) {
	li := NewLinkedList()
	li.Add(1)
	li.Add(2)
	li.Add(3)
	node, err := li.GetNth(1)
	if err != nil {
		t.Log(err)
	}
	t.Log(node.Data)
}

func TestLength(t *testing.T) {
	li := NewLinkedList()
	li.Add(1)
	li.Add(2)
	li.Add(3)
	li.Delete(1)
	li.Delete(1)
	li.Delete(1)
	li.Delete(1)
	t.Log(li.Length())
}
