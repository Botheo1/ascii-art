package ascii

import (
	"strings"
)

// Transform the text into the ascii art version
func AsciiArt(texte string, banner string) string {
	chars := string(OpenFile(banner + ".txt"))
	lines := strings.Split(chars, "\n")
	newArg := strings.Split(texte, "\n")
	return PrintFormated(newArg, lines)
}
