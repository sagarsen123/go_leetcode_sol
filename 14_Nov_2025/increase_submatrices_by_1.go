package main

import "fmt"

func rangeAddQueries(n int, queries [][]int) [][]int {
	mat := make([][]int, n)
	for i := range n {
		mat[i] = make([]int, n)
	}

	for _, query := range queries {
		sr := query[0]
		sc := query[1]
		er := query[2]
		ec := query[3]

		for row := sr; row <= er; row++ {
			for col := sc; col <= ec; col++ {
				mat[row][col]++
			}
		}
	}
	return mat
}

func main() {
	fmt.Println(rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
}
