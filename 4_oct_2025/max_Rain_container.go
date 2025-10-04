package main

import "fmt"

func maxArea(height []int) int {
	ans := 0
	i := 0
	j := len(height) - 1
	for i < j {
		ans = max(ans, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i += 1
		} else {
			j -= 1
		}
	}
	return ans
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
