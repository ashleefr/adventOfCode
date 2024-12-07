package day_6

func CalculateSumOfVisitedPositions(schema [][]byte) int {
	rows := len(schema)
	cols := len(schema[0])

	xPosition, yPosition, direction := findCurrentPosition(schema)

	visited := make(map[[2]int]bool)
	visited[[2]int{xPosition, yPosition}] = true

	for {
		xNext := xPosition + directions[direction].dx
		yNext := yPosition + directions[direction].dy

		if xNext < 0 || xNext >= rows || yNext < 0 || yNext >= cols {
			break
		}

		if schema[xNext][yNext] == '#' {
			direction = (direction + 1) % 4
		} else {
			xPosition, yPosition = xNext, yNext
			visited[[2]int{xPosition, yPosition}] = true
		}
	}

	return len(visited)
}

func findCurrentPosition(schema [][]byte) (int, int, int) {
	for i := 0; i < len(schema); i++ {
		for j := 0; j < len(schema[0]); j++ {
			switch schema[i][j] {
			case '^':
				return i, j, 0
			case '>':
				return i, j, 1
			case 'v':
				return i, j, 2
			case '<':
				return i, j, 3
			}
		}
	}
	return 0, 0, 0
}
