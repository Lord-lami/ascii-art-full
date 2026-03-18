package justify

import (
	"asciiart/measure"
	"asciiart/paint"
	"regexp"
	"strings"
)

// Right takes in text, subStr and color strings
// it returns the painted art of the text with
// the art justified to the right of the terminal.
func Right(text, subStr, color string) string {
	terminalW := measure.TerminalWidth()

	var b strings.Builder

	newlineRe := regexp.MustCompile(regexp.QuoteMeta("\n"))
	newlinePositions := newlineRe.FindAllStringIndex(text, -1)
	if len(newlinePositions) > 0 {
		newlinePositions = append(newlinePositions,
			[]int{newlinePositions[len(newlinePositions)-1][1], len(text)})
	} else {
		newlinePositions = [][]int{{0, len(text)}}
	}
	previousIndex := 0
	paintTextArt := paint.CommissionPainter(text, subStr, color)
	for _, position := range newlinePositions {
		line := text[previousIndex:position[1]]
		// The -1 is to stop the terminal from wraping the text when printing
		// reaches the end of the terminal. That is, characters should not be
		// printed at the exact end of the terminal.
		rightShiftSize := terminalW - measure.GetStringArtWidth(line) - 1
		coloredLineArt := paintTextArt(previousIndex, position[1])
		b.WriteString(rightShiftArtSegment(coloredLineArt, rightShiftSize))
		previousIndex = position[1]
	}
	return b.String()
}
