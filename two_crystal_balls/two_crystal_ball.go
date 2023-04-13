package twocrystalballs

import (
	"math"
)

// Two Crystal Balls
// Find the closest point where the two balls will break
// if they are droped from certain height
func TwoCrystalBalls(breaks []bool) int {
	// using square root as jmp point
	jmpPoint := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jmpPoint
	for ; i < len(breaks); i += jmpPoint {
		if breaks[i] {
			break
		}
	}

	i -= jmpPoint

	// backtracking to find breaking point of another ball
	for j := 0; j < jmpPoint && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}
