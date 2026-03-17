package paint

import (
	"asciiart/art"
	"regexp"
	"strings"
)

// PaintSubstring takes a string str, a substring subStr and a color string.
// It returns the art of the str with all matches of subStr painted in the
// specified color.
func PaintSubstring(str, subStr, color string, start, stop int) string {
	paintBrush := brush(color)
	reset := "\033[0m"

	var b strings.Builder
	drawStrArt := art.CommissionArtist(str)
	subStr = strings.ReplaceAll(subStr, "\\n", "\n")
	// drawStrArt := art.DrawTextArt(str, 0, len(str))

	if subStr == "" || str == subStr {
		b.WriteString(paintBrush)
		b.WriteString(drawStrArt(start, stop))
		b.WriteString(reset)
		return b.String()
	}

	subStrRe := regexp.MustCompile(regexp.QuoteMeta(subStr))
	positions := subStrRe.FindAllStringIndex(str, -1)
	previousIndex := start

	for _, position := range positions {
		if position[0] > stop {
			break
		}
		if position[0] < start {
			position[0] = start
		}
		if position[1] > stop {
			position[1] = stop
		}
		if position[1] < start {
			break
		}
		b.WriteString(drawStrArt(previousIndex, position[0]))

		b.WriteString(paintBrush)
		b.WriteString(drawStrArt(position[0], position[1]))
		b.WriteString(reset)

		previousIndex = position[1]
	}

	b.WriteString(drawStrArt(previousIndex, stop))
	return b.String()
}
