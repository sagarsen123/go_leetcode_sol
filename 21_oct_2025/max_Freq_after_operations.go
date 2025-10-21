package main

import (
	"fmt"
	"slices"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	maxVal := slices.Max(nums) + k
	freq := make([]int, maxVal+1)
	ans := 0

	for _, i := range nums {
		freq[i]++
	}

	// calculate cumulative freq
	for i := 1; i <= maxVal; i++ {
		freq[i] += freq[i-1]
	}

	for i := 0; i <= maxVal; i++ {
		if freq[i] == 0 {
			continue
		}

		lnum := max(0, i-k)
		rnum := min(maxVal, i+k)

		totalInRange := freq[rnum]
		if lnum > 0 {
			totalInRange -= freq[lnum-1]
		}

		currElementCount := freq[i]
		if i > 0 {
			currElementCount -= freq[i-1]
		}

		needConversion := totalInRange - currElementCount
		currAns := currElementCount + min(numOperations, needConversion)
		ans = max(ans, currAns)
	}
	return ans
}

func main() {
	fmt.Println(maxFrequency([]int{1, 4, 5}, 1, 2))
}
