package problem

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{}
	}

	result := make([]int, l+1)

	carry := 1
	for i := l - 1; i >= 0; i-- {
		sum := digits[i] + carry
		carry = sum / 10
		result[i+1] = sum % 10
	}

	if carry > 0 {
		result[0] = 1
		return result
	}

	return result[1:]
}
