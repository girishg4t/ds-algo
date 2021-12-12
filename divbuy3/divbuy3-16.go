package divbuy3

import (
	"fmt"
	"strconv"
)

func Solution(S string) int {
	count := 0
	for start, ch := range S {
		for i := 0; i < 10; i++ {
			if string(ch) == fmt.Sprintf("%d", i) {
				continue
			}
			s := fmt.Sprintf("%s%d%s", string(S[:start]), i, string(S[start+1:]))
			n, _ := strconv.ParseInt(s, 10, 64)
			fmt.Println(n)
			if n%3 == 0 {
				count++
			}
		}
	}
	n, _ := strconv.ParseInt(S, 10, 64)
	if n%3 == 0 {
		count++
	}
	return count
}
