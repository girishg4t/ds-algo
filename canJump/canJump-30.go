package canjump

import "math"

//[2,3,1,1,4]
/*func CanJump(nums []int) bool {
	numsmap := make([]bool, len(nums))
	if len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 && !numsmap[len(nums)-1] && !numsmap[i] {
			return false
		}
		if numsmap[len(nums)-1] {
			return true
		}
		min := int(math.Min(float64(nums[i]+i), float64(len(nums)-1)))
		for j := i; j < min; j++ {
			numsmap[j+1] = true
		}
	}
	return numsmap[len(nums)-1]
}*/

func CanJump1(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	numsmap := make([]bool, len(nums))
	lastIndex := len(nums) - 1
	numsmap[0] = true
	if nums[0] == lastIndex {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	for i := lastIndex - 1; i > -1; i-- {
		min := int(math.Min(float64(nums[i]+i), float64(len(nums)-1)))
		for j := i; j < min; j++ {
			numsmap[j+1] = true
		}
	}
	for _, val := range numsmap {
		if !val {
			return false
		}
	}
	return true
}

func CanJump(nums []int) bool {
	jumpIndex := len(nums) - 1
	for i := len(nums) - 2; i > -1; i-- {
		if nums[i]+i >= jumpIndex {
			jumpIndex = i
		}
	}
	return jumpIndex == 0
}
