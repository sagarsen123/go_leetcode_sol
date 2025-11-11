package main

import "fmt"

type key struct {
	idx, used0, used1 int
}

func solve(idx, used0, used1 int, vals [][]int, mp map[key]int, m, n int) int {
	if idx >= len(vals) {
		return 0
	}

	key := key{idx, used0, used1}
	value, exists := mp[key]
	if exists {
		return value
	}

	ans := solve(idx+1, used0, used1, vals, mp, m, n)

	z, o := vals[idx][0], vals[idx][1]
	if used0+z <= m && used1+o <= n {
		ans = max(ans, 1+solve(idx+1, used0+z, used1+o, vals, mp, m, n))
	}
	mp[key] = ans
	return ans
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	vals := make([][]int, l)
	for idx := range strs {
		z := 0

		for _, ch := range strs[idx] {
			if ch == '0' {
				z++
			}
			vals[idx] = []int{z, len(strs[idx]) - z}
		}
	}

	mp := make(map[key]int)

	return solve(0, 0, 0, vals, mp, m, n)
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
}
