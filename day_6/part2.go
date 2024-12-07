package day_6

func CalculateDifferentPositions(schema [][]byte) int {
	rows := len(schema)
	cols := len(schema[0])

	xPosition, yPosition, direction := findCurrentPosition(schema)

	validPositions := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if schema[i][j] == '^' || schema[i][j] == '>' || schema[i][j] == 'v' || schema[i][j] == '<' {
				continue
			}

			updatedSchema := make([][]byte, rows)
			for k := 0; k < rows; k++ {
				updatedSchema[k] = make([]byte, cols)
				copy(updatedSchema[k], schema[k])
			}

			updatedSchema[i][j] = '#'

			if checkForCycle(updatedSchema, xPosition, yPosition, direction) {
				validPositions++
			}
		}
	}

	return validPositions
}

func checkForCycle(schema [][]byte, xPosition, yPosition, direction int) bool {
	rows := len(schema)
	cols := len(schema[0])

	visited := make(map[[2]int]int)

	for {
		xNext := xPosition + directions[direction].dx
		yNext := yPosition + directions[direction].dy

		if xNext < 0 || xNext >= rows || yNext < 0 || yNext >= cols {
			return false
		}

		if schema[xNext][yNext] == '#' {
			direction = (direction + 1) % 4
		} else {
			xPosition, yPosition = xNext, yNext
			if visited[[2]int{xPosition, yPosition}] == 5 {
				return true
			}
			visited[[2]int{xPosition, yPosition}]++
		}
	}
}
