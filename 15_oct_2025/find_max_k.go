package main

import "fmt"

func maxIncreasingSubarrays(nums []int) int {
	k := 0
	n := len(nums)
	curr, prev := 1, 0
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			curr++
		} else {
			prev = curr
			curr = 1
		}
		k = max(k, curr/2)
		k = max(k, min(prev, curr))
	}
	return k
}

func main() {
	fmt.Println(maxIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}))
}
