package uniqueMorseCodeWords

import (
	"fmt"
	"testing"
)

func TestFun(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	num := uniqueMorseRepresentations(words)
	fmt.Println(num)
}
