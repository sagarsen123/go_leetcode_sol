package main

import "fmt"

func check(nums []int, start int, end int) bool {

	for i := start; i < end-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	for start := 0; start+2*k <= n; start++ {
		left := check(nums, start, start+k)
		right := check(nums, start+k, start+2*k)
		if left && right {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(hasIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3))
}
