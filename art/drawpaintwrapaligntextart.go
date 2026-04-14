package art

import (
	"asciiart/measure"
	"asciiart/paint"
	"bytes"
	"embed"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

// Embedding
//
//go:embed lineart.tmpl
var templateFS embed.FS

// DrawPaintWrapAlignTextArt takes a text string and returns
// the art of a text from start index to stop index (half open range).
//
// Note: the strings that are passed will have all `\n`s`
// replaced by newline characters.
func DrawPaintWrapAlignTextArt(text, subStr, color, alignment string) string {
	// Production
	lineArtTmpl, err := template.ParseFS(templateFS, "lineart.tmpl")

	// Development
	// lineArtTmpl, err := template.New("").ParseFiles("art/lineart.tmpl")

	terminalWidth := measure.TerminalWidth(text)
	brush, reset := paint.Brush(color)

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

	var b strings.Builder

	wordSpace := ""
	totalLineSpaces := 0
	lineSpacesSeen := 0

	var lineArtForm struct {
		LineArt              []string
		CharArt              [][]byte
		LineSpace, WordSpace string
		Brush, Reset         string
	}
	var lineBuilder bytes.Buffer
	lineArtWidth := 0
	// Start Newline
	lineArtForm.LineArt = make([]string, 8)

	if err != nil {
		panic(err)
	}

	for i, char := range text + "\n" {
		charArt, charArtWidth := drawCharArt(byte(char))

		if char != '\n' {
			// Fill CharArt
			lineArtForm.CharArt = bytes.Split(charArt, []byte{'\n'})
		} else {
			// Fill CharArt with nothing
			lineArtForm.CharArt = make([][]byte, 8)
		}

		if lineArtWidth+charArtWidth >= terminalWidth || char == '\n' {
			if lineArtWidth == 0 {
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

		// Aligning
		if lineArtWidth == 0 && char != '\n' {
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
		fmt.Println(paintPositions)
		if len(paintPositions) > 0 && i == paintPositions[0][0] {
			lineArtForm.Brush = brush
			lineArtForm.Reset = reset
			if len(paintPositions) > 1 && paintPositions[1][0] == paintPositions[0][1] {
				paintPositions = paintPositions[1:]
			}
		} else if len(paintPositions) > 0 && i == paintPositions[0][1] {
			lineArtForm.Brush = ""
			lineArtForm.Reset = ""
			paintPositions = paintPositions[1:]
		}

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
		err1 := lineArtTmpl.Execute(&lineBuilder, lineArtForm)
		// fmt.Printf("\033[32m"+"After:\nOriginal:%q\n" /*+"Copy    :%q\n\n"*/ +"\033[0m", lineArtForm /*, hold*/)
		if err1 != nil {
			panic(err1)
		}
		lineArtWidth += charArtWidth

		// Save Line
		lineArtForm.LineArt = strings.Split(lineBuilder.String(), "\n")

		// Erase
		lineBuilder.Reset()
		lineArtForm.CharArt = [][]byte{}
		lineArtForm.LineSpace = ""
		lineArtForm.WordSpace = ""
	}
	return b.String()
}
