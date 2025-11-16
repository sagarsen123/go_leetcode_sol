package main

import "fmt"

func numSub(s string) int {
	const mod = 1_000_000_007

	ans := 0
	curr := 0

	for _, ch := range s {
		if ch == '1' {
			curr++
			ans = (ans + curr) % mod
		} else {
			curr = 0
		}
	}

	return ans
}

func main() {
	fmt.Println(numSub("111111"))
}
