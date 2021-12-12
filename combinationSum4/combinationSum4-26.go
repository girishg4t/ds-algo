package combinationsum4

func CombinationSum4(nums []int, target int) int {
	numbermap := make(map[int]int)
	numbermap[0] = 1
	for i := 1; i < target+1; i++ {
		numbermap[i] = 0
		for _, n := range nums {
			if val, ok := numbermap[i-n]; ok {
				numbermap[i] = numbermap[i] + val
			}
		}
	}

	return numbermap[target]
}
