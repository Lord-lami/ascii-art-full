package oversabi

import (
	"asciiart/measure"
	"strings"
)

func Wrap(text string) string {
	terminlWidth := measure.TerminalWidth()
	var wrappedTextBuilder strings.Builder
	var lineBuilder strings.Builder
	bar := 0
	textLines := strings.SplitAfter(text, "\n")
	var WrapAroundSep func([]string, string)
	WrapAroundSep = func(lines []string, sep string) {
		textWords := []string{}
		for _, line := range lines {
			textWords = append(textWords, strings.SplitAfter(line, sep)...)
		}

		for i := 0; i < len(textWords); i++ {
			word := textWords[i]
			wordLength := measure.GetStringArtWidth(word)
			if bar+wordLength < terminlWidth {
				bar += wordLength
				wrappedTextBuilder.WriteString(word)
				if word[len(word)-1] == '\n' {
					bar = 0
					continue
				}
			} else {
				if wordLength < terminlWidth {
					wrappedTextBuilder.WriteByte('\n')
					bar = 0
					i--
				} else {
					WrapAroundSep([]string{word}, "")
				}
			}
		}
	}
	WrapAroundSep(textLines, " ")
	wrappedTextBuilder.WriteString(lineBuilder.String())
	return wrappedTextBuilder.String()
}
