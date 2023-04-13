package twocrystalballs

import (
	"math"
)

// TwoCrystalBalls finds the closest point where two crystal balls will break
// if they are dropped from a certain height. It takes an array of boolean values
// as input, where each boolean value represents whether or not the crystal ball
// breaks at that point. The function returns the index of the closest breaking point
// of the two crystal balls, or -1 if both balls do not break.
// Time complexity: O(sqrt(n)).
func TwoCrystalBalls(breaks []bool) int {
	// using square root as jmp point
	jmpPoint := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	// Starting from the jump point and iterating through each jump point,
	// find the first breaking point of one of the two crystal balls.
	i := jmpPoint
	for ; i < len(breaks); i += jmpPoint {
		if breaks[i] {
			break
		}
	}
	// Backtrack to find the breaking point of the other crystal ball,
	// iterating one element at a time.
	i -= jmpPoint
	for j := 0; j < jmpPoint && i < len(breaks); j, i = j+1, i+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}
