package linearsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	assert.Equal(t, true, LinearSearch([]int{1, 2, 3}, 1))
	assert.Equal(t, false, LinearSearch([]int{1, 2, 3}, 4))
}
