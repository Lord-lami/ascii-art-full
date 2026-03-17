package justify

import (
	"strings"
)

func rightShiftArtSegment(artSegment string, shiftSize int) string {
	var b strings.Builder
	paddingLine := strings.Repeat(" ", shiftSize)
	if artSegment != "\n" {
		b.Write([]byte{'\n', '\n', '\n', '\n', '\n', '\n', '\n'})
	}
	
	b.WriteString(paddingLine)
	b.WriteString(artSegment)
	return b.String()
}
