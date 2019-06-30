package countingbits

import "testing"

func TestCounting(t *testing.T) {

	num := 16

	result := countBits(num)
	t.Log(result)
}
