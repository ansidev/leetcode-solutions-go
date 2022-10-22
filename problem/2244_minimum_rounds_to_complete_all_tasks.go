package problem

func minimumRounds(tasks []int) int {
	freqMap := make(map[int]int)
	minLevel := tasks[0]
	l := len(tasks)
	for i := 0; i < l; i++ {
		freqMap[tasks[i]]++
		if tasks[i] < minLevel {
			minLevel = tasks[i]
		}
	}

	round := 0
	for _, count := range freqMap {
		if count == 1 {
			return -1
		}

		for count%3 != 0 {
			count -= 2
			round++
		}
		round += count / 3
	}
	return round
}
