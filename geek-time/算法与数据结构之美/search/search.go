package search

func BinarySearch(data []int, target int) int {
	return binarySearch(data, 0, len(data)-1, target)
}

func binarySearch(data []int, l, r, t int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if data[mid] == t {
		return mid
	} else if data[mid] < t {
		return binarySearch(data, mid+1, r, t)
	} else {
		return binarySearch(data, l, mid-1, t)
	}
}
