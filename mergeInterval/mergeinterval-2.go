package mergeinterval

import (
	"math"
	"sort"
)

// func Merge(intervals [][]int) [][]int {
// 	result := [][]int{}
// 	newArr := []int{}
// 	for _, interval := range intervals {
// 		for _, n := range interval {
// 			newArr = append(newArr, n)
// 		}
// 	}
// 	if len(newArr) == 2 {
// 		result = append(result, newArr)
// 	}
// 	start := 0
// 	middle1 := 1
// 	middle2 := 2
// 	end := 3

// 	for end <= len(newArr) {
// 		mergedArr := []int{}
// 		if newArr[start] <= newArr[middle1] &&
// 			newArr[middle1] >= newArr[middle2] &&
// 			newArr[middle2] <= newArr[end] {
// 			min := newArr[start]
// 			max := newArr[end]
// 			for _, n := range newArr[start : start+4] {
// 				min = int(math.Min(float64(min), float64(n)))
// 				max = int(math.Max(float64(max), float64(n)))
// 			}
// 			mergedArr = append(mergedArr, min)
// 			mergedArr = append(mergedArr, max)
// 			result = append(result, mergedArr)
// 		} else {
// 			mergedArr = append(mergedArr, newArr[start])
// 			mergedArr = append(mergedArr, newArr[middle1])
// 			result = append(result, mergedArr)
// 			mergedArr = []int{}
// 			mergedArr = append(mergedArr, newArr[middle2])
// 			mergedArr = append(mergedArr, newArr[end])
// 			result = append(result, mergedArr)
// 		}

// 		start = end + 1
// 		middle1 = start + 1
// 		middle2 = start + 2
// 		end = start + 3
// 	}
// 	return result

// }

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := [][]int{{
		intervals[0][0],
		intervals[0][1],
	}}
	for i := 1; i < len(intervals); i++ {

		count := len(result)
		record := result[count-1]
		firstmin := int(math.Min(float64(intervals[i][0]), float64(record[0])))
		secondMax := int(math.Max(float64(intervals[i][1]), float64(record[1])))

		if (firstmin == record[0] && secondMax == record[1]) ||
			(firstmin == intervals[i][0] && secondMax == intervals[i][1]) ||
			(record[0] <= intervals[i][0] && record[1] >= intervals[i][0]) {
			record[0] = firstmin
			record[1] = secondMax
		} else {
			result = append(result, []int{
				intervals[i][0],
				intervals[i][1],
			})
		}
	}
	return result
}
