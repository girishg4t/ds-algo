package pacificatlantic

import (
	"fmt"
	"strconv"
	"strings"
)

func PacificAtlantic(heights [][]int) [][]int {
	collen := len(heights[0])
	rowlen := len(heights)
	pac := make(map[string]bool)
	atl := make(map[string]bool)

	for column := 0; column < collen; column++ {
		dfs(0, column, pac, heights[0][column], heights)
		dfs(rowlen-1, column, atl, heights[rowlen-1][column], heights)
	}
	for row := 0; row < rowlen; row++ {
		dfs(row, 0, pac, heights[row][0], heights)
		dfs(row, collen-1, atl, heights[row][collen-1], heights)
	}
	result := [][]int{}
	for key, _ := range pac {
		if ok, _ := atl[key]; ok {
			cors := []int{}
			res := strings.Split(key, ",")
			x, _ := strconv.ParseInt(res[0], 0, 8)
			cors = append(cors, int(x))
			y, _ := strconv.ParseInt(res[1], 0, 8)
			cors = append(cors, int(y))
			result = append(result, cors)
		}
	}
	return result
}

func dfs(r, c int, visited map[string]bool, preHeight int, heights [][]int) {
	if r < 0 || c < 0 || r == len(heights) || c == len(heights[0]) ||
		visited[fmt.Sprintf("%d,%d", r, c)] || heights[r][c] < preHeight {
		return
	}

	visited[fmt.Sprintf("%d,%d", r, c)] = true
	dfs(r+1, c, visited, heights[r][c], heights)
	dfs(r-1, c, visited, heights[r][c], heights)
	dfs(r, c+1, visited, heights[r][c], heights)
	dfs(r, c-1, visited, heights[r][c], heights)
}
