package day_6

import (
	"bufio"
	"fmt"
	"os"
)

func Calculate() {
	var schema [][]byte
	ReadMapFromFile(&schema)
	fmt.Println("Day 6!")

	// Part 1
	totalVisitedPositions := CalculateSumOfVisitedPositions(schema)
	fmt.Printf("Total visited positions: %d\n", totalVisitedPositions)
	//// Part 2
	totalOfDifferentPositions := CalculateDifferentPositions(schema)
	fmt.Printf("Total of different positions: %d\n", totalOfDifferentPositions)
}

func ReadMapFromFile(schema *[][]byte) {
	file, _ := os.Open("day_6/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		*schema = append(*schema, []byte(line))
	}
}

var directions = []struct{ dx, dy int }{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}
