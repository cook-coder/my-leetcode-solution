package removeOuterParentheses

import "testing"

func TestFunc(t *testing.T) {
	s := "(()())(())"
	// s := "(()())(())(()(()))"
	// s := "()()"
	r := removeOuterParentheses(s)
	t.Log(r)
}

func BenchmarkFunc(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := []string{"(()())(())", "(()())(())(()(()))", "()()"}
		removeOuterParentheses(s[i%3])
	}

	b.StopTimer()
}

func BenchmarkFuncV2(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := []string{"(()())(())", "(()())(())(()(()))", "()()"}
		removeOuterParenthesesV2(s[i%3])
	}

	b.StopTimer()
}
