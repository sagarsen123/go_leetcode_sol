package main

import "fmt"

func countOperations2(num1 int, num2 int) int {
	operations := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		operations++
	}
	return operations
}

func countOperations(num1 int, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return 0
	}

	if num1 < num2 {
		return countOperations(num2, num1)
	}
	return num1/num2 + countOperations(num2, num1%num2)
}

func main() {
	fmt.Println(countOperations(3, 2))
	fmt.Println(countOperations2(3, 2))
}
