package main

import (
	"container/list"
	"fmt"
)

func findLexSmallestString(s string, a, b int) string {
	st := make(map[string]bool)
	q := list.New()

	st[s] = true
	q.PushBack(s)
	ans := s

	for q.Len() > 0 {
		curr := q.Remove(q.Front()).(string)

		// update answer if lexicographically smaller
		if curr < ans {
			ans = curr
		}

		// ---------- Addition Operation ----------
		temp := []byte(curr)
		for i := 1; i < len(temp); i += 2 {
			temp[i] = byte(((int(temp[i]-'0') + a) % 10) + '0')
		}
		addStr := string(temp)
		if !st[addStr] {
			st[addStr] = true
			q.PushBack(addStr)
		}

		// ---------- Rotation Operation ----------
		n := len(curr)
		rotStr := curr[n-b:] + curr[:n-b]
		if !st[rotStr] {
			st[rotStr] = true
			q.PushBack(rotStr)
		}
	}

	return ans
}
func main() {
	fmt.Println(findLexSmallestString("5525", 9, 2))
}
