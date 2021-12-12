package twosum

import (
	"math"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
		sort.Ints(nums)
		result = append(result, nums)
		return result
	}

	checkDup := make(map[int]int)
	for _, n := range nums {
		num := -n
		if n < 0 {
			num = int(math.Abs(float64(n)))
		}
		checkDup[num] = n
	}

	newArr := []int{}
	for _, val := range checkDup {
		newArr = append(newArr, val)
	}
	nums = newArr
	arrlen := len(nums) - 1

	for i := 0; i < arrlen-1; i++ {
		for j := i + 1; j < arrlen; j++ {
			val, ok := checkDup[nums[i]+nums[j]]
			if ok {
				newArr := []int{
					nums[i], nums[j], val,
				}
				sort.Ints(newArr)
				result = append(result, newArr)
			}

		}
	}
	return result
}
