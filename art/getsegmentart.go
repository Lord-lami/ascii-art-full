package art

import (
	"strings"
)

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
