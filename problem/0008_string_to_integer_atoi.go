package problem

import "math"

func myAtoi(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	i := 0
	for i < l && s[i] == ' ' {
		i++
	}

	if i < l && s[i] != '+' &&
		s[i] != '-' &&
		s[i] != '0' &&
		s[i] != '1' &&
		s[i] != '2' &&
		s[i] != '3' &&
		s[i] != '4' &&
		s[i] != '5' &&
		s[i] != '6' &&
		s[i] != '7' &&
		s[i] != '8' &&
		s[i] != '9' {
		return 0
	}

	isNegative := false
	if i < l && s[i] == '-' {
		isNegative = true
		i++
	} else if i < l && s[i] == '+' {
		i++
	}

	if i < l && (s[i] == '+' || s[i] == '-') {
		return 0
	}

	var n float64 = 0
	for i < l && s[i] >= '0' && s[i] <= '9' {
		n = n*10 + float64(s[i]-'0')
		i++
	}

	if isNegative {
		n = -n
	}

	if n < math.MinInt32 {
		return math.MinInt32
	}
	if n > math.MaxInt32 {
		return math.MaxInt32
	}

	return int(n)
}
