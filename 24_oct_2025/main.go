package main

import "fmt"

func is_balanced(n int) bool {
	mp := make(map[int]int)
	for n > 0 {
		mp[n%10]++
		n /= 10
	}

	for key, value := range mp {
		if key != value {
			return false
		}
	}
	return true
}
func nextBeautifulNumber(n int) int {
	n++
	for is_balanced(n) == false {
		n++
	}
	return n
}
func main() {
	fmt.Println(nextBeautifulNumber(1000))
}
