package justify

import (
	"asciiart/paint"
	"regexp"
	"strings"
)

func Center(text, subStr, color string) string {
	terminalW := terminalWidth()
	var b strings.Builder

	newlineRe := regexp.MustCompile(regexp.QuoteMeta("\n"))
	newlinePositions := newlineRe.FindAllStringIndex(text, -1)
	previousIndex := 0
	for _, position := range newlinePositions {
		line := text[previousIndex:position[1]]
		// The -1 is to stop the terminal from wraping the text when printing
		// reaches the end of the terminal. That is, characters should not be
		// printed at the exact end of the terminal.
		rightShiftSize := terminalW - getStringArtWidth(line)
		coloredLineArt := paint.PaintSubstring(text, subStr, color, previousIndex, position[1])
		if previousIndex == 0 {
			coloredLineArt = coloredLineArt[7:]
		}
		b.WriteString(rightShiftArtSegment(coloredLineArt, rightShiftSize/2))
		previousIndex = position[1]
	}
	rightShiftSize := terminalW - getStringArtWidth(text[previousIndex:]) - 1
	coloredLineArt := paint.PaintSubstring(text, subStr, color, previousIndex, len(text))
	b.WriteString(rightShiftArtSegment(coloredLineArt, rightShiftSize/2))
	return b.String()
}
