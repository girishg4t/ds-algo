package findmin

import "math"

// 4, 5, 6, 7, 0, 1, 2
func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left < right-1 {
		mid = (left + (right-left)/2)
		if nums[mid] > nums[left] {
			left = mid
		} else {
			right = mid
		}
	}
	return int(math.Min(float64(nums[left]), float64(nums[right])))
}
