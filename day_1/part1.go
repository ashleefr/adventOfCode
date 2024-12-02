package day_1

import (
	"math"
	"sort"
)

func CalculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		totalDistance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return totalDistance
}
