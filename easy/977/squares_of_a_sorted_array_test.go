package squaresofasortedarray

import "testing"

func TestFunc(t *testing.T) {
	A := []int{-4, -1, 0, 3, 10}
	B := sortedSquares(A)
	t.Log(B)
}
