package day_4

import (
	"bufio"
	"fmt"
	"os"
)

func Calculate() {
	var grid [][]rune
	ReadFileToGrid(&grid)
	fmt.Println("Day 4!")

	// Part 1
	totalFoundXMAS := FindAllXMAS(grid)
	fmt.Printf("Total found XMAS: %d\n", totalFoundXMAS)
	// Part 2
	totalFoundMASInX := FindAllMASInX(grid)
	fmt.Printf("Total found MAS in X-form: %d\n\n", totalFoundMASInX)
}

func ReadFileToGrid(grid *[][]rune) {
	file, _ := os.Open("day_4/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		*grid = append(*grid, line)
	}
}
