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
	ReadTwoListsFromFile(&leftList, &rightList)
	fmt.Println("Day 1!")

	// Part 1
	totalDistance := CalculateTotalDistance(leftList, rightList)
	fmt.Printf("Total Distance: %d\n", totalDistance)
	// Part 2
	similarityScore := CalculateSimilarityScore(leftList, rightList)
	fmt.Printf("Total similarity score: %d\n\n", similarityScore)
}

func ReadTwoListsFromFile(leftList, rightList *[]int) {
	file, _ := os.Open("day_1/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		*leftList = append(*leftList, left)
		*rightList = append(*rightList, right)
	}
}
