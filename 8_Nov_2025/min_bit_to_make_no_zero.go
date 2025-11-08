package main

import "fmt"

func minimumOneBitOperations(n int) int {
	ans := n
	ans ^= ans >> 16
	ans ^= ans >> 8
	ans ^= ans >> 4
	ans ^= ans >> 2
	ans ^= ans >> 1
	return ans
}

func main() {
	fmt.Println(minimumOneBitOperations(3))
}
