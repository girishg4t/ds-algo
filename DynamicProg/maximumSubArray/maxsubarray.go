package maximumsubarray

import (
	"math"
)

var nums = []int{
	-2, 1, -3, 4, -1, 2, 1, -5, 4,
}

/*
1
1
4
3


*/
func FindTotal() int {
	var prev = 0.0
	var max = -1.0

	for i := 0; i < len(nums); i++ {
		prev = math.Max(float64(prev)+float64(nums[i]), float64(nums[i]))
		max = math.Max(max, prev)
	}
	return int(max)
}
