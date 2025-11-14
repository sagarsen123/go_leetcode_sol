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

// DIFFERENCE ARRAY APPROACH
func rangeAddQueries2(n int, queries [][]int) [][]int {
	mat := make([][]int, n)

	for i := range n {
		mat[i] = make([]int, n)
	}

	for _, query := range queries {
		r1 := query[0]
		c1 := query[1]
		r2 := query[2]
		c2 := query[3]

		for row := r1; row <= r2; row++ {
			mat[row][c1]++
			if c2+1 < n {
				mat[row][c2+1]--
			}
		}

	}
	for row := range n {
		for col := 1; col < n; col++ {
			mat[row][col] += mat[row][col-1]
		}
	}
	return mat
}

func main() {
	fmt.Println(rangeAddQueries(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
	fmt.Println(rangeAddQueries2(3, [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}}))
}
