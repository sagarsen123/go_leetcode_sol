package main

import (
	"fmt"
)

func minNumberOperations(target []int) int {
	ans := 0
	prev := 0
	for _, curr := range target {
		if curr > prev {
			ans += (curr - prev)
		}
		prev = curr
	}
	return ans
}

func main() {
	fmt.Println(minNumberOperations([]int{1, 2, 3, 2, 1}))
}
