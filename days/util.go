package days

import "math"

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}

func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
