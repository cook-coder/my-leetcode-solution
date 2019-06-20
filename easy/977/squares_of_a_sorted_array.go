package squaresofasortedarray

func sortedSquares(A []int) []int {
	ALength := len(A)
	B := make([]int, ALength, ALength)
	l, r := 0, ALength-1
	i := r
	for l <= r {
		if A[l]+A[r] < 0 {
			B[i] = A[l] * A[l]
			l++
		} else {
			B[i] = A[r] * A[r]
			r--
		}
		i--
	}

	return B
}
