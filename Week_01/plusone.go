package main

func plusOne(digits []int) []int {
	res := make([]int, len(digits)+1)
	start := len(digits) - 1
	cnt := 1
	overflow := false
	for ; start >= 0; start-- {

		sum := digits[start] + cnt
		if sum < 10 {
			res[start+1] = sum
			cnt = 0
			continue
		}

		res[start+1] = 0
		overflow = start == 0
	}

	if overflow {
		res[0] = 1
		return res
	}

	return res[1:]
}
