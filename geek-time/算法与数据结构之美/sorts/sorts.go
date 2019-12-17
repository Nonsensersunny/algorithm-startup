package sorts

import (
	"math"
)

// BubbleSort sorts int array in bubble sort
func BubbleSort(data []int) {
	if len(data) < 2 {
		return
	}
	for end := len(data) - 1; end >= 0; end-- {
		altered := false
		for i := 0; i < end; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				altered = true
			}
		}
		if !altered {
			break
		}
	}
}

// InsertionSort sorts int array in insertion sort
func InsertionSort(data []int) {
	if len(data) < 2 {
		return
	}
	for i := 1; i < len(data); i++ {
		j, temp := i-1, data[i]
		for ; j >= 0; j-- {
			if temp < data[j] {
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = temp
	}
}

// SelectionSort sorts int array in selection sort
func SelectionSort(data []int) {
	if len(data) < 2 {
		return
	}
	for i := 0; i < len(data)-1; i++ {
		leastIndex := i
		for j := leastIndex; j < len(data); j++ {
			if data[j] < data[leastIndex] {
				leastIndex = j
			}
		}
		data[i], data[leastIndex] = data[leastIndex], data[i]
	}
}

// MergeSort sorts int array in merge sort
func MergeSort(data []int) {
	if len(data) < 2 {
		return
	}
	mergeSort(data, 0, len(data)-1)
}

func mergeSort(data []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	mergeSort(data, l, mid)
	mergeSort(data, mid+1, r)
	merge(data, l, mid, r)
}

func merge(data []int, left, mid, right int) {
	llen, rlen := mid-left+1, right-mid
	lArr, rArr := make([]int, llen+1), make([]int, rlen+1)

	i, j := 0, 0
	for ; i < llen; i++ {
		lArr[i] = data[left+i]
	}
	lArr[i] = math.MaxInt64

	for ; j < rlen; j++ {
		rArr[j] = data[mid+j+1]
	}
	rArr[j] = math.MaxInt64

	for i, j = 0, 0; left < right+1; left++ {
		if lArr[i] < rArr[j] {
			data[left] = lArr[i]
			i++
		} else {
			data[left] = rArr[j]
			j++
		}
	}
}

// QuickSort sorts int array in quick sort
func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(data, l, r)
	quickSort(data, l, p-1)
	quickSort(data, p+1, r)
}

func partition(data []int, l, r int) int {
	pivot, i := data[r], l
	for j := l; j < r-1; j++ {
		if data[j] < pivot {
			data[j], data[i] = data[i], data[j]
			i++
		}
	}
	data[i], data[r] = data[r], data[i]
	return i
}

// HeapSort sorts int array in heap sort
func HeapSort(data []int) {
	for i := 0; i < len(data); i++ {
		heapify(data, len(data)-i)
		data[0], data[len(data)-i-1] = data[len(data)-i-1], data[0]
	}
}

func heapify(data []int, size int) {
	for i := size; i >= 0; i-- {
		maxHeapify(data, i, size)
	}
}

func maxHeapify(data []int, current, size int) {
	if current >= size {
		return
	}
	l, r, max := current*2+1, current*2+2, current
	if l < size {
		if data[l] > data[max] {
			max = l
		}
	}
	if r < size {
		if data[r] > data[max] {
			max = r
		}
	}
	if current != max {
		data[max], data[current] = data[current], data[max]
		maxHeapify(data, max, size)
	}
}

func minHeapify(data []int, current, size int) {
	if current >= size {
		return
	}
	l, r, min := current*2+1, current*2+2, current
	if l < size {
		if data[l] < data[min] {
			min = l
		}
	}
	if r < size {
		if data[r] < data[min] {
			min = r
		}
	}
	if min != current {
		data[min], data[current] = data[current], data[min]
		minHeapify(data, min, size)
	}
}
