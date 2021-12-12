package minmumwindowsubstring

//"ADOBECODEBANC", t = "ABC"
//Input: string = “geeksforgeeks”, pattern = “ork”
func MinWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	store := make(map[rune]int)
	for _, c := range t {
		store[c]++
	}
	refRest := make(map[rune]int)
	copymap(store, refRest)
	str := ""
	min := s
	start := len(t)
	for _, c := range s {
		str = str + string(c)
		if start == len(store) {
			str = string(c)
		}
		if store[c] > 0 {
			store[c]--
			if store[c] == 0 {
				delete(store, c)
			}
		}
		if len(store) == 0 {
			copymap(refRest, store)
			if len(str) < len(min) {
				min = str
			}
			str = ""
		}
	}
	if len(store) > 0 && min == "" {
		return ""
	}
	return min
}

func copymap(src, dest map[rune]int) {
	for k, v := range src {
		dest[k] = v
	}
}
