package day_5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculate() {
	var rules []string
	var orders [][]int
	ReadFileToTwoSlices(&rules, &orders)
	fmt.Println("Day 5!")

	//// Part 1
	totalSumOfMiddlePages := CalculateSumOfMiddlePages(rules, orders)
	fmt.Printf("Total sum of middle pages: %d\n", totalSumOfMiddlePages)
	//// Part 2
	CalculateSumOfIncorrectMiddlePages := CalculateSumOfIncorrectMiddlePages(rules, orders)
	fmt.Printf("Total sorted sum of middle pages: %d\n\n", CalculateSumOfIncorrectMiddlePages)
}

func ReadFileToTwoSlices(rules *[]string, orders *[][]int) {
	file, _ := os.Open("day_5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	readingRules := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingRules = false
			continue
		}

		if readingRules {
			*rules = append(*rules, line)
		} else {
			var nums []int
			parts := strings.Split(line, ",")
			for _, part := range parts {
				num, _ := strconv.Atoi(part)
				nums = append(nums, num)
			}
			*orders = append(*orders, nums)
		}
	}
}
