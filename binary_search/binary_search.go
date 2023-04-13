package binarysearch

// Binary Search
// Dividing and looking whether the value falls under first or second half
// Data must be sorted for binary search
// Time complexity: log(N)
func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		mid := (low + high) / 2
		val := haystack[mid]
		if val == needle {
			return true
		} else if val > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}
