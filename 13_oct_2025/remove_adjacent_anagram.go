package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	runs := []rune(s)
	sort.Slice(runs, func(i, j int) bool {
		return runs[i] < runs[j]
	})
	return string(runs)
}

func removeAnagrams(words []string) []string {
	n := len(words)
	i := 0
	answer := []string{words[0]}

	for i < n {
		j := i + 1
		sorted_i := sortString(words[i])
		for j < n && sorted_i == sortString(words[j]) {
			j++
		}
		if j >= n {
			break
		}
		answer = append(answer, words[j])
		i = j
	}
	return answer
}

func main() {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	fmt.Println(removeAnagrams(words))
}
