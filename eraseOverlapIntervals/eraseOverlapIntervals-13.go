package eraseoverlapintervals

import (
	"math"
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	maxcount := 0

	for i := 0; i < len(intervals); i++ {
		count := 1
		nextSeq := intervals[i]
		for j := i + 1; j < len(intervals); j++ {
			if nextSeq[1] <= intervals[j][0] {
				nextSeq = intervals[j]
				count++
			}
			maxcount = int(math.Max(float64(count), float64(maxcount)))
		}
	}
	return len(intervals) - maxcount
}
