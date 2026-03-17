package justify

import (
	"asciiart/paint"
	"fmt"
	"regexp"
	"strings"
)

func Center(text, subStr, color string) string {
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
	fmt.Println("Here")
	previousIndex := 0
	paintTextArt := paint.CommissionPainter(text, subStr, color)
	for _, position := range newlinePositions {
		
		line := text[previousIndex:position[1]]
		rightShiftSize := (terminalW - GetStringArtWidth(line)) / 2
		coloredLineArt := paintTextArt(previousIndex, position[1])
		b.WriteString(rightShiftArtSegment(coloredLineArt, rightShiftSize))
		previousIndex = position[1]
	}
	return b.String()
}
