package uniquepaths

func UniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	memo[0][0] = 1
	return path(m, n, memo)
}

func path(m, n int, memo [][]int) int {
	for i := 0; i < m; i++ {
		memo[i][0] = 1
	}
	for i := 0; i < n; i++ {
		memo[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i][j] = memo[i-1][j] + memo[i][j-1]
		}
	}
	return memo[m-1][n-1]
}
