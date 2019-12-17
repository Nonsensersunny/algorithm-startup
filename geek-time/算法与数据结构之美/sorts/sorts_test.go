package sorts

import (
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	data := []int{3, 1, 6, 4, 7, 5, 8, 9, 2}
	start := time.Now()
	BubbleSort(data)
	dur := time.Since(start)
	t.Logf("Sorted:%v spent:%v", data, dur)
}

func TestInsertionSort(t *testing.T) {
	data := []int{3, 1, 6, 4, 7, 5, 8, 9, 2}
	start := time.Now()
	InsertionSort(data)
	dur := time.Since(start)
	t.Logf("Sorted:%v spent:%v", data, dur)
}

func TestSelectionSort(t *testing.T) {
	data := []int{3, 1, 6, 4, 7, 5, 8, 9, 2}
	start := time.Now()
	SelectionSort(data)
	dur := time.Since(start)
	t.Logf("Sorted:%v spent:%v", data, dur)
}

func TestMergeSort(t *testing.T) {
	data := []int{3, 1, 6, 4, 7, 5, 8, 9, 2}
	start := time.Now()
	MergeSort(data)
	dur := time.Since(start)
	t.Logf("Sorted:%v spent:%v", data, dur)
}

func TestQuickSort(t *testing.T) {
	data := []int{3, 1, 6, 4, 7, 5, 8, 9, 2}
	start := time.Now()
	QuickSort(data)
	dur := time.Since(start)
	t.Logf("Sorted:%v spent:%v", data, dur)
}

func TestHeapSort(t *testing.T) {
	data := []int{3, 1, 6, 4, 7, 5, 8, 9, 2}
	start := time.Now()
	HeapSort(data)
	dur := time.Since(start)
	t.Logf("Sorted:%v spent:%v", data, dur)
}
