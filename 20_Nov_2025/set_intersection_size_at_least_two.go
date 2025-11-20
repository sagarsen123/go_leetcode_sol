package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	ans := 0
	first := -1
	second := -1
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	for i := range intervals {
		l := intervals[i][0]
		r := intervals[i][1]

		if l <= first {
			continue
		}

		if l > second {
			ans += 2
			first = r - 1
			second = r
		} else {
			ans += 1
			first = second
			second = r
		}
	}

	return ans
}

func main() {
	fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}))
}
