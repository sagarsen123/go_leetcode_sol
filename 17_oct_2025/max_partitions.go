package main

import "fmt"

var s string
var k int
var mp map[int]int

func solve(i int, unique int, canchange bool) int {
	key := (i << 27) | (unique << 1)
	if canchange {
		key |= 1
	}

	if val, ok := mp[key]; ok {
		return val
	}

	if i >= len(s) {
		return 0
	}

	currCharIdx := int(s[i] - 'a')
	new_unique := unique | (1 << currCharIdx)
	unique_count := bitsCount(new_unique)
	result := 0

	if unique_count > k {
		result = 1 + solve(i+1, 1<<currCharIdx, canchange)
	} else {
		result = solve(i+1, new_unique, canchange)
	}

	if canchange {
		for new_idx := 0; new_idx < 26; new_idx++ {
			new_unique_set := unique | (1 << new_idx)
			new_unique_count := bitsCount(new_unique_set)
			if new_unique_count > k {
				tmp := 1 + solve(i+1, 1<<new_idx, false)
				if tmp > result {
					result = tmp
				}
			} else {
				tmp := solve(i+1, new_unique_set, false)
				if tmp > result {
					result = tmp
				}
			}
		}
	}

	mp[key] = result
	return result
}

func maxPartitionsAfterOperations(str string, kk int) int {
	s = str
	k = kk
	mp = make(map[int]int)
	return solve(0, 0, true) + 1
}

func bitsCount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	fmt.Println(maxPartitionsAfterOperations("abcba", 2))
}
