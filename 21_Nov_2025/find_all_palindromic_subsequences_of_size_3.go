package main

import "fmt"

func countPalindromicSubsequence(s string) int {
	indices := make([][]int, 26)
	for i := range 26 {
		indices[i] = []int{-1, -1}
	}
	result := 0
	n := len(s)

	for i := range n {
		idx := s[i] - 'a'
		if indices[idx][0] == -1 {
			indices[idx][0] = i
		}
		indices[idx][1] = i
	}

	for idx := range 26 {
		if indices[idx][0] == -1 {
			continue
		}

		mp := make(map[rune]bool)
		for j := indices[idx][0] + 1; j < indices[idx][1]; j++ {
			ch := rune(s[j])
			mp[ch] = true
		}
		result += len(mp)
	}

	return result
}

func main() {
	fmt.Println(countPalindromicSubsequence("aabca"))
}
