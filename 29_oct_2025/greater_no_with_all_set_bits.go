package main

import "fmt"

func smallestNumber(n int) int {
	r := 1
	for r < n {
		r = r<<1 | 1
	}
	return r
}
func main() {
	fmt.Println(smallestNumber(5))
}
