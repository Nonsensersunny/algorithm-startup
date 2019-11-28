// Array emulation

// Package array contains array operations
package array

import (
	"errors"
)

var errList = map[string]string{
	"E001": "index out of bound",
	"E002": "full array",
	"E003": "invalid length",
	"E004": "target not found",
}

// Array struct
type Array struct {
	Length int
	Data   []int
	count  int
}

// NewArray returns a new array of n ints
func NewArray(n int) (*Array, error) {
	if n < 1 {
		return nil, errors.New(errList["E001"])
	}
	return &Array{
		Length: n,
		Data:   make([]int, n),
		count:  0,
	}, nil
}

// Add adds element of array
func (a *Array) Add(num int) error {
	if a.count == a.Length {
		return errors.New(errList["E002"])
	}
	a.Data[a.count] = num
	a.count++
	return nil
}

// DeleteNth delete nth number of array
func (a *Array) DeleteNth(n int) error {
	if n > a.count || n < 1 {
		return errors.New(errList["E004"])
	}
	for ; n < a.count-1; n++ {
		a.Data[n-1] = a.Data[n]
	}
	a.count--
	return nil
}

// Insert insert element
func (a *Array) Insert(pos, num int) error {
	if pos < 1 || pos > a.count {
		return errors.New(errList["E004"])
	}
	if a.count == a.Length {
		return errors.New(errList["E002"])
	}
	for i := a.count; i > pos-1; i-- {
		a.Data[i] = a.Data[i-1]
	}
	a.Data[pos-1] = num
	a.count++
	return nil
}

// Replace replace element
func (a *Array) Replace(pos, val int) error {
	if pos < 1 || pos > a.count {
		return errors.New(errList["E004"])
	}
	a.Data[pos-1] = val
	return nil
}

// FindX find target element
func (a *Array) FindX(x int) (int, error) {
	for i := 0; i < a.count; i++ {
		if a.Data[i] == x {
			return i + 1, nil
		}
	}
	return -1, errors.New(errList["E004"])
}

// GetNth get nth element
func (a *Array) GetNth(n int) (int, error) {
	if n >= a.count || n < 1 {
		return -1, errors.New(errList["E001"])
	}
	return a.Data[n-1], nil
}

// Size returns count
func (a *Array) Size() int {
	return a.count
}
