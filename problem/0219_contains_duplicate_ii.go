package problem

func containsNearbyDuplicate(nums []int, k int) bool {
	l := len(nums)
	m := make(map[int]int)
	for i := 0; i < l; i++ {
		if j, ok := m[nums[i]]; ok && i-j <= k {
			return true
		}

		m[nums[i]] = i
	}

	return false
}
