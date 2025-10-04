package main

import (
	"container/heap"
	"fmt"
)

// Item represents a cell in the heightMap
type Item struct {
	height int
	row    int
	col    int
}

// MinHeap implements heap.Interface for Items based on height
type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
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

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	n, m := len(heightMap), len(heightMap[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// visited matrix
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	h := &MinHeap{}
	heap.Init(h)

	// push left and right columns
	for i := 0; i < n; i++ {
		for _, j := range []int{0, m - 1} {
			heap.Push(h, Item{heightMap[i][j], i, j})
			visited[i][j] = true
		}
	}

	// push top and bottom rows (excluding corners)
	for j := 1; j < m-1; j++ {
		for _, i := range []int{0, n - 1} {
			heap.Push(h, Item{heightMap[i][j], i, j})
			visited[i][j] = true
		}
	}

	answer := 0

	for h.Len() > 0 {
		curr := heap.Pop(h).(Item)
		height, i, j := curr.height, curr.row, curr.col

		for _, d := range directions {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < n && y >= 0 && y < m && !visited[x][y] {
				visited[x][y] = true
				if height > heightMap[x][y] {
					answer += height - heightMap[x][y]
				}
				heap.Push(h, Item{max(height, heightMap[x][y]), x, y})
			}
		}
	}

	return answer
}

func main() {
	rainWater := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}
	fmt.Println(trapRainWater(rainWater))
}
