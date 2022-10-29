package problem

func maxPower(s string) int {
	l := len(s)

	if l == 0 || l == 1 {
		return l
	}

	currentCount := 0
	maxCount := 0
	var previousChar rune
	for _, currentChar := range s {
		if currentChar == previousChar {
			currentCount++
		} else {
			currentCount = 1
		}
		previousChar = currentChar
		if maxCount < currentCount {
			maxCount = currentCount
		}
	}
	return maxCount
}
