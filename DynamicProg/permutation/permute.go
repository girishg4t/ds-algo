package permutation

var i = 0
var orgStr = "ab"
var a = []string{orgStr}

// a b c
// abc bac cba
// acb bca cab
func shuffle() []string {
	var newArr []string
	for _, str := range a {
		j := i
		chars := []rune(str)
		for j < len(str) {
			chars[i], chars[j] = chars[j], chars[i]
			newArr = append(newArr, string(chars))
			j++
		}
	}
	i++
	a = newArr
	if i < len(orgStr)-1 {
		return shuffle()
	}
	return a
}
