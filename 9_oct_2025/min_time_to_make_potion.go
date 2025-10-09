package main

import "fmt"

func minTime(skill []int, mana []int) int64 {
	var n int = len(skill)
	var m int = len(mana)

	potion := make([]int64, n)

	for j := range m {

		potion[0] += int64(mana[j]) * int64(skill[0])

		for i := 1; i < n; i++ {
			potion[i] = max(potion[i], potion[i-1]) + int64(mana[j])*int64(skill[i])
		}

		for i := n - 1; i > 0; i-- {
			potion[i-1] = potion[i] - int64(mana[j])*int64(skill[i])
		}
	}
	return potion[n-1]
}

func main() {
	skill := []int{1, 5, 2, 4}
	mana := []int{5, 1, 4, 2}
	fmt.Println(minTime(skill, mana))
}
