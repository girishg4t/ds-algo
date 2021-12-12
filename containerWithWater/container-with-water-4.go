package containerwithwater

import "math"

// 1 8 6 2 5 4 8 3 7
//min = 1, end = 8, end = 8
//min = 1, max =8
//
func MaxArea(height []int) int {
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		min := int(math.Min(float64(height[i]), float64(height[j])))
		max = int(math.Max(float64(max), float64(min*(j-i))))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
