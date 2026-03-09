// Package art provides tools to print the ascii art of a text.
package art

import (
	"asciiart/banners"
)

// charRowArt holds the character row art of a byte and
// the character index of the byte in its string.
type charRowArt struct {
	RowArt    []byte
	CharIndex int
}

// fetchCharRowArtConc fetches the row art of a character, char,
// at the position, position, in the string from the file
//
//	banners.BannerFile
//
// and channels it to the channel, ch.
func fetchCharRowArtConc(char byte, position [2]int, ch chan charRowArt) {
	if char < 32 {
		panic("control characters are illegal except newlines")
	}
	charIndex := int(char - 32)
	bannerLinePosition := 1 + position[1] + charIndex*9
	offset := banners.BannerLineIndex[bannerLinePosition]
	linelength := banners.BannerLineIndex[bannerLinePosition+1] - offset - 1
	rowArt := make([]byte, linelength)
	_, err := banners.BannerFile.ReadAt(rowArt, offset)
	if err != nil {
		panic(err)
	}

	charArt := charRowArt{RowArt: rowArt, CharIndex: position[0]}
	ch <- charArt
}
