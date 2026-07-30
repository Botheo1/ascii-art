package ascii

// Print the ascii art representation of the text
func PrintFormated(words []string, lines []string) string {
	var result string
	for _, word := range words {
		if word == "" {
			result += "\n"
			continue
		}
		for index := 1; index <= 8; index++ {
			for _, letter := range word {
				position := int(letter - 32)
				currentLine := (position * 9) + index
				result += (lines[currentLine])
			}
			result += "\n"
		}
	}
	return result
}
