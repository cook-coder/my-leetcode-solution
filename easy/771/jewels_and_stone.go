package jewelsAndStone

// 二逼解法
func BadSolution(J string, S string) int {
	var jCount int
	for _, j := range J {
		for _, s := range S {
			if j == s {
				jCount++
			}
		}
	}
	return jCount
}

// 改进过的解法
func ImprovedSolution(J string, S string) int {
	jRunes := []rune(J)
	sRunes := []rune(S)
	sMap := make(map[rune]int)
	for _, s := range sRunes {
		sMap[s]++
	}
	var jCount int
	for _, j := range jRunes {
		jCount += sMap[j]
	}
	return jCount
}
