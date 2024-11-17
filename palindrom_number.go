func isPalindrome(x int) bool {
	return x == reverse(x)
}

func reverse(x int) int {
	result := 0

	for {
		if x <= 0 {
			return result
		}
		result = result*10 + (x % 10)
		x /= 10
	}
}