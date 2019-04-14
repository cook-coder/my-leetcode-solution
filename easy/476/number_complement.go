package numberComplement

import (
	"strconv"
)

func findComplement(num int) int {
	binaryStr := strconv.FormatInt(int64(num), 2)
	var r []rune
	for _, b := range binaryStr {
		if b == 48 {
			b = 49
		} else {
			b = 48
		}
		r = append(r, b)
	}
	result, _ := strconv.ParseInt(string(r), 2, 0)
	return int(result)
}
