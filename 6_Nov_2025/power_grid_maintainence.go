package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ---------------- DFS for components ----------------

func dfs(node, compID int, graph map[int][]int, visited []bool, compMap map[int]int, compNodes map[int][]int) {
	if visited[node] {
		return
	}
	visited[node] = true
	compMap[node] = compID
	compNodes[compID] = append(compNodes[compID], node)
	for _, nbr := range graph[node] {
		dfs(nbr, compID, graph, visited, compMap, compNodes)
	}
}

// ---------------- Main logic ----------------

func processQueries(c int, connections [][]int, queries [][]int) []int {
	graph := make(map[int][]int)
	for _, e := range connections {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, c+1)
	compMap := make(map[int]int)     // node -> component
	compNodes := make(map[int][]int) // component -> list of nodes

	for i := 1; i <= c; i++ {
		if !visited[i] {
			dfs(i, i, graph, visited, compMap, compNodes)
		}
	}

	// Initialize per-component heap and active set
	compHeap := make(map[int]*IntHeap)
	active := make(map[int]map[int]bool)
	for comp, nodes := range compNodes {
		h := &IntHeap{}
		for _, node := range nodes {
			heap.Push(h, node)
		}
		heap.Init(h)
		compHeap[comp] = h

		active[comp] = make(map[int]bool)
		for _, node := range nodes {
			active[comp][node] = true
		}
	}

	var result []int
	for _, q := range queries {
		qType, node := q[0], q[1]
		compID := compMap[node]

		if qType == 1 {
			if active[compID][node] {
				result = append(result, node)
			} else {
				h := compHeap[compID]
				found := -1
				// Lazy removal of inactive nodes
				for h.Len() > 0 {
					top := (*h)[0]
					if !active[compID][top] {
						heap.Pop(h)
						continue
					}
					found = top
					break
				}
				result = append(result, found)
			}
		} else {
			// remove node
			active[compID][node] = false
		}
	}

	return result
}

func main() {
	c := 5
	conections := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	queries := [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
	fmt.Println(processQueries(c, conections, queries))
}
