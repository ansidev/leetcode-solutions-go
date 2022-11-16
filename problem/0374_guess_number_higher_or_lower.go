package problem

var pick = 0

func guess(n int) int {
	if n == pick {
		return 0
	}
	if n < pick {
		return 1
	}
	return -1
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *                1 if num is lower than the picked number
 *                otherwise return 0
 * func guess(num int) int;
 */
func guessNumber(n int) int {
	minNum, maxNum := 1, n
	for minNum < maxNum {
		x := minNum + (maxNum-minNum)/2
		check := guess(x)

		if check == 0 {
			return x
		} else if check == -1 {
			maxNum = x - 1
		} else {
			minNum = x + 1
		}
	}

	return minNum
}
