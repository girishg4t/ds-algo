package characterreplacement

import "math"

// "AABABCC" k = 2
func CharacterReplacement(s string, k int) int {
	left := 0
	right := 0
	max := 0
	lettersToReplace := 0
	uniqueCount := 0

	store := make([]int, 26)
	for right = 0; right < len(s); right++ {
		c := rune(s[right]) - rune('A')
		store[c]++
		uniqueCount = int(math.Max(float64(store[c]), float64(uniqueCount)))
		lettersToReplace = right - left + 1 - uniqueCount
		if lettersToReplace > k {
			left++
			i := rune(s[left]) - rune('A')
			store[i]--
		} else {
			max = int(math.Max(float64(max), float64(right-left+1)))
		}
	}
	return max
}
