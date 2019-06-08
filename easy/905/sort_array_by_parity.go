package sortArrayByParity

func sortArrayByParity(A []int) []int {
	even := []int{}
	odd := []int{}
	for _, a := range A {
		if a%2 == 0 {
			even = append(even, a)
		} else {
			odd = append(odd, a)
		}
	}
	return append(even, odd...)
}
