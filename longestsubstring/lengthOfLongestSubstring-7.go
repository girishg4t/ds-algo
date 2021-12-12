package longestsubstring

import (
	"math"
)

func LengthOfLongestSubstring(s string) int {
	store := make(map[string]bool)
	max := 0
	count := 0
	for i, ch := range s {
		c := string(ch)
		if store[c] {
			store = make(map[string]bool)
			store[c] = true
			count = backtrack(i-1, s, store)
		} else {
			store[c] = true
			count++
		}

		max = int(math.Max(float64(max), float64(count)))
	}
	return max
}

func backtrack(i int, s string, store map[string]bool) int {
	count := 1
	for !store[string(s[i])] {
		count++
		store[string(s[i])] = true
		i--
	}
	return count
}
