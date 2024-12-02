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

	// Day 2 Task 1
	totalSafeReports := CalculateSafeReports(reports)
	fmt.Printf("Total Distance: %d\n", totalSafeReports)
	// Day 2 Task 2
	totalSafeReportsWithDampener := CalculateSafeReportsWithDampener(reports)
	fmt.Printf("Total Distance: %d\n", totalSafeReportsWithDampener)
}

func ReadLinesFromStdinToSlice(lines *[][]int) {
	fmt.Println("Day 2. Input:")
	scanner := bufio.NewScanner(os.Stdin)
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
