package paint

import (
	"asciiart/art"
	"regexp"
	"strings"
)

func PaintSubstring(str, subStr, color string) string {
	paintBrush := brush(color)
	reset := "\033[0m"

	if paintBrush == "" {
		return ""
	}

	var b strings.Builder

	if subStr == "" || str == subStr {
		b.WriteString(paintBrush)
		textArt, _ := art.GetSegmentArt(str, nil)
		b.WriteString(textArt)
		b.WriteString(reset)
		return b.String()
	}

	subStrRe := regexp.MustCompile(regexp.QuoteMeta(subStr))
	positions := subStrRe.FindAllStringIndex(str, -1)
	previousIndex := 0
	segment := ""
	segmentArt := ""
	var continuation func(string) string

	if len(positions) > 0 && positions[0][0] == 0 {
		segment = str[positions[0][0]:positions[0][1]]
		segmentArt, continuation = art.GetSegmentArt(segment, continuation)
		b.WriteString(paintBrush)
		b.WriteString(segmentArt)
		b.WriteString(reset)
		previousIndex = positions[0][1]
		positions = positions[1:]
	}
	if len(positions) > 0 && positions[0][0] != 0 {
		segment = str[previousIndex:positions[0][0]]
		segmentArt, continuation = art.GetSegmentArt(segment, continuation)
		b.WriteString(segmentArt)
		previousIndex = positions[0][0]
	}

	for i, position := range positions {
		segment = str[previousIndex:position[1]]
		segmentArt, continuation = art.GetSegmentArt(segment, continuation)
		b.WriteString(paintBrush)
		b.WriteString(segmentArt)
		b.WriteString(reset)
		previousIndex = position[1]

		if i+1 < len(positions) {
			segment = str[previousIndex:positions[i+1][0]]
			segmentArt, continuation = art.GetSegmentArt(segment, continuation)
			b.WriteString(segmentArt)
			previousIndex = positions[i+1][0]
		}
	}
	segment = str[previousIndex:]
	segmentArt, _ = art.GetSegmentArt(segment, continuation)
	b.WriteString(segmentArt)
	return b.String()
}
