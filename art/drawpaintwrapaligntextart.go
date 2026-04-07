package art

import (
	"asciiart/measure"
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
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

	var b strings.Builder
	wordSpace := ""
	lineSpace := ""
	totalLineSpaces := 0
	lineSpacesSeen := 0
	lineArtTemplate := template.New("")
	multiply := func(a, b int) int {
		return a * b
	}
	lineArtTemplate.Funcs(template.FuncMap{"multiply": multiply})
	_, err := lineArtTemplate.ParseFiles("art/text-templates/art.txt")
	if err != nil {
		panic(err)
	}

	var lineArtForm struct {
		LineArt      []byte
		LineArtWidth int
		CharArt      []byte
		CharArtWidth int
		Brush        string
		Reset        string
		LineSpace    string
		WordSpace    string
	}
	for i, char := range text + "\n" {
		charArt, charArtWidth := drawCharArt(byte(char))
		if lineArtWidth+charArtWidth >= terminalWidth {
			var ArtBuilder bytes.Buffer
			lineArtForm.CharArt = []byte{'\n', '\n', '\n', '\n', '\n', '\n', '\n', '\n'}
			lineArtForm.CharArtWidth = 1
			lineArtTemplate.ExecuteTemplate(&ArtBuilder, "art.txt", lineArtForm)
			lineArtForm.LineArt = ArtBuilder.Bytes()
			b.WriteString(lineSpace)
			b.Write(lineArtForm.LineArt)
			lineArtForm.LineArt = []byte{}
			lineArtForm.LineArtWidth = 0
			lineArtWidth = 0
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
				lineSpace = strings.Repeat(" ", shiftSize)
			case "center":
				shiftSize := (terminalWidth - nextLineArtWidth) / 2
				lineSpace = strings.Repeat(" ", shiftSize)
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
		if brush != "" {
			fmt.Println(paintPositions)
			if len(paintPositions) > 0 && i == paintPositions[0][1] {
				lineArtForm.Reset = "\033[0m"
				paintPositions = paintPositions[1:]
			}
			if len(paintPositions) > 0 && i == paintPositions[0][0] {
				lineArtForm.Brush = brush
			}
		}

		// Drawing
		if char == '\n' && len(lineArtForm.LineArt) == 0 ||
			len(lineArtForm.LineArt) > 0 &&
				lineArtForm.LineArt[len(lineArtForm.LineArt)-1] == '\n' {
			lineArtForm.LineArt = append(lineArtForm.LineArt, '\n')
		} else {
			var ArtBuilder bytes.Buffer
			lineArtForm.CharArt = charArt
			lineArtForm.CharArtWidth = charArtWidth
			lineArtTemplate.ExecuteTemplate(&ArtBuilder, "art.txt", lineArtForm)
			lineArtForm.LineArt = ArtBuilder.Bytes()
			lineArtForm.LineArtWidth += charArtWidth
			fmt.Println(lineArtForm.LineArtWidth)
			lineArtWidth += charArtWidth
		}
		lineArtForm.Brush = ""
		lineArtForm.Reset = ""

		// Drawing
		// b.WriteString(charArtStr)
		// lineArtWidth += charArtWidth

		// if i == len(text)-1 {
		// 	// Stop coloring if we're at the
		// 	// end of the text and still coloring
		// 	if len(paintPositions) > 0 && paintPositions[0][1] == len(text) {
		// 		lineArtForm.LineArt = append(lineArtForm.LineArt, []byte("\033[0m")...)
		// 		paintPositions = paintPositions[1:]
		// 	}
		// }

		// Aligning
		if totalLineSpaces > 0 && char == ' ' {
			lineSpacesSeen++
			if lineSpacesSeen == totalLineSpaces {

				wordSpace = wordSpace[1:]
			}
			lineArtForm.LineArt = append(lineArtForm.LineArt, []byte(wordSpace)...)
		}

		if char == '\n' {
			b.WriteString(lineSpace)
			b.Write(lineArtForm.LineArt)
			lineArtForm.LineArt = []byte{}
			lineArtForm.LineArtWidth = 0
			lineArtWidth = 0
		}
	}
	return b.String()
}
