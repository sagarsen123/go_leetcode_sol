package main

import (
	"fmt"
	"math"
)

func solve(energy []int, idx int, s int, k int, t []int) int {
	if idx >= s {
		return 0
	}

	if t[idx] != math.MinInt64 {
		return t[idx]
	}
	t[idx] = energy[idx] + solve(energy, idx+k, s, k, t)
	return t[idx]
}

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	ans := math.MinInt64

	t := make([]int, n)
	for i := range t {
		t[i] = math.MinInt64
	}

	for i := range n {
		ans = max(ans, solve(energy, i, n, k, t))
	}
	return ans
}

func main() {
	energy := []int{5, 2, -10, -5, 1}

	fmt.Println(maximumEnergy(energy, 3))
}
