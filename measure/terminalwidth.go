// Package measure provides tools for measuring the width
// of art and the terminal.
package measure

import (
	"strings"
)

// TerminalWidth returns the current width of the terminal.
func TerminalWidth(text string) (width int) {
	return getWidestLineWidth(text)
}

// getWidestLineWidth returns the width of the line with
// the widest art.
func getWidestLineWidth(text string) int {
	longestLineWidth := 0

	for line := range strings.SplitSeq(text, "\n") {
		lineLength := 0
		for _, char := range line {
			lineLength += GetCharArtWidth(byte(char))
		}
		longestLineWidth = max(lineLength, longestLineWidth)
	}
	return longestLineWidth + 1
}
