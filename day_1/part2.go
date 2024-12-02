package day_1

func CalculateSimilarityScore(leftList, rightList []int) int {
	totalScore := 0

	for i := 0; i < len(leftList); i++ {
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				totalScore += leftList[i]
			}
		}
	}

	return totalScore
}
