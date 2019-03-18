package toLowerCase

import "testing"

func TestToLowerCase(t *testing.T) {
	str := "Hello Wolrd!"
	str = "here"
	str = "LOVELY"
	t.Log(toLowerCase(str))
}
