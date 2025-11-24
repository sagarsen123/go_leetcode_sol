package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	n := len(nums)
	answer := make([]bool, n)
	curr := 0
	for i := range n {
		curr = ((curr << 1) | nums[i]) % 5
		answer[i] = curr == 0
	}
	return answer
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))
}
