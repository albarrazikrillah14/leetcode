func getRow(rowIndex int) []int {
	result := make([][]int, rowIndex+1)

	for i := 0; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1

		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}

		result[i] = row
	}

	if rowIndex == 0 {
		return []int{1}
	}
	return result[rowIndex]
}