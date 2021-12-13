package powerset

// Initially, one empty subset [[]]
// Adding 1 to []: [[], [1]];
// Adding 2 to [] and [1]: [[], [1], [2], [1, 2]];
// Adding 3 to [], [1], [2] and [1, 2]: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]].

func FindPowerSet(nums []int) [][]int {
	finArr := [][]int{{}}
	for _, e := range nums {
		newArr := [][]int{}
		for _, ie := range finArr {
			ie = append(ie, e)
			newArr = append(newArr, ie)
		}
		finArr = append(finArr, newArr...)
	}
	return finArr
}
