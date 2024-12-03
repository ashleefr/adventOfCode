package day_3

import (
	"regexp"
	"strconv"
)

func CalculateResultOfMultiplication(reports []string) int {
	resultOfMultiplication := 0

	for _, line := range reports {
		regexpStr := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
		allNeededStr := regexpStr.FindAllString(line, -1)

		for _, val := range allNeededStr {
			resultOfMultiplication += MultiplicationOfMul(val)
		}
	}
	return resultOfMultiplication
}

func MultiplicationOfMul(mul string) int {
	regexpStr := regexp.MustCompile(`mul\((\d+),\s*(\d+)\)`)

	matches := regexpStr.FindStringSubmatch(mul)

	firstNumber, _ := strconv.Atoi(matches[1])
	secondNumber, _ := strconv.Atoi(matches[2])
	return firstNumber * secondNumber
}
