package art

import (
	"strings"
)

// getTextArt returns the art of a text(multi-line or not)
// and continueLineArt, a function that can print additional 
// art on the last line of the art.
//
// The returned function should not be used with an empty string ("") under normal circumstances
// it will return the art of the last line excluding the first character.
func getTextArt(text string) (textArt string, continueLineArt func( addition string) string) {
	var b strings.Builder
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			start, rest := getLineArt(line)
			b.WriteString(start)
			b.WriteString(rest(""))
			continueLineArt = rest
		} else {
			b.WriteByte('\n')
			continueLineArt = nil
		}
	}
	textArt = b.String()
	return
}
