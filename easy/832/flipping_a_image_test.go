package flippingAImage

import (
	"fmt"
	"testing"
)

func TestFlippingAImage(t *testing.T) {
	// var a = [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	var a = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	b := flipAndInvertImage(a)
	fmt.Println(b)
}
