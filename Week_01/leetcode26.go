package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	left, right := 0, 1

	for right < len(nums) {

		for right < len(nums) && nums[left] == nums[right] {
			right++
		}

		if right == len(nums) {
			fmt.Println(nums)
			return left + 1
		}

		if left+1 < right {
			nums[left+1], nums[right] = nums[right], nums[left+1]
		}

		left++
		right++
	}

	fmt.Println(nums)
	return left + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
