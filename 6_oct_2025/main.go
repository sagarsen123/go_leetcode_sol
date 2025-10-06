package main

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	t int
	x int
	y int
}

type minH []Cell

func (h minH) Len() int           { return len(h) }
func (h minH) Less(i, j int) bool { return h[i].t < h[j].t }
func (h minH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minH) Push(x any)        { *h = append(*h, x.(Cell)) }
func (h *minH) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func swimInWater(grid [][]int) int {
	n := len(grid)
	directions := [][]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, n)
	}

	minh := &minH{{grid[0][0], 0, 0}}
	visited[0][0] = true

	heap.Init(minh)

	for minh.Len() > 0 {
		cell := heap.Pop(minh).(Cell)
		t, x, y := cell.t, cell.x, cell.y

		if x == n-1 && y == n-1 {
			return t
		}

		for _, d := range directions {
			r, c := x+d[0], y+d[1]
			if r >= 0 && r < n && c >= 0 && c < n && !visited[r][c] {
				visited[r][c] = true
				heap.Push(minh, Cell{max(t, grid[r][c]), r, c})
			}
		}
	}
	return -1
}

func main() {
	grid := [][]int{
		{0, 2},
		{1, 3},
	}
	fmt.Println(swimInWater(grid)) // Output: 3
}
