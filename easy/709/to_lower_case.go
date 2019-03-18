package toLowerCase

func toLowerCase(str string) string {
	strRunes := []rune(str)
	for i, c := range strRunes {
		if c > 64 && c < 91 {
			strRunes[i] += 32
		}
	}
	return string(strRunes)
}
