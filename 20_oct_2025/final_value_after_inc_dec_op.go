package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	mp := map[string]int{
		"X++": 1,
		"++X": 1,
		"--X": -1,
		"X--": -1,
	}
	ans := 0
	for _, i := range operations {
		ans += mp[i]
	}
	return ans
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
