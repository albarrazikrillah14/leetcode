func longestCommonPrefix(strs []string) string {
	result := []rune{}

	if len(strs) == 1 {
		return strs[0]
	}

	for idx, value := range strs[0] {
		for j := range len(strs) - 1 {
			if idx > len(strs[j+1])-1 {
				return string(result)

			}

			if rune(strs[j+1][idx]) != value {
				return string(result)
			}
		}
		result = append(result, rune(value))
	}

	return string(result)
}