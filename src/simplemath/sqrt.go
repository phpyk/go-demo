package simplemath

import (
	"math"
)

func Sqrt(val int) int {
	result := math.Sqrt(float64(val))
	return int(result)
}

