/*
Given an array of characters,  return an integer indicating the length of *continuous* subarray with max 2 distinct character

Input: ['A', 'B', 'C', 'A', 'C']
Output: 3
Explanation: the subarray ['C', 'A', 'C'] is the longest  continious subarrary with 2 distict charaters

Input: ['A', 'B', 'B', 'A', 'C', 'A', 'A', 'C', 'B']
Output: 5
Explanation: ['A', 'C', 'A', 'A', 'C']

Input: ['A', 'A', 'A']
Output: 3

*/
//['A', 'B', 'B', 'A', 'C', 'A', 'A', 'C', 'B']

package findduplicates

import "math"

var result = 0
var max = 0
var currentIndex = 0

func ContinuosSubArray(arr []string) int {
	if len(arr) <= 2 || len(arr) < max {
		result = len(arr)
		max = int(math.Max(float64(max), float64(result)))
		return max
	}
	result = 2
	var firstPoint = arr[0]
	var secondPoint = arr[1]

	if firstPoint == secondPoint {
		for _, elem := range arr[2:] {
			if elem != secondPoint {
				secondPoint = elem
				max = int(math.Max(float64(max), float64(result)))
				break
			}
		}
	}

	for _, elem := range arr[2:] {
		if elem == firstPoint || elem == secondPoint {
			result++
		} else {
			max = int(math.Max(float64(max), float64(result)))
			currentIndex++
			return ContinuosSubArray(arr[1:])
		}
	}
	return int(math.Max(float64(max), float64(result)))

}

// n *log(n)
