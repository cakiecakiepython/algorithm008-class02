package main

func combine(n int, k int) [][]int {

	if n < k {
		return nil
	}

	return combineHelper(1, n, k)
}

func combineHelper(left int, right int, k int) [][]int {

	if k == 0 {
		return nil
	}

	res := [][]int(nil)

	for start := left; start+k-1 <= right; start++ {

		subRes := combineHelper(start+1, right, k-1)
		if len(subRes) == 0 {
			res = append(res, []int{start})
			continue
		}

		for _, subNums := range subRes {
			res = append(res, append([]int{start}, subNums...))
		}
	}

	return res
}
