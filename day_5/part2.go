package day_5

import (
	"strconv"
	"strings"
)

func CalculateSumOfIncorrectMiddlePages(rules []string, orders [][]int) int {
	sumOfMiddlePages := 0
	for _, order := range orders {
		if !isValidOrder(rules, order) {
			sumOfMiddlePages += findMiddlePage(correctOrder(rules, order))
		}
	}
	return sumOfMiddlePages
}

func correctOrder(rules []string, order []int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, page := range order {
		inDegree[page] = 0
	}

	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])

		if _, okBefore := inDegree[before]; okBefore {
			if _, okAfter := inDegree[after]; okAfter {
				graph[before] = append(graph[before], after)
				inDegree[after]++
			}
		}
	}

	var sortedOrder []int
	var queue []int

	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sortedOrder
}
