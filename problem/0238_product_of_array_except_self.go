package problem

func productExceptSelf(nums []int) []int {
	// calculate left product and right product of each element
	// then multiply each element of two arrays has same index
	// finally, we have the result array
	l := len(nums)
	product := 1
	result := make([]int, l)

	result[0] = 1
	for i := 1; i < l; i++ {
		product *= nums[i-1]
		result[i] = product
	}

	product = 1
	for i := l - 2; i >= 0; i-- {
		product *= nums[i+1]
		result[i] *= product
	}

	return result
}
