package problem

const mod = 1e9 + 7

func numberOfWays1(startPos int, endPos int, k int) int {
	dp := [1001][1001]int{}
	for i := 1; i <= k; i++ {
		dp[i][i] = 1
		for j := 0; j < i; j++ {
			dp[i][j] = (dp[i-1][abs(j-1)] + dp[i-1][j+1]) % mod
		}
	}

	return dp[k][abs(startPos-endPos)]
}

func numberOfWays2(startPos int, endPos int, k int) int {
	pattern := make([]int, k+1)
	for i := range pattern {
		pattern[i] = -1
	}
	cache := make([][]int, startPos+k+1000)
	for j := range cache {
		cache[j] = make([]int, k+1)
		copy(cache[j], pattern)
	}

	return move(startPos, endPos, k, cache)
}

func move(start int, end int, k int, cache [][]int) int {
	if k == 0 && start == end {
		return 1
	}

	if k == 0 {
		return 0
	}

	if cache[start+999][k] != -1 {
		return cache[start+999][k]
	}

	a := move(start-1, end, k-1, cache)
	b := move(start+1, end, k-1, cache)

	cache[start+999][k] = (a + b) % (1e9 + 7)

	return cache[start+999][k]
}
