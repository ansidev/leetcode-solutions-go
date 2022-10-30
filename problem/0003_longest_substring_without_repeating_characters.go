package problem

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 || l == 1 {
		return l
	}

	result, left := 1, 0
	lastIndex := map[byte]int{}
	lastIndex[s[0]] = 0
	for right := 1; right < l; right++ {
		if v, ok := lastIndex[s[right]]; ok {
			if left < v+1 {
				left = v + 1
			}
		}
		lastIndex[s[right]] = right
		if result < right-left+1 {
			result = right - left + 1
		}
	}

	return result
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	l := len(s)
	for left < l {
		// The bitSet corresponding to the character on the right is marked true,
		// indicating that this character is repeated at the X position and
		// needs to move forward on the left until X is marked as false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= l || right >= l {
			break
		}
	}
	return result
}
