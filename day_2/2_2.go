package day_2

func CalculateSafeReportsWithDampener(reports [][]int) int {
	totalScore := 0
	for i := 0; i < len(reports); i++ {
		if isReportSafeWithDampener(reports[i]) {
			totalScore++
		}
	}
	return totalScore
}

func isReportSafeWithDampener(report []int) bool {
	isIncreasing := true
	isDecreasing := true
	flag := true
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff < -3 || diff > 3 || diff == 0 {
			if flag {
				flag = false
				continue
			}
			return false
		}

		if diff > 0 {
			if flag {
				flag = false
				continue
			}
			isDecreasing = false
		} else if diff < 0 {
			if flag {
				flag = false
				continue
			}
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
