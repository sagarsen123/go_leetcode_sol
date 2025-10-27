package main

import (
	"fmt"
	"strings"
)

func numberOfBeams(bank []string) int {
	i := 0
	n := len(bank)
	ans := 0
	for i < n {
		for i < n && !strings.Contains(bank[i], "1") {
			i++
		}

		if i >= n {
			break
		}

		j := i + 1
		for j < n && !strings.Contains(bank[j], "1") {
			j++
		}
		if j >= n {
			break
		}
		count_r1 := strings.Count(bank[i], "1")
		count_r2 := strings.Count(bank[j], "1")
		ans += (count_r1 * count_r2)
		i++
	}
	return ans
}

func main() {
	fmt.Println(numberOfBeams([]string{"011001", "000000", "010100", "001000"}))
}
