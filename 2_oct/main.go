package main

import fmt "fmt"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := 0
	empty := 0

	for numBottles > 0 {
		ans += numBottles
		empty += numBottles
		numBottles = 0
		for empty >= numExchange {
			numBottles += 1
			empty -= numExchange
			numExchange += 1
		}
	}
	return ans
}

func main() {
	fmt.Println(maxBottlesDrunk(13, 6))
}
