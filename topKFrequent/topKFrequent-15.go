package topkfrequent

import "math"

func topKFrequent(nums []int, k int) []int {
	elems := make(map[int]int)
	result := make([]int, k)
	for _, n := range nums {
		elems[n]++
	}
	for i := 0; i < k; i++ {
		result[i] = math.MinInt64
		foundkey := nums[0]
		for key, val := range elems {
			if val > result[i] {
				result[i] = val
				foundkey = key
			}
		}
		result[i] = foundkey
		delete(elems, result[i])
	}

	return result
}
