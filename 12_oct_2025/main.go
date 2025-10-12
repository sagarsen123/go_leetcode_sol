package main

import "fmt"

const MOD int64 = 1_000_000_007

type key struct {
	i, remM, remK, carry int
}

func powmod(a, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = (res * a) % MOD
		}
		a = (a * a) % MOD
		b >>= 1
	}
	return res
}

func magicalSum(m int, k int, nums []int) int {
	n := len(nums)

	// Precompute combinations C(m, c)
	comb := make([][]int64, m+1)
	for i := 0; i <= m; i++ {
		comb[i] = make([]int64, m+1)
		comb[i][0] = 1
		for j := 1; j <= i; j++ {
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % MOD
		}
	}

	memo := make(map[key]int64)

	var popcount = func(x int) int {
		cnt := 0
		for x > 0 {
			cnt += x & 1
			x >>= 1
		}
		return cnt
	}

	var dp func(i, remM, remK, carry int) int64
	dp = func(i, remM, remK, carry int) int64 {
		if remM < 0 || remK < 0 {
			return 0
		}
		if i == n {
			if remM == 0 && remK == popcount(carry) {
				return 1
			}
			return 0
		}

		keyState := key{i, remM, remK, carry}
		if val, ok := memo[keyState]; ok {
			return val
		}

		var res int64 = 0
		val := int64(nums[i])

		for cnt := 0; cnt <= remM; cnt++ {
			c := comb[remM][cnt]
			powVal := powmod(val, int64(cnt))
			contribution := (c * powVal) % MOD

			newCarry := carry + cnt
			bit0 := newCarry & 1
			nextRemK := remK - bit0
			nextCarry := newCarry >> 1

			sub := dp(i+1, remM-cnt, nextRemK, nextCarry)
			if sub > 0 {
				res = (res + (contribution*sub)%MOD) % MOD
			}
		}

		memo[keyState] = res
		return res
	}

	return int(dp(0, m, k, 0))
}

func main() {
	nums := []int{5, 4, 3, 2, 1}
	m := 2
	k := 2
	fmt.Println(magicalSum(m, k, nums)) // example output
}
