package main

func moveZeroes(nums []int)  {

    if len(nums) <= 1 {
        return
    }

    left := 0
    for left < len(nums) && nums[left] != 0 {
        left++
    }

    if left == len(nums) {
        return
    }

    right := left + 1

    for right < len(nums) {

        for right < len(nums) && nums[right] == 0 {
            right++
        }

        if right == len(nums) {
            return
        }

        nums[left], nums[right] = nums[right], nums[left]
        left++
        right++
    }
}

