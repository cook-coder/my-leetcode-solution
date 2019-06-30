package countingbits

func countBits(num int) []int {
	result := make([]int, num+1)
	result[0] = 0
	base := 1 // 2 ** 0
	for i := 1; i <= num; i++ {
		if i < base {
			result[i] = result[base/2] + result[i-base/2]
		} else if i == base {
			result[i] = 1
			base *= 2
		}
	}
	return result
}

func countBitsVersion0(num int) []int {
	result := []int{0}
	base := 1
	for i := 1; i <= num; i++ {
		if i < base {
			result = append(result, result[base/2]+result[i-base/2])

		} else if i == base {
			result = append(result, 1)
			base *= 2
		}
	}
	return result
}
