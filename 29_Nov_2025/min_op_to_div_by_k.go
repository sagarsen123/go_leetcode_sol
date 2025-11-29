package main

import "fmt"

func minOperations(nums []int, k int) int {
	    sum := 0
		    for _,v := range nums {
			        sum += v
					    }

						    return sum % k
							}
}


func main(){
	fmt.Println(minOperations([]int{3,9,7}, 5))
}