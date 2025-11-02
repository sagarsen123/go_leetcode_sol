package main

import "fmt"

func update_grid(m int, n int, x int, y int, grid [][]int) {
	if 0 <= x && x < m && 0 <= y && y < n {
		for r := y + 1; r < n && grid[x][r] != 1 && grid[x][r] != 2; r++ {
			grid[x][r] = 3
		}
		for l := y - 1; l >= 0 && grid[x][l] != 1 && grid[x][l] != 2; l-- {
			grid[x][l] = 3
		}
		for d := x + 1; d < m && grid[d][y] != 1 && grid[d][y] != 2; d++ {
			grid[d][y] = 3
		}
		for u := x - 1; u >= 0 && grid[u][y] != 1 && grid[u][y] != 2; u-- {
			grid[u][y] = 3
		}
	}
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)

	for i := range grid {
		grid[i] = make([]int, n)
	}

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = 1
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 2
	}

	for _, guard := range guards {
		update_grid(m, n, guard[0], guard[1], grid)
	}
	answer := 0
	for row := range m {
		for col := range n {
			if grid[row][col] == 0 {
				answer++
			}
		}
	}

	return answer
}

func main() {
	walls := [][]int{{0, 1}, {2, 2}, {1, 4}}
	guards := [][]int{{0, 0}, {1, 1}, {2, 3}}
	fmt.Println(countUnguarded(4, 6, guards, walls))
}
