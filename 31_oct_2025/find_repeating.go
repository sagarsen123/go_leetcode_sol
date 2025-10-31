package main

import "fmt"

func getSneakyNumbers(nums []int) []int {
	mp := make(map[int]int)
	for _, value := range nums {
		mp[value]++
	}
	ans := []int{}
	for key := range mp {
		if mp[key] > 1 {
			ans = append(ans, key)
		}
	}
	return ans
}

func main() {
	fmt.Println(getSneakyNumbers([]int{0, 1, 1, 0}))
}
