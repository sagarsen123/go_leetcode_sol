package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0

	for i < n {
		if bits[i] == 0 {
			i++
		} else if bits[i] == 1 && (bits[i+1] == 1 || bits[i+1] == 0) {
			bits[i] = 1
			bits[i+1] = 1
			i += 2
		}
	}
	return bits[n-1] == 0
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0, 0}))
}
