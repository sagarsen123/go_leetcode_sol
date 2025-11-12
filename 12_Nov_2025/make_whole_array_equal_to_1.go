package main

import (
	"fmt"
	"math"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func minOperations(nums []int) int {
	n := len(nums)

	ans := math.MaxInt

	g := 0
	for _, num := range nums {
		if num == 1 {
			g++
		}
	}

	if g > 0 {
		return n - g
	}

	for i := range n - 1 {
		g = nums[i]
		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				ans = min(ans, j-i)
				break
			}
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans + n - 1
}

func main() {
	fmt.Println(minOperations([]int{2, 6, 3, 4}))
}
