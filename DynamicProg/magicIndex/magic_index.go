package magicindex

func Magicindex(a []int, low, high int) int {
	m := (low + high) / 2
	for low <= high {
		if a[m] == m {
			return m
		}
		if m < a[m] {
			return Magicindex(a, low, m-1)
		}
		if m > a[m] {
			return Magicindex(a, m+1, high)
		}
	}

	return -1
}
