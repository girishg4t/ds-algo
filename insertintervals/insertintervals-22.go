package insertintervals

import "math"

// 1 2 3 5 6 9 => 1 5 6 9
// 1 2 3 5 6 7 8 10 12 16 =>  1 2 3 4 5 6 7 8 10 12 16 =>  1 2 3 10 12 16
func Insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}
	for _, interval := range intervals {
		if newInterval == nil {
			result = append(result, interval)
			continue
		}
		if interval[1] < newInterval[0] {
			result = append(result, interval)
			continue
		} else if newInterval == nil || interval[0] > newInterval[1] {
			result = append(result, newInterval)
			result = append(result, interval)
			newInterval = nil
			break
		} else {
			min := int(math.Min(float64(interval[0]), float64(newInterval[0])))
			max := int(math.Max(float64(interval[1]), float64(newInterval[1])))
			newInterval[0] = min
			newInterval[1] = max
		}
	}
	if newInterval != nil {
		result = append(result, newInterval)
	}
	return result
}
