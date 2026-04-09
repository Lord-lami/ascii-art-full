package art

import (
	"asciiart/measure"
	"regexp"
	"strconv"
	"strings"
)

// DrawPaintWrapAlignTextArt takes a text string and returns
// the art of a text from start index to stop index (half open range).
//
// Note: the strings that are passed will have all `\n`s`
// replaced by newline characters.
func DrawPaintWrapAlignTextArt(text, subStr, brush, alignment string) string {
	terminalWidth := measure.TerminalWidth()
	if subStr == "" {
		subStr = text
	}
	text = strings.ReplaceAll(text, "\\n", "\n")
	subStr = strings.ReplaceAll(subStr, "\\n", "\n")
	subStrRe := regexp.MustCompile(regexp.QuoteMeta(subStr))
	paintPositions := subStrRe.FindAllStringIndex(text, -1)
	lineArtWidth := 0
	const (
		moveUp7  = "\033[7A"
		moveDown = "\033[B"
	)
	var b strings.Builder
	wordSpace := ""
	totalLineSpaces := 0
	lineSpacesSeen := 0

	for i, char := range text {
		charArt, charArtWidth := drawCharArt(byte(char))
		charArtStr := string(charArt)

		if lineArtWidth+charArtWidth >= terminalWidth || charArtStr == "\n" {
			lineArtWidth = 0
		}

		if charArtStr != "\n" {
			if lineArtWidth == 0 {
				b.Write([]byte{'\n', '\n', '\n', '\n', '\n', '\n', '\n'})
			}
			b.WriteString(moveUp7)
			moveBackCAL := "\033[" + strconv.Itoa(charArtWidth) + "D"
			charArtStr = strings.Replace(charArtStr, "\n", moveDown+moveBackCAL, 7)
		}

		// Aligning
		if lineArtWidth == 0 {
			nextLineArtWidth := 0
			j := 0
			for ; i+j < len(text) && text[i+j] != '\n'; j++ {
				nextCharsWidth := measure.GetCharArtWidth(text[i+j])
				if nextLineArtWidth+nextCharsWidth >= terminalWidth {
					break
				}
				nextLineArtWidth += nextCharsWidth
			}
			switch alignment {
			case "left":
			case "right":
				shiftSize := terminalWidth - nextLineArtWidth - 1
				lineSpace := strings.Repeat(" ", shiftSize)
				b.WriteString(lineSpace)
			case "center":
				shiftSize := (terminalWidth - nextLineArtWidth) / 2
				lineSpace := strings.Repeat(" ", shiftSize)
				b.WriteString(lineSpace)
			case "justify":
				totalLineSpaces = max(strings.Count(text[i:i+j], " "), 1)
				lineSpacesSeen = 0
				shiftSize := (terminalWidth - nextLineArtWidth) / totalLineSpaces
				wordSpace = strings.Repeat(" ", shiftSize)
			default:
				panic("alignment must be one of: left, right, center and justify")
			}

		}

		// Painting
		if len(paintPositions) > 0 && i == paintPositions[0][1] {
			b.WriteString("\033[0m")
			paintPositions = paintPositions[1:]
		}
		if len(paintPositions) > 0 && i == paintPositions[0][0] {
			b.WriteString(brush)
		}

		// Drawing
		b.WriteString(charArtStr)
		lineArtWidth += charArtWidth

		// Stop coloring if we're at the
		// end of the text and still coloring
		if len(paintPositions) > 0 &&
			i == len(text)-1 &&
			paintPositions[0][1] == len(text) {
			b.WriteString("\033[0m")
			paintPositions = paintPositions[1:]
		}

		// Aligning
		if totalLineSpaces > 0 && char == ' ' {
			lineSpacesSeen++
			if lineSpacesSeen == totalLineSpaces {

				wordSpace = wordSpace[1:]
			}
			b.WriteString(wordSpace)
		}
	}
	return b.String()
}
