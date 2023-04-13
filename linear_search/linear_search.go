package linearsearch

// Linear Search:
// Searching every index, easiest to implement but slowest
// Time complexity: O(N)
func LinearSearch(haystack []int, needle int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
