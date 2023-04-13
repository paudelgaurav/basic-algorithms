package twocrystalballs

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)
	data := make([]bool, 10000)
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	assert.Equal(t, idx, TwoCrystalBalls(data))
	assert.Equal(t, -1, TwoCrystalBalls(make([]bool, 821)))
}
