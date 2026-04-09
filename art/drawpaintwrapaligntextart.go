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
	if brush == "" {
		paintPositions = [][]int{}
	}
	lineArtWidth := 0
	const (
		moveUp7  = "\033[7A"
		moveDown = "\033[B"
	)
	var b strings.Builder
	var lineBuilder bytes.Buffer
	wordSpace := ""
	totalLineSpaces := 0
	lineSpacesSeen := 0

	var lineArtForm struct {
		LineArt              []string
		CharArt              [][]byte
		LineSpace, WordSpace string
		Brush, Reset         string
	}
	// Start Newline
	lineArtForm.LineArt = make([]string, 8)

	lineArtTmpl := template.New("")
	_, err := lineArtTmpl.ParseFiles("art/lineart.tmpl")
	if err != nil {
		panic(err)
	}

	for i, char := range text + "\n" {
		charArt, charArtWidth := drawCharArt(byte(char))

		if lineArtWidth != 0 {
			if lineArtWidth+charArtWidth >= terminalWidth || char == '\n' {
				if lineArtWidth == -1 {
					// Add Single Line
					b.WriteString(lineArtForm.LineArt[0])
					b.WriteByte('\n')
				} else {
					// Add Line
					b.WriteString(strings.Join(lineArtForm.LineArt, "\n"))
					b.WriteByte('\n')
				}

				// Start Newline
				lineArtForm.LineArt = make([]string, 8)
				lineArtWidth = 0
			}
		}

		if char != '\n' {
			// Fill CharArt
			lineArtForm.CharArt = bytes.Split(charArt, []byte{'\n'})
		} else {
			// Write nothing
			lineArtForm.CharArt = make([][]byte, 8)
			if lineArtWidth == 0 {
				lineArtWidth = -1
			}
		}

		// Aligning
		if lineArtWidth < 1 {
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
				lineArtForm.LineSpace = strings.Repeat(" ", shiftSize)
			case "center":
				shiftSize := (terminalWidth - nextLineArtWidth) / 2
				lineArtForm.LineSpace = strings.Repeat(" ", shiftSize)
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
		if len(paintPositions) > 0 && i == paintPositions[0][0] {
			lineArtForm.Brush = brush
		}
		if len(paintPositions) > 0 && i == paintPositions[0][1]-1 {
			lineArtForm.Reset = "\033[0m"
			paintPositions = paintPositions[1:]
		}

		// Stop coloring if we're at the
		// end of the text and still coloring
		if len(paintPositions) > 0 &&
			i == len(text)-1 &&
			paintPositions[0][1] == len(text) {
			lineArtForm.Reset = "\033[0m"
			paintPositions = paintPositions[1:]
		}

		fmt.Printf("%q %q %v\n", lineArtForm.Brush, lineArtForm.Reset, paintPositions)

		// Aligning
		if totalLineSpaces > 0 && char == ' ' {
			lineSpacesSeen++
			if lineSpacesSeen == totalLineSpaces {
				wordSpace = wordSpace[1:]
			}
			lineArtForm.WordSpace = wordSpace
		}

		// Drawing Line
		//
		// The commented out parts are for testing
		// with lineArtForm.LineArt as [][]byte
		// there was a bug where ExecuteTemplate seemed to
		// mutate lineArtForm.LineArt
		// Make the necessary changes to the code before testing
		//
		// h, _ := json.Marshal(lineArtForm)
		// var hold struct {
		// 	LineArt, CharArt     [][]byte
		// 	LineSpace, WordSpace string
		// 	Brush, Reset         string
		// }
		// json.Unmarshal(h, &hold)
		// fmt.Printf("\033[31m"+"Before:\nOriginal:%q\n" /*+"Copy    :%q\n\n"*/ +"\033[0m",lineArtForm /*, hold*/)
		err1 := lineArtTmpl.ExecuteTemplate(&lineBuilder, "lineart.tmpl", lineArtForm)
		// fmt.Printf("\033[32m"+"After:\nOriginal:%q\n" /*+"Copy    :%q\n\n"*/ +"\033[0m", lineArtForm /*, hold*/)
		if err1 != nil {
			panic(err1)
		}
		fmt.Printf("%q\n", lineArtForm.LineArt)
		fmt.Println(lineBuilder.String())
		fmt.Println()
		lineArtWidth += charArtWidth

		// Save Line
		lineArtForm.LineArt = strings.Split(lineBuilder.String(), "\n")
		fmt.Printf("changed:%q\n\n", lineArtForm.LineArt)

		// Erase
		lineBuilder.Reset()
		lineArtForm.CharArt = [][]byte{}
		lineArtForm.LineSpace = ""
		lineArtForm.WordSpace = ""
		lineArtForm.Brush = ""
		lineArtForm.Reset = ""
	}
	return b.String()
}
