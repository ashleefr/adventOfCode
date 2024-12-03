package day_3

import (
	"bufio"
	"fmt"
	"os"
)

func Calculate() {
	var reports []string
	ReadLinesFromStdin(&reports)
	fmt.Println("Day 3!")

	// Part 1
	resultOfMultiplication := CalculateResultOfMultiplication(reports)
	fmt.Printf("Result of multiplication: %d\n", resultOfMultiplication)
	// Part 2
	resultOfEnabledMultiplication := CalculateResultOfEnabledMultiplication(reports)
	fmt.Printf("Result of enabled multiplication: %d\n\n", resultOfEnabledMultiplication)
}

func ReadLinesFromStdin(lines *[]string) {
	file, _ := os.Open("day_3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		//strLines := strings.Fields(line)
		*lines = append(*lines, line)
	}
}
