package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	prev := -1
	for i, curr := range nums {
		if curr == 1 {
			if prev != -1 && i-prev <= k {
				return false
			}
			prev = i
		}
	}
	return true
}

func main() {
	fmt.Println(kLengthApart([]int{1, 0, 0, 0, 1, 0, 0, 1}, 2))
}
