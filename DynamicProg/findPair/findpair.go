package findpair

var k = 1
var pair = "()"

var insertPair = "()"

func findPair(n int) []string {
	var i = 1
	var pairArr []string = []string{
		pair,
	}
	for i < n {
		var newArr = []string{}
		for _, str := range pairArr {
			for j, c := range str {
				char := str[:j] + string(c) + pair + str[j+1:]
				if char != str {
					newArr = append(newArr, char)
				}
			}
		}
		pairArr = newArr
		i++
	}
	return pairArr
}
