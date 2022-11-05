package problem

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1

	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums[n]
}
