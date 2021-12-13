package steps

func CalculateSteps(n int) int {

	if n == 1 || n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return CalculateSteps(n-3) + CalculateSteps(n-2) + CalculateSteps(n-1)
}

func ClimbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a := 1
	b := 1
	c := 2
	d := a + b + c
	for i := 2; i < n; i++ {
		a = b
		b = c
		c = d
		d = a + b + c
	}
	return c
}
