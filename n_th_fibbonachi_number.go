func tribonacci(n int) int {
	if n <= 2 {
		if n == 0 {
			return 0
		}

		return 1
	}

	result := make([]int, n+1)
	result[0] = 0
	result[1] = 1
	result[2] = 1

	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2] + result[i-3]
	}

	return result[n]
}