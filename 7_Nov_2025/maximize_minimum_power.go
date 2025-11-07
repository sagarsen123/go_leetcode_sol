package main

import (
	"fmt"
	"slices"
)

func check(mid int, diff []int, r int, k int, n int) bool {
	temp := slices.Clone(diff)

	currSum := 0

	for i := range n {
		currSum += temp[i]

		if currSum < mid {
			need := mid - currSum
			if need > k {
				return false
			}
			k -= need
			currSum += need

			if i+2*r+1 < n {
				temp[2*r+1+i] -= need
			}

		}
	}
	return true
}

func sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
func maxPower(stations []int, r int, k int) int64 {

	n := len(stations)
	diff := make([]int, n)

	for i := range stations {
		diff[max(0, i-r)] += stations[i]
		if i+r+1 < n {
			diff[i+r+1] -= stations[i]
		}
	}
	left := slices.Min(stations)
	right := sum(stations) + k

	var result int64
	for left <= right {
		mid := left + (right-left)/2

		if check(mid, diff, r, k, n) {
			result = int64(mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func main() {
	stations := []int{1, 2, 4, 5, 0}
	r := 1
	k := 2
	fmt.Println(maxPower(stations, r, k))
}
