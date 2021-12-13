package paintfill

import "fmt"

var pixels [][]int = [][]int{{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 0, 0},
	{1, 0, 0, 1, 1, 0, 1, 1},
	{1, 2, 2, 2, 2, 0, 1, 0},
	{1, 1, 1, 2, 2, 0, 1, 0},
	{1, 1, 1, 2, 2, 2, 2, 0},
	{1, 1, 1, 1, 1, 2, 1, 1},
	{1, 1, 1, 1, 1, 2, 2, 1},
}

var newColor = 3

func Fill(x, y int) {
	found := true
	var val = pixels[x][y]
	for found {
		var left = x - 1
		var right = x + 1
		var up = y - 1
		var down = y + 1
		var count = 0
		fmt.Println(pixels)
		if val == pixels[left][y] {
			count++
			pixels[left][y] = newColor
			Fill(left, y)
		}
		if val == pixels[right][y] {
			count++
			pixels[right][y] = newColor
			Fill(right, y)
		}
		if val == pixels[x][up] {
			count++
			pixels[x][up] = newColor
			Fill(x, up)
		}
		if val == pixels[x][down] {
			count++
			pixels[x][down] = newColor
			Fill(x, down)
		}
		if count < 1 {
			found = false
		}
	}
	// fmt.Println(pixels[x][y])
}
