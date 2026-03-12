package art

import (
	"strings"
)

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
