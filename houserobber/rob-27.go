package houserobber

func Rob(nums []int) int {
	max := 0
	pos := 0
	for i, n := range nums {
		if n > max {
			max = n
			pos = i
		}
	}

}

func findLeft(pos int, nums []int) int {
	if len(nums[:pos]) == 2 {
		if nums[pos+1] > nums[pos+2] {

		}
	}
}

func findRight(pos int, nums []int) {
	if len(nums[pos:]) == 2 {

	}
}
