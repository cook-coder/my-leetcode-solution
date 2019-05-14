package plusone

func plusOne(digits []int) []int {
	plusOneIndex := len(digits) - 1
	for {
		digits[plusOneIndex]++
		if digits[plusOneIndex] > 9 {
			digits[plusOneIndex] = 0
			plusOneIndex--
			if plusOneIndex < 0 {
				return append([]int{1}, digits...)
			}
		} else {
			break
		}

	}
	return digits
}
