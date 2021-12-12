package numdecodings

import "strconv"

// 22260890
func NumDecodings(s string) int {
	arr := make([]int, len(s)+1)

	num, _ := strconv.ParseInt(string(s[0]), 0, 64)
	arr[0] = 1
	if num != 0 {
		arr[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		first := s[i-1 : i]
		num, _ := strconv.ParseInt(first, 0, 64)
		if num > 0 && num < 10 {
			arr[i] = arr[i] + arr[i-1]
		}
		second := s[i-2 : i]
		num, _ = strconv.ParseInt(second, 0, 64)
		if num > 9 && num < 27 {
			arr[i] = arr[i-2] + arr[i]
		}
	}
	return arr[len(s)]
}
