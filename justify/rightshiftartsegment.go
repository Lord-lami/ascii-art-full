// Package justify provides tools for justifying/arranging/aligning
// the art of a text on the terminal.
package justify

import (
	"strings"
)

// rightShiftArtSegment is the core of the justify package.
// It shifts an art segment (artSegment) to the right
// shiftSize amount of times.
func rightShiftArtSegment(artSegment string, shiftSize int) string {
	var b strings.Builder
	// fmt.Println(len("\033[7A"))
	paddingLine := strings.Repeat(" ", shiftSize)
	start, _, _ := strings.Cut(artSegment, "\033[7A")
	if len(start) >= 7 {
		b.WriteString(start)
		remaining, works := strings.CutPrefix(artSegment, start)
		if !works {
			panic("something is horribly wrong")
		}
		artSegment = remaining
	}
	b.WriteString(paddingLine)
	b.WriteString(artSegment)
	// fmt.Printf("%q\n", b.String())
	return b.String()
}
