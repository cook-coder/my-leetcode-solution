package toLowerCase

import "fmt"

func toLowerCase(str string) string {
	strRunes := []rune(str)
	for i, c := range strRunes {
		fmt.Println(i, c)
		if c > 64 && c < 91 {
			strRunes[i] += 32
		}
	}
	return string(strRunes)
}
