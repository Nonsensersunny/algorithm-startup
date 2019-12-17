package search

import (
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	start := time.Now()
	index := BinarySearch(data, 8)
	dur := time.Since(start)
	t.Logf("Search:%v got:%d spent:%v", data, index, dur)
}
