package longestconsecutive

import (
	"math"
)

func LongestConsecutive(nums []int) int {
	seq := make(map[int]bool)
	min := 9999999999999
	max := -999999999999
	for _, n := range nums {
		min = int(math.Min(float64(min), float64(n)))
		max = int(math.Max(float64(max), float64(n)))
		seq[n] = true
	}
	result := 0
	prevresult := 0
	for key, _ := range seq {
		if !seq[key-1] {
			min = key
			for seq[min] {
				result++
				min++
			}
			prevresult = int(math.Max(float64(prevresult), float64(result)))

		}
		result = 0
	}
	return prevresult
}
