package main

import (
	"fmt"
	"math"
)

func minNumberOperations(target []int) int {
	ans := 0
	prev := 0
	for _, curr := range target {
		if math.Abs(float64(curr)) > math.Abs(float64(prev)) {
			ans += int(math.Abs(float64(curr)) - math.Abs(float64(prev)))
		}
		prev = curr
	}
	return ans
}

func main() {
	fmt.Println(minNumberOperations([]int{1, 2, 3, 2, 1}))
}
