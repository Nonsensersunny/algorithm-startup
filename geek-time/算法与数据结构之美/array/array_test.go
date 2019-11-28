// This file contains tests for array
package array

import (
	"testing"
)

func TestNewArray(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	t.Log(array.Data)
}

func TestAdd(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	if err = array.Add(1); err != nil {
		t.FailNow()
	}
	array.Add(2)
	array.Add(3)
	t.Log(array.Data, array.Size())
}

func TestDeleteNth(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	_ = array.Add(1)
	if err = array.DeleteNth(1); err != nil {
		t.FailNow()
	}
	t.Log(array.Data, array.Size())
}

func TestInsert(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	_ = array.Add(1)
	err = array.Insert(1, 2)
	if err != nil {
		t.FailNow()
	}
	array.Insert(1, 5)
	t.Log(array.Data)
}

func TestReplace(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	_ = array.Add(1)
	_ = array.Add(2)
	if err = array.Replace(1, 13); err != nil {
		t.FailNow()
	}
	array.Replace(2, 100)
	t.Log(array.Data, array.Size())
}

func TestFindX(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	_ = array.Add(1)
	_ = array.Add(2)
	_ = array.Add(3)
	pos, err := array.FindX(5)
	t.Log(pos, err)
}

func TestGetNth(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	_ = array.Add(1)
	_ = array.Add(2)
	_ = array.Add(3)
	pos, err := array.GetNth(1)
	t.Log(pos, err)
}

func TestSize(t *testing.T) {
	array, err := NewArray(10)
	if err != nil {
		t.FailNow()
	}
	_ = array.Add(1)
	_ = array.Add(2)
	_ = array.Add(3)
	t.Log(array.Size())
}
