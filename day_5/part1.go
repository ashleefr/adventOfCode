package day_5

import (
	"strconv"
	"strings"
)

func CalculateSumOfMiddlePages(rules []string, orders [][]int) int {
	sumOfMiddlePages := 0
	for _, order := range orders {
		if isValidOrder(rules, order) {
			sumOfMiddlePages += findMiddlePage(order)
		}
	}
	return sumOfMiddlePages
}

func isValidOrder(rules []string, order []int) bool {
	mapIndex := make(map[int]int)
	for i, page := range order {
		mapIndex[page] = i
	}

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])

		beforeIndex, beforeExists := mapIndex[before]
		afterIndex, afterExists := mapIndex[after]

		if !beforeExists || !afterExists {
			continue
		}

		if beforeIndex > afterIndex {
			return false
		}
	}

	return true
}

func findMiddlePage(order []int) int {
	return order[len(order)/2]
}
