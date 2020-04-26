package main

import "sort"

func permuteUnique(nums []int) [][]int {

	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	sort.Ints(nums)
	res := [][]int(nil)
	for i := range nums {

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		leftNums := []int(nil)
		if i != 0 {
			leftNums = nums[:i]
		}

		rightNums := []int(nil)
		if i != len(nums)-1 {
			rightNums = nums[i+1:]
		}

		subNums := make([]int, len(nums)-1)
		copy(subNums, leftNums)
		copy(subNums[len(leftNums):], rightNums)

		subRes := permuteUnique(subNums)
		for _, sub := range subRes {
			res = append(res, append([]int{nums[i]}, sub...))
		}
	}

	return res
}
