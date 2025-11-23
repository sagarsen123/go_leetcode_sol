package main

import (
	"fmt"
	"math"
	"sort"
)

func maxSumDivThree(nums []int) int {
	rem1 := []int{}
	rem2 := []int{}

	t := 0
	for _, num := range nums {
		t += num
		if num%3 == 1 {
			rem1 = append(rem1, num)
		} else if num%3 == 2 {
			rem2 = append(rem2, num)
		}
	}

	r := t % 3

	if r == 0 {
		return t
	}

	ans := 0

	sort.Ints(rem1)
	sort.Ints(rem2)

	r1 := math.MaxInt
	r2 := math.MaxInt
	if r == 1 {
		if len(rem1) > 0 {
			r1 = rem1[0]
		}
		if len(rem2) > 1 {
			r2 = rem2[0] + rem2[1]
		}
	} else {
		if len(rem2) > 0 {
			r2 = rem2[0]
		}
		if len(rem1) > 1 {
			r1 = rem1[0] + rem1[1]
		}
	}
	ans = max(ans, t-min(r1, r2))
	return ans
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
}
