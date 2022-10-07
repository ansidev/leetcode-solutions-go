package problem

func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		x := target - n

		if j, ok := m[x]; ok {
			return []int{i, j}
		}

		m[n] = i
	}

	return []int{}
}
