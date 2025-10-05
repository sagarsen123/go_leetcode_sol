package main

import "fmt"

var row, col int
var directions = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func pacificAtlantic(heights [][]int) [][]int {
	row = len(heights)
	col = len(heights[0])

	a := make([][]bool, row)
	p := make([][]bool, row)

	for i := 0; i < row; i++ {
		a[i] = make([]bool, col)
		p[i] = make([]bool, col)
	}
	var ans [][]int
	var dfs func(x, y int, check [][]bool)
	dfs = func(x, y int, check [][]bool) {
		if check[x][y] {
			return
		}

		check[x][y] = true

		if a[x][y] && p[x][y] {
			ans = append(ans, []int{x, y})
		}

		for _, d := range directions {
			r, c := x+d[0], y+d[1]
			if r >= 0 && r < row && c >= 0 && c < col && heights[r][c] >= heights[x][y] {
				dfs(r, c, check)
			}
		}
	}

	for r := 0; r < row; r++ {
		dfs(r, 0, p)
		dfs(r, col-1, a)
	}

	for c := 0; c < col; c++ {
		dfs(0, c, p)
		dfs(row-1, c, a)
	}
	return ans
}

func main() {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}

	fmt.Println(pacificAtlantic(heights))
}
