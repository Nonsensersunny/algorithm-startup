// LinkedList emulation

// Package list contains list operations
package list

import (
	"errors"
	"fmt"
)

var errList = map[string]string{
	"E001": "index out of bound",
	"E002": "full array",
	"E003": "invalid length",
	"E004": "target not found",
}

// LinkedList struct
type LinkedList struct {
	Data int
	Next *LinkedList
}

// NewLinkedList returns a new linkedlist
func NewLinkedList() *LinkedList {
	return &LinkedList{
		Next: nil,
	}
}

// Add node
func (l *LinkedList) Add(data int) {
	temp := &LinkedList{
		Data: data,
		Next: l.Next,
	}
	l.Next = temp
}

// AddBefore insert data before pos
func (l *LinkedList) AddBefore(pos, data int) error {
	return nil
}

// Delete element
func (l *LinkedList) Delete(pos int) error {
	if l.Next == nil {
		return nil
	}
	temp := l
	for ; pos > 0; pos-- {
		temp = temp.Next
		if l == nil {
			return errors.New(errList["E001"])
		}
	}
	l.Next = l.Next.Next
	return nil
}

// GetNth returns nth node
func (l *LinkedList) GetNth(n int) (*LinkedList, error) {
	if l.Next == nil {
		return nil, errors.New(errList["E001"])
	}
	temp := l
	for ; n > 0; n-- {
		temp = temp.Next
		if temp == nil {
			return nil, errors.New(errList["E001"])
		}
	}
	return temp, nil
}

//Length return the length of the list
func (l *LinkedList) Length() int {
	if l.Next == nil {
		return 0
	}
	cnt, tmp := -1, l
	for tmp != nil {
		cnt++
		tmp = tmp.Next
	}
	return cnt
}

// PrintList prints list
func (l *LinkedList) PrintList() {
	if l.Next == nil {
		fmt.Println("empty list")
		return
	}
	temp := l.Next
	fmt.Println("-----START-----")
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
	fmt.Println("-----END-----")
}
