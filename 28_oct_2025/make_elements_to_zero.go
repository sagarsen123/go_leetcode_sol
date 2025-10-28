package main

import (
	"fmt"
	"math"
)

func countValidSelections(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	l := 0
	ans := 0
	for i := range nums {
		r := total - l
		if nums[i] == 0 {
			if l == r {
				ans += 2
			} else if math.Abs(float64(l-r)) == 1 {
				ans += 1
			}
		}
		l += nums[i]
	}
	return ans
}

func main() {
	fmt.Println(countValidSelections([]int{1, 0, 2, 0, 3}))
}
