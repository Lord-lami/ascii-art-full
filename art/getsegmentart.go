package art

import (
	"strings"
)


// GetSegmentArt is the crown jewel of the art package.
// It takes a segment string (multi-line or not) which
// can be a slice of a string from any position and 
// a continueLineArt function which can be nil.
// It returns the segmentArt and the moreLineArt function
// which can print additional string art on the last line of segmentArt art.
//
// The returned function should not be used with an empty string ("") under normal circumstances
// it will return the art of the last line excluding the first character.
func GetSegmentArt(segment string,
	continueLineArt func(addition string) string) (
	segmentArt string,
	moreLineArt func(addition string) string) {
	if continueLineArt == nil && segment != "" {
		return getTextArt(segment)
	}
	var b strings.Builder
	remainder, otherLines, found := strings.Cut(segment, "\\n")
	if remainder != "" {
		b.WriteString(continueLineArt(remainder))
		moreLineArt = continueLineArt
	}
	if found {
		b.WriteString("\n")
		moreLineArt = nil
	}
	if otherLines != "" {
		var otherLinesArt string
		otherLinesArt, moreLineArt = getTextArt(otherLines)
		b.WriteString(otherLinesArt)
	}
	segmentArt = b.String()
	return
}
