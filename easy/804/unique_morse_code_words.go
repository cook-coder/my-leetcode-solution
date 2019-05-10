package uniqueMorseCodeWords

func uniqueMorseRepresentations(words []string) int {
	var words2Morse = map[rune]string{
		97:  ".-",
		98:  "-...",
		99:  "-.-.",
		100: "-..",
		101: ".",
		102: "..-.",
		103: "--.",
		104: "....",
		105: "..",
		106: ".---",
		107: "-.-",
		108: ".-..",
		109: "--",
		110: "-.",
		111: "---",
		112: ".--.",
		113: "--.-",
		114: ".-.",
		115: "...",
		116: "-",
		117: "..-",
		118: "...-",
		119: ".--",
		120: "-..-",
		121: "-.--",
		122: "--..",
	}
	uniqueMorseWords := make(map[string]byte)
	var eachWord string
	for _, word := range words {
		for _, char := range word {
			eachWord += words2Morse[char]
		}
		uniqueMorseWords[eachWord] = 0
		eachWord = ""
	}
	return len(uniqueMorseWords)
}
