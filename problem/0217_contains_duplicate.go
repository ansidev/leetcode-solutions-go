package problem

func containsDuplicate(nums []int) bool {
	freqMap := make(map[int]bool)

	for _, n := range nums {
		if freqMap[n] {
			return true
		}

		freqMap[n] = true
	}

	return false
}
