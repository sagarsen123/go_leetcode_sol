package main

import "fmt"

// method 1
func hasSameDigits(s string) bool {
	a := []byte(s)
	n := len(a)
	for n > 2 {
		for i := 0; i < n-1; i++ {
			a[i] = (a[i] + a[i+1]) % 10
		}
		n--
	}
	return a[0] == a[1]
}

// method 2
// func hasSameDigits(s string) bool {
//     n := len(s)
//     intarr := make([]int, n)
//     for i := range s{
//         intarr[i] = int(s[i]-'0')
//     }

//     for n > 2{
//         for i:=0;i+1<n;i++{
//             intarr[i] = ( intarr[i] + intarr[i+1] ) % 10
//         }
//         n--
//     }
//     return intarr[0]==intarr[1]
// }

func main() {
	fmt.Println(hasSameDigits("3902"))
	fmt.Println(hasSameDigits("34789"))
}
