package main

import "fmt"

func minimumOperations(nums []int) int {
	ans := 0

	for _, val := range nums {
		if val%3 == 0 {
			continue
		}
		ans += 1
	}
	return ans
}

func main() {
	fmt.Println(minimumOperations([]int{1, 2, 3, 4}))
}
