package main

import "fmt"

func minCost(colors string, neededTime []int) int {
	i := 0
	ans := 0
	n := len(neededTime)

	for i < n {
		j := i + 1
		for i < n && j < n && colors[i] == colors[j] {
			if neededTime[i] < neededTime[j] {
				ans += neededTime[i]
				i = j
			} else {
				ans += neededTime[j]
			}
			j++
		}
		i = j
	}
	return ans
}

func main() {
	colors := "abaac"
	neededTime := []int{1, 2, 3, 4, 5}
	fmt.Println(minCost(colors, neededTime))
}
