package day_3

import (
	"regexp"
)

func CalculateResultOfEnabledMultiplication(reports []string) int {
	resultOfEnabledMultiplication := 0
	enabled := true
	for _, line := range reports {
		regexpStr := regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
		allNeededStr := regexpStr.FindAllString(line, -1)

		for _, val := range allNeededStr {
			switch val {
			case "do()":
				enabled = true
			case "don't()":
				enabled = false
			default:
				if enabled {
					resultOfEnabledMultiplication += MultiplicationOfMul(val)
				}
			}

		}
	}
	return resultOfEnabledMultiplication
}
