package problem

func firstMissingPositive(nums []int) int {
	n := len(nums)

	i := 0
	for i < n {
		if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for j := 0; j < n; j++ {
		if nums[j] != j+1 {
			return j + 1
		}
	}
	return n + 1
}
