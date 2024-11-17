func romanToInt(s string) int {
	result := 0

	for idx, value := range s {
		if idx > 0 {
			if isBigger(rune(value), rune(s[idx-1])) {
				result -= 2 * baseValue(rune(s[idx-1]))
			}
		}

		result += baseValue(value)
	}

	return result
}

func isBigger(a rune, b rune) bool {
	return baseValue(a) > baseValue(b)
}

func baseValue(s rune) int {
	switch s {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}