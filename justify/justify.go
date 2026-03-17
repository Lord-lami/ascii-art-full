package justify

import (
	"asciiart/paint"
	"regexp"
	"strings"
)

func Justify(text, subStr, color string) string {
	terminalW := TerminalWidth()
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
	for _, nlPosition := range newlinePositions {
		line := text[previousIndex:nlPosition[1]]
		shiftableRe := regexp.MustCompile(`\s+[\S\n]*`)
		shiftablePositions := shiftableRe.FindAllStringIndex(line, -1)
		numberOfShiftables := len(shiftablePositions)
		if numberOfShiftables == 0 {
			panic("left align not implemented")
		}
		spaceSize := (terminalW - GetStringArtWidth(line))
		rightShiftSize := spaceSize / numberOfShiftables
		coloredLineArt := paintTextArt(previousIndex, previousIndex+shiftablePositions[0][0])
		b.WriteString(coloredLineArt)
		totalRightShifts := 0
		for _, shiftablePosition := range shiftablePositions {
			coloredLineArt := paintTextArt(previousIndex+shiftablePosition[0],
				previousIndex+shiftablePosition[1])
			// To precisely avoid end of terminal wrapping
			totalRightShifts += rightShiftSize
			if totalRightShifts >= spaceSize {
				rightShiftSize -= totalRightShifts - spaceSize + 1
			}
			b.WriteString(rightShiftArtSegment(coloredLineArt, rightShiftSize))
		}
		previousIndex = nlPosition[1]
	}
	return b.String()
}
