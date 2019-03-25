package toLowerCase

import (
	"strings"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	str := "Hello Wolrd!"
	str = "here"
	str = "LOVELY"
	str = "😁😆Hello！ État壹"
	t.Log(toLowerCase(str))
	t.Log(strings.ToLower(str))
}
