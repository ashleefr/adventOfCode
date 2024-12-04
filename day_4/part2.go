package day_4

func FindAllMASInX(grid [][]rune) int {
	totalMASInX := 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 'A' {
				totalMASInX += IsItMAXInX(grid, i, j)
			}
		}
	}

	return totalMASInX
}

func IsItMAXInX(grid [][]rune, i, j int) int {
	if ((grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M')) &&
		((grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') || (grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M')) {
		return 1
	}
	return 0
}
