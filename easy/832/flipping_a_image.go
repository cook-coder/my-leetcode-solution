package flippingAImage

func flipAndInvertImage(A [][]int) [][]int {
	var B [][]int
	for i := 0; i < len(A); i++ {
		var b []int
		for j := len(A[i]) - 1; j > -1; j-- {
			b = append(b, 1-A[i][j])
		}
		B = append(B, b)
	}
	return B
}

func flipAndInvertImage2(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		length := len(A[i])
		for j := 0; j <= (length-1)/2; j++ {
			A[i][j], A[i][length-j-1] = A[i][length-j-1]^1, A[i][j]^1
		}
	}
	return A
}
