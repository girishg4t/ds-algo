package recmulty

func Multy(a, b int) int {
	i := 2
	d := b + b
	for i < a {
		d = d + d
		i = i + i
	}
	for i > a {
		d = d - a
		i = i - 1
	}
	return d
}
