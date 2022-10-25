package problem

func maxSubArray(nums []int) int {
	currentMaxSum := nums[0]
	globalSum := nums[0]

	for _, x := range nums[1:] {
		if currentMaxSum+x > x {
			currentMaxSum += x
		} else {
			currentMaxSum = x
		}

		if globalSum < currentMaxSum {
			globalSum = currentMaxSum
		}
	}

	return globalSum
}
