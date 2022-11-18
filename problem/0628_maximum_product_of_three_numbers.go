package problem

import "sort"

func maximumProduct(nums []int) int {
	l := len(nums)

	if l == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	sort.Ints(nums)

	return max(nums[0]*nums[1]*nums[l-1], nums[l-3]*nums[l-2]*nums[l-1])
}
