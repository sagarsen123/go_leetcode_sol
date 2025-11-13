package main

import "fmt"

func maxOperations(s string) int {
	ans := 0
	ones := 0
	for i, ch := range s {
		if ch == '1' {
			ones++
		} else if i > 0 && s[i] < s[i-1] {
			ans += ones
		}
	}
	return ans
}

func main() {
	fmt.Println(maxOperations("1001101"))
}
