package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert.True(t, BinarySearch([]int{1, 2, 3, 4, 5}, 1))
	assert.True(t, BinarySearch([]int{1, 2, 3, 4, 5}, 2))
	assert.True(t, BinarySearch([]int{1, 2, 3, 4, 5}, 3))
	assert.True(t, BinarySearch([]int{1, 2, 3, 4, 5}, 4))
	assert.True(t, BinarySearch([]int{1, 2, 3, 4, 5}, 5))
	assert.False(t, BinarySearch([]int{1, 2, 3, 4, 5}, 0))
	assert.False(t, BinarySearch([]int{1, 2, 3, 4, 5}, 15))
}
