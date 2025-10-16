package main

import "fmt"

func findSmallestInteger(nums []int, value int) int {
	mp := make([]int, value)
	for _, i := range nums {
		r := ((i % value) + value) % value
		mp[r]++
	}
	ans := 0
	for mp[ans%value] > 0 {
		mp[ans%value]--
		ans++
	}
	return ans
}

func main() {
	fmt.Println(findSmallestInteger([]int{3, 0, 3, 2, 4, 2, 1, 1, 0, 4}, 5))
}
