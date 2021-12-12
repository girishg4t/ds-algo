package findwords2

func FindWords(board [][]byte, words []string) []string {
	result := []string{}
	for _, word := range words {
		if findword(word, board) {
			result = append(result, word)
		}
	}
	return result
}

func findword(word string, board [][]byte) bool {
	visited := make(map[int]int)
	return dfs(0, 0, board, visited, word, 0)
}

func dfs(row, column int, board [][]byte, visited map[int]int, word string, index int) bool {
	if row == len(board[0]) || row == -1 ||
		column == len(board) || column == -1 {
		return false
	}
	visited[row] = column
	dfs(row+1, column, board, visited, word, index+1)
	dfs(row-1, column, board, visited, word, index+1)
	dfs(row, column+1, board, visited, word, index+1)
	dfs(row, column-1, board, visited, word, index+1)
	return true
}
