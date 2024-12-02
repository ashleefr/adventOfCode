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
	if isReportSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		modifiedReport := append([]int(nil), report...)
		modifiedReport = append(modifiedReport[:i], modifiedReport[i+1:]...)

		if isReportSafe(modifiedReport) {
			return true
		}
	}

	return false
}
