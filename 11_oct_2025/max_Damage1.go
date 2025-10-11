package main

import (
	"fmt"
	"sort"
)

func solve(nums []int64, idx int, mp map[int64]int64, dp []int64) int64 {
	if idx >= len(nums) {
		return 0
	}

	if dp[idx] != -1 {
		return dp[idx]
	}

	skip := solve(nums, idx+1, mp, dp)

	j := sort.Search(len(nums), func(k int) bool {
		return nums[k] > nums[idx]+2
	})

	curr := nums[idx]*mp[nums[idx]] + solve(nums, j, mp, dp)

	dp[idx] = max(skip, curr)
	return dp[idx]
}

func maximumTotalDamage1(power []int) int64 {
	mp := make(map[int64]int64)
	for _, p := range power {
		mp[int64(p)]++
	}

	// Collect unique sorted numbers
	nums := make([]int64, 0, len(mp))
	for k := range mp {
		nums = append(nums, k)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	dp := make([]int64, len(nums))
	for i := range dp {
		dp[i] = -1
	}

	return solve(nums, 0, mp, dp)
}

func maximumTotalDamage2(power []int) int64 {
	// Frequency map
	mp := make(map[int64]int64)
	for _, p := range power {
		mp[int64(p)]++
	}

	// Unique sorted numbers
	nums := make([]int64, 0, len(mp))
	for k := range mp {
		nums = append(nums, k)
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	n := len(nums)
	dp := make([]int64, n)
	for i := range dp {
		dp[i] = -1
	}

	var max func(a, b int64) int64
	max = func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	// Recursive closure
	var solve func(idx int) int64
	solve = func(idx int) int64 {
		if idx >= n {
			return 0
		}
		if dp[idx] != -1 {
			return dp[idx]
		}

		// Option 1: skip current
		skip := solve(idx + 1)

		// Option 2: take current and jump
		j := sort.Search(n, func(k int) bool { return nums[k] > nums[idx]+2 })
		take := nums[idx]*mp[nums[idx]] + solve(j)

		dp[idx] = max(skip, take)
		return dp[idx]
	}

	return solve(0)
}

func main() {
	fmt.Println(maximumTotalDamage1([]int{1, 1, 3, 4}))
	fmt.Println(maximumTotalDamage2([]int{1, 1, 3, 4}))
}
