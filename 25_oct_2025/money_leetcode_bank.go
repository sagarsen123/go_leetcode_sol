package main

import "fmt"

func totalMoney(n int) int {
	full_week := n / 7
	remaing_days := n % 7

	answer := full_week*28 + (7 * full_week * (full_week - 1) / 2)
	full_week++
	for day := 0; day < remaing_days; day++ {
		answer += (full_week + day)
	}
	return answer
}

func main() {
	// This is a placeholder main function.
	fmt.Println(totalMoney(20))
}
