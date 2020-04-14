package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}

	if m == 0 {
		for idx, num := range nums2 {
			nums1[idx] = num
		}
		return
	}

	start1, start2 := 0, 0
	head := m

	for start1 < m && start2 < n {

		head = head % (m + n)

		if nums1[start1] < nums2[start2] {
			nums1[head] = nums1[start1]
			start1++
		} else {
			nums1[head] = nums2[start2]
			start2++
		}

		head++
	}

	if start1 == m {
		for start2 < n {

			head = head % (m + n)
			nums1[head] = nums2[start2]
			head++
			start2++
		}
	}

	reverse(nums1, 0, m-1)
	reverse(nums1, m, m+n-1)
	reverse(nums1, 0, m+n-1)
}

func reverse(nums []int, start int, end int) {

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
