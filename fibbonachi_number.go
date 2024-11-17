func fib(n int) int {
	f := make([]int, n+1)

	if n <= 1 {
		return n
	}

	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n]
}