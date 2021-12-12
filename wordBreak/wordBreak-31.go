package wordbreak

//"leetcode", wordDict = ["leet","code"]
func WordBreak(s string, wordDict []string) bool {
	str := ""
	if s == "" {
		return true
	}
	for i := 0; i < len(s); i++ {
		str = str + string(s[i])
		for _, word := range wordDict {
			if str == word {
				if !WordBreak(s[i+1:], wordDict) {
					continue
				}
			}
		}

	}
	if s == "" {
		return true
	}
	return false
}
