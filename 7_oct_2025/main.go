package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	lakeRainDay := make(map[int]int)
	var zeroDays []int

	for i := range n {
		lake := rains[i]

		if lake == 0 {
			zeroDays = append(zeroDays, i)
			ans[i] = 1
		} else {
			day, ok := lakeRainDay[lake]
			if ok {
				idx := sort.Search(len(zeroDays), func(j int) bool {
					return zeroDays[j] > day
				})

				if idx == len(zeroDays) {
					return []int{}
				}

				zeroDay := zeroDays[idx]
				ans[zeroDay] = lake

				zeroDays = append(zeroDays[:idx], zeroDays[idx+1:]...)
			}
			lakeRainDay[lake] = i
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(avoidFlood([]int{1, 2, 0, 0, 2, 1}))
	// Output: [-1 -1 2 1 -1 -1]
}
