package day_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculate() {
	var reports [][]int
	ReadLinesFromStdinToSlice(&reports)
	fmt.Println("Day 2!")

	// Part 1
	totalSafeReports := CalculateSafeReports(reports)
	fmt.Printf("Total Distance: %d\n", totalSafeReports)
	// Part 2
	totalSafeReportsWithDampener := CalculateSafeReportsWithDampener(reports)
	fmt.Printf("Total Distance with Dampener: %d\n\n", totalSafeReportsWithDampener)
}

func ReadLinesFromStdinToSlice(lines *[][]int) {
	file, _ := os.Open("day_2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		strLines := strings.Fields(line)
		var levels []int
		for _, str := range strLines {
			level, _ := strconv.Atoi(str)
			levels = append(levels, level)
		}

		*lines = append(*lines, levels)
	}
}
