package main

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))
	p := len(potions)

	sort.Ints(potions)

	for i, spell := range spells {
		s := 0
		e := p - 1

		for s <= e {
			m := s + (e-s)/2
			product := int64(spell) * int64(potions[m])
			if product >= success {
				e = m - 1
			} else {
				s = m + 1
			}
		}
		ans[i] = p - s
	}
	return ans
}

func main() {
	spells := []int{10, 20, 30, 40}
	potions := []int{1, 2, 3, 4, 5}
	success := int64(100)

	fmt.Println(successfulPairs(spells, potions, success)) // [0 1 2]
}
