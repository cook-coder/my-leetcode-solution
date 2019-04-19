package removeOuterParentheses

import "bytes"

/*
ascii:
	40 => (
	41 => )
*/
func removeOuterParentheses(S string) string {
	var leftParentheses []rune
	var partialParentheses []rune
	var r string
	for _, s := range S {
		partialParentheses = append(partialParentheses, s)
		if s == 40 {
			leftParentheses = append(leftParentheses, s)
		} else {
			leftParentheses = leftParentheses[:len(leftParentheses)-1]
			if len(leftParentheses) == 0 {
				r += string(partialParentheses[1 : len(partialParentheses)-1])
				partialParentheses = partialParentheses[0:0]
			}
		}
	}
	return r
}

func removeOuterParenthesesV2(S string) string {
	ans := new(bytes.Buffer)
	l, r := 0, 0
	for i := range S {

		if S[i] == '(' {
			l++
		} else {
			r++
			if l == r {
				ans.WriteString(S[i-l-r+2 : i])
				l = 0
				r = 0
			}
		}
	}

	return ans.String()
}
