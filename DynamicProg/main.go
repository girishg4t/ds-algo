package main

import (
	"fmt"

	findduplicates "github.com/girishg4t/javascriptTests/DynamicProg/findDuplicates"
)

func main() {
	// i := steps.ClimbStairs(4)
	// i := move.Move(3, 3)
	// a := []int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}
	// i := magicindex.Magicindex(a, 0, len(a))
	// i := powerset.FindPowerSet([]int{1, 2, 3})
	// i := recmulty.Multy(30, 35)
	//i := permutation.Permute()
	//fmt.Println(i)
	//paintfill.Fill(4, 4)
	//i := coin.GetTotal(2, 0)
	// i := maximumsubarray.FindTotal()
	i := findduplicates.ContinuosSubArray([]string{})
	fmt.Println(i)
}

// [1,7,3,6,5,6]
