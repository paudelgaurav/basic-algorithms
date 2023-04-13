package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	assert.True(t, LinearSearch([]int{1, 2, 3}, 1))
	assert.False(t, LinearSearch([]int{1, 2, 3}, 4))
}
