package icpatest

import (
	"fmt"
)

func SolveTest(str string, strs []string) int {
	count := 0
	charMap := make(map[string]bool)
	for _, ch := range str {
		charMap[string(ch)] = true
	}
	for _, s := range strs {
		found := true
		for _, c := range s {
			if !charMap[string(c)] {
				found = false
				break
			}
		}
		if found {
			count++
		}
	}
	return count
}

func miniMaxSum(arr []int32) {
	min := 9223372036854775807
	max := -9223372036854775807
	minindex := 0
	maxindex := 0
	for i, n := range arr {
		if int(n) <= min {
			min = int(n)
			minindex = i
		}

		if int(n) >= max {
			max = int(n)
			maxindex = i
		}
	}
	totalmin := 0
	totalmax := 0

	for i, n := range arr {
		if i != maxindex {

			totalmin = totalmin + int(n)
		}
		if i != minindex {
			totalmax = totalmax + int(n)
		}
	}
	fmt.Println(totalmin, totalmax)
}
