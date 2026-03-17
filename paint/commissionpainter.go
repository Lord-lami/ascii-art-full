package paint

import (
	"asciiart/art"
	"regexp"
	"strings"
)

// CommissionPainter takes a string text, a substring subStr and a color string.
// It returns the art of the text with all matches of subStr painted in the
// specified color.
func CommissionPainter(text, subStr, color string) (paintTextArt func(start, stop int) string) {
	paintBrush := brush(color)
	reset := "\033[0m"

	drawStrArt := art.CommissionArtist(text)
	subStr = strings.ReplaceAll(subStr, "\\n", "\n")
	paintTextArt = func(start, stop int) string {
		var b strings.Builder
		if subStr == "" || text == subStr {
			b.WriteString(paintBrush)
			b.WriteString(drawStrArt(start, stop))
			b.WriteString(reset)
			return b.String()
		}

		subStrRe := regexp.MustCompile(regexp.QuoteMeta(subStr))
		positions := subStrRe.FindAllStringIndex(text, -1)
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
				continue
			}

			b.WriteString(drawStrArt(previousIndex, position[0]))

			if position[0] < position[1] {
				b.WriteString(paintBrush)
				b.WriteString(drawStrArt(position[0], position[1]))
				b.WriteString(reset)
			}

			previousIndex = position[1]
		}

		b.WriteString(drawStrArt(previousIndex, stop))
		return b.String()
	}
	return
}
