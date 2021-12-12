package numberofislands

/*Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
*/
func NumIslands(grid [][]byte) int {
	rowlen := len(grid)
	collen := len(grid[0])
	land := 0

	for row := 0; row < rowlen; row++ {
		for column := 0; column < collen; column++ {
			if string(grid[row][column]) == "1" {
				dfs(row, column, grid)
				land++
			}
		}
	}

	return land
}

func dfs(r, c int, grid [][]byte) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || string(grid[r][c]) != "1" {
		return
	}
	grid[r][c] = byte(50)

	dfs(r+1, c, grid)
	dfs(r-1, c, grid)
	dfs(r, c+1, grid)
	dfs(r, c-1, grid)
}
