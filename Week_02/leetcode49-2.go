package main

func nthUglyNumber(n int) int {

	if n <= 0 {
		return -1
	}

	dp := []int{1}
	idx2, idx3, idx5 := 0, 0, 0
	for n > 1 {

		x := dp[idx2] * 2
		y := dp[idx3] * 3
		z := dp[idx5] * 5

		m := min(min(x, y), z)
		if m == x {
			idx2++
		}

		if m == y {
			idx3++
		}

		if m == z {
			idx5++
		}

		dp = append(dp, m)
		n--
	}

	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
