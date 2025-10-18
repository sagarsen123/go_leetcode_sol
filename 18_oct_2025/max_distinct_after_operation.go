package main

import (
	"fmt"
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)

	answer := 0
	prev := math.MinInt
	n := len(nums)
	for i := 0; i < n; i++ {
		if prev < nums[i]-k {
			prev = nums[i] - k
			answer += 1
		} else if prev < nums[i]+k {
			prev += 1
			answer += 1
		}
	}
	return answer
}
func main() {
	fmt.Println(maxDistinctElements([]int{1, 2, 2, 3, 3, 4}, 2))
}
