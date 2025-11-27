package main

import (
	"fmt"
	"math"
)

func maxi(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefix := make([]int64, n)
	prefix[0] = int64(nums[0])

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + int64(nums[i])
	}

	ans := int64(math.MinInt64)

	for start := 0; start < k; start++ {
		var currSum int64 = 0
		i := start
		for i < n && i+k-1 < n {
			j := i + k - 1
			subSum := prefix[j]
			if i > 0 {
				subSum -= prefix[i-1]
			}
			currSum = maxi(subSum, subSum+currSum)
			ans = maxi(ans, currSum)
			i += k
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSubarraySum([]int{-1, -2, -3, -4, -5}, 4))
}
