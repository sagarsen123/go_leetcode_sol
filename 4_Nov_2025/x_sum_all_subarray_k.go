package main

import (
	"fmt"
	"sort"
)

func getTopFreqKeys(mp map[int]int, x int) []int {
	type kv struct {
		key, val int
	}
	pairs := make([]kv, 0, len(mp))

	for k, v := range mp {
		pairs = append(pairs, kv{k, v})
	}

	// Sort by frequency desc, then key desc
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].val == pairs[j].val {
			return pairs[i].key > pairs[j].key
		}
		return pairs[i].val > pairs[j].val
	})

	topKeys := make([]int, 0, x)
	for i := 0; i < len(pairs) && i < x; i++ {
		topKeys = append(topKeys, pairs[i].key)
	}
	return topKeys
}

// findXSum computes the windowed top-x frequency sum
func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}

	mp := make(map[int]int)
	ans := []int{}

	// initialize frequency map for the first window
	for i := 0; i < k; i++ {
		mp[nums[i]]++
	}

	for i := 0; i <= n-k; i++ {
		// Get top x keys
		topKeys := getTopFreqKeys(mp, x)

		curr := 0
		for _, key := range topKeys {
			curr += key * mp[key]
		}
		ans = append(ans, curr)

		// Slide the window
		mp[nums[i]]--
		if mp[nums[i]] == 0 {
			delete(mp, nums[i])
		}
		if i+k < n {
			mp[nums[i+k]]++
		}
	}

	return ans
}

func main() {
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2)) // Example usage
}
