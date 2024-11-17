func longestPalindrome(s string) int {
	hashMap := map[rune]int{}
	result := 0

	for _, char := range s {
		hashMap[char]++

		if hashMap[char]%2 == 0 {
			result += 2
		}
	}

	if result < len(s) {
		result++
	}
	return result
}