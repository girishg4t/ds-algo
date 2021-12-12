package countsubstrings

// abaab
func CountSubstrings(s string) int {
	result := []string{}
	slen := len(s) - 1
	pos := 0
	index := 0
	data := make(map[string]bool)
	for pos <= slen {
		for index <= slen {
			leftStr := ""
			rightStr := ""
			i := index
			j := i + pos
			if j > slen {
				break
			}
			for i <= j && s[i] == s[j] {
				newStr := s[i : j+1]
				if data[newStr] {
					newStr = leftStr + newStr + rightStr
					result = append(result, newStr)
					break
				}
				leftStr = leftStr + string(s[i])
				if i != j {
					rightStr = string(s[j]) + rightStr
				}
				i++
				j--
			}
			if i >= j {
				str := leftStr + rightStr
				if len(str) > 1 {
					data[str] = true
				}
				result = append(result, leftStr+rightStr)
			}
			index = index + 1
		}
		pos++
		index = 0
	}
	return len(result)
}
