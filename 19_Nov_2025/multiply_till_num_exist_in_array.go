package main

import (
	"fmt"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	n := len(nums)
	sort.Ints(nums)
	s := 0
	e := n - 1
	check := true
	for check {
		check = false

		for s <= e {
			mid := s + (e-s)/2
			if nums[mid] == original {
				s = mid + 1
				e = n - 1
				check = true
				break
			} else if nums[mid] > original {
				e = mid - 1
			} else {
				s = mid + 1
			}
		}
		if !check {
			break
		}
		original *= 2
	}
	return original
}

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
}
