package main

func rotate(nums []int, k int) {

	if len(nums) <= 1 {
		return
	}

	k = k % len(nums)

	if k == 0 {
		return
	}

	right := len(nums) - k
	left := right - 1

	reverse(nums, 0, left)
	reverse(nums, right, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start int, end int) {

	if start >= end {
		return
	}

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
