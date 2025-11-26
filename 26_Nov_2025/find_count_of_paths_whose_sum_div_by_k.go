package main

import "fmt"

var mod = 1000000007
var K int
var m, n int

func solve(grid [][]int, x, y, currSum int, mp map[[3]int]int) int {
	if x >= m || y >= n {
		return 0
	}

	currSum = (currSum + grid[x][y]) % K
	if x == m-1 && y == n-1 {
		if currSum%K == 0 {
			return 1
		}
		return 0
	}

	key := [3]int{x, y, currSum}
	if val, ok := mp[key]; ok {
		return val
	}

	down := solve(grid, x, y+1, currSum, mp)
	right := solve(grid, x+1, y, currSum, mp)

	mp[key] = (down + right) % mod
	return mp[key]

}

func numberOfPaths(grid [][]int, k int) int {
	K = k
	m, n = len(grid), len(grid[0])
	mp := make(map[[3]int]int)
	return solve(grid, 0, 0, 0, mp)
}

func main() {
	grid := [][]int{
		{5, 2, 4},
		{3, 0, 5},
		{0, 7, 2},
	}
	k := 3
	fmt.Println(numberOfPaths(grid, k))
}
