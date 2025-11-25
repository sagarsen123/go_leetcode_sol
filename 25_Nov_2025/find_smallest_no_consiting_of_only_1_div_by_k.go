package main

import "fmt"

func smallestRepunitDivByK(k int) int {
	if k == 1 {
		return 1
	}
	num := 0
	for l := 1; l <= k; l++ {
		num = (num*10 + 1) % k
		if num == 0 {
			return l
		}
	}
	return -1
}
func main() {
	fmt.Println(smallestRepunitDivByK(7))
}
