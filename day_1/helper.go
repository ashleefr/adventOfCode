package day_1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculate() {
	var leftList, rightList []int
	ReadTwoListFromStdin(&leftList, &rightList)

	// Day 1 Task 1
	totalDistance := CalculateTotalDistance(leftList, rightList)
	fmt.Printf("Total Distance: %d\n", totalDistance)
	// Day 1 Task 2
	similarityScore := CalculateSimilarityScore(leftList, rightList)
	fmt.Printf("Total similarity score: %d\n", similarityScore)
}

func ReadTwoListFromStdin(leftList, rightList *[]int) {
	fmt.Println("Day 1. Input:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Fields(line)
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		*leftList = append(*leftList, left)
		*rightList = append(*rightList, right)
	}
}
