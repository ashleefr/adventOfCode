package day_2

func CalculateSafeReports(reports [][]int) int {
	totalScore := 0
	for i := 0; i < len(reports); i++ {
		if isReportSafe(reports[i]) {
			totalScore++
		}
	}
	return totalScore
}

func isReportSafe(report []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
