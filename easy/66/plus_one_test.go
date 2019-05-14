package plusone

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {

	digits := []int{
		9, 2, 9,
	}
	a := plusOne(digits)
	fmt.Println(a)
}
