package day_4

func FindAllXMAS(grid [][]rune) int {
	totalXMAS := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 'X' {
				totalXMAS += CountXMASInGridCell(grid, i, j)
			}
		}
	}

	return totalXMAS
}

func CountXMASInGridCell(grid [][]rune, i, j int) int {
	counter := 0

	// up
	if CheckWordInDirection(grid, "XMAS", i, j, -1, 0) {
		counter++
	}
	// up-right
	if CheckWordInDirection(grid, "XMAS", i, j, -1, 1) {
		counter++
	}
	// right
	if CheckWordInDirection(grid, "XMAS", i, j, 0, 1) {
		counter++
	}
	// down-right
	if CheckWordInDirection(grid, "XMAS", i, j, 1, 1) {
		counter++
	}
	// down
	if CheckWordInDirection(grid, "XMAS", i, j, 1, 0) {
		counter++
	}
	// down-left
	if CheckWordInDirection(grid, "XMAS", i, j, 1, -1) {
		counter++
	}
	// left
	if CheckWordInDirection(grid, "XMAS", i, j, 0, -1) {
		counter++
	}
	// up-left
	if CheckWordInDirection(grid, "XMAS", i, j, -1, -1) {
		counter++
	}

	return counter
}

func CheckWordInDirection(grid [][]rune, word string, row, col, dirRow, dirCol int) bool {
	wordLen := len(word)
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < wordLen; i++ {
		newRow := row + i*dirRow
		newCol := col + i*dirCol
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}
		if grid[newRow][newCol] != rune(word[i]) {
			return false
		}
	}
	return true
}
