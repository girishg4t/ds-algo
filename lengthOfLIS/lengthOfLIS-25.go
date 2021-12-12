package lengthoflis

import (
	"math"
)

func LengthOfLIS(nums []int) int {
	result := make([]int, len(nums))
	index := len(nums) - 1
	result[index] = 1
	max := 1
	for i := index - 1; i > -1; i-- {
		j := i + 1
		for j < len(nums) {
			currentVal := 1
			if nums[i] < nums[j] {
				currentVal = 1 + result[j]
			}

			result[i] = int(math.Max(float64(currentVal), float64(result[i])))
			j++
		}
		max = int(math.Max(float64(result[i]), float64(max)))
	}
	return max
}
