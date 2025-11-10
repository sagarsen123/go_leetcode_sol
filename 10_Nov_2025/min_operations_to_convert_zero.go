package main

import "fmt"

func minOperations(nums []int) int {
	ans := 0
	stack := []int{0}

	for _, num := range nums {
		for stack[len(stack)-1] > num {
			stack = stack[0 : len(stack)-1]
		}
		if num == 0 {
			continue
		}
		if stack[len(stack)-1] < num {
			stack = append(stack, num)
			ans++
		}
	}
	return ans

}

func main() {
	fmt.Println(minOperations([]int{1, 2, 1, 2, 1, 2}))
}
