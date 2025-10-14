package main

import "fmt"

func check(nums []int, start int, k int, n int) bool {
	if n-start < k {
		return false
	}
	for i := start; i < start+k-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	for i := range n {
		left := check(nums, i, k, n)
		right := check(nums, i+k, k, n)
		if left && right {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(hasIncreasingSubarrays([]int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}, 3))
}
