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

// Delete element
func (l *LinkedList) Delete(pos int) error {
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

// PrintList prints list
func (l *LinkedList) PrintList() {
	temp := l.Next
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}
