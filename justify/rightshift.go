package justify

import (
	"strings"
)

func rightShiftArtSegment(artSegment string, shiftSize int) string {
	var b strings.Builder
	paddingLine := strings.Repeat("-", shiftSize)
	// fmt.Printf("%q\n\n", artSegment)
	if len(artSegment) > 7 && artSegment[:7] == "\n\n\n\n\n\n\n" {
		// log.Fatal("Working!")
		b.WriteString(artSegment[:7])
		artSegment = artSegment[7:]
	}
	b.WriteString(paddingLine)
	b.WriteString(artSegment)
	return b.String()
}
