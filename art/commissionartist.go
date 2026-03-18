package art

import (
	"asciiart/oversabi"
	"strconv"
	"strings"
)

// CommissionArtist takes a text string and returns a function drawTextArt
// that draws the art of a text from start index to stop index (half open range).
//
// Note: the string that is passed must have all \n that represent newlines
// replaced by newline characters before being passed.
//
// Use:
//
//	str = strings.ReplaceAll(str, "\\n", "\n")
//
// to replace all \n occurences with newlines.
func CommissionArtist(text string) (drawTextArt func(start, stop int) string) {
	text = oversabi.Wrap(text) // For optionally wrapping the text
	const (
		moveUp7  = "\033[7A"
		moveDown = "\033[B"
	)
	newline := false
	drawTextArt = func(start, stop int) string {
		var b strings.Builder
		if start == 0 {
			newline = true
		}
		for _, char := range text[start:stop] {
			if char >= 32 && char < 127 || char == '\n' {
				charArt, charArtWidth := drawCharArt(byte(char))
				charArtStr := string(charArt)
				if charArtStr != "\n" {
					if newline {
						b.Write([]byte{'\n', '\n', '\n', '\n', '\n', '\n', '\n'})
						newline = false
					}
					b.WriteString(moveUp7)
					moveBackCAL := "\033[" + strconv.Itoa(charArtWidth) + "D"
					charArtStr = strings.Replace(charArtStr, "\n", moveDown+moveBackCAL, 7)
				} else {
					newline = true
				}
				b.WriteString(charArtStr)
			} else {
				panic("invalid ASCII character " + string(char))
			}
		}
		return b.String()
	}
	return
}
