package groupanagram

func GroupAnagrams(strs []string) [][]string {
	result := [][]string{}
	return recurse(strs, result)
}
func recurse(strs []string, result [][]string) [][]string {
	if len(strs) == 0 {
		return result
	}
	data := []string{}
	c := strs[0]
	for j := 1; j < len(strs); j++ {
		if isAnagram(c, strs[j]) {
			data = append(data, strs[j])
			if len(strs) > 0 {
				newStr := append(strs[:j], strs[j+1:]...)
				strs = newStr
			}
		}
	}
	data = append(data, c)
	result = append(result, data)

	return recurse(strs[1:], result)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smap := make(map[rune]int)
	tmap := make(map[rune]int)

	for i, c := range s {
		smap[c]++
		tmap[rune(t[i])]++
	}

	for key, val := range smap {
		if tmap[key] != val {
			return false
		}
	}
	return true
}
