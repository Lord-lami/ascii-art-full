// Package art provides tools to print the ascii art of a text.
package art

import (
	"asciiart/banners"
)

// drawCharArt draws the art of a character, char
// from banner file:
//
//	banners.BannerFile
//
// and returns the art as a slice of [byte] and the
// width of the art as an [int]
func drawCharArt(char byte) ([]byte, int) {
	if char == '\n' {
		return []byte{'\n'}, 0
	}
	if char < 32 || char >= 127 {
		panic("invalid ASCII character: " + string(char))
	}
	charIndex := int(char - 32)
	bannerLinePosition := 1 + charIndex*9
	charArt := []byte{}
	charArtLength := banners.BannerLineIndex[bannerLinePosition+1] -
		banners.BannerLineIndex[bannerLinePosition]
	for range 8 {
		offset := banners.BannerLineIndex[bannerLinePosition]
		rowArt := make([]byte, charArtLength)
		_, err := banners.BannerFile.ReadAt(rowArt, offset)
		if err != nil {
			panic(err)
		}
		charArt = append(charArt, rowArt...)
		bannerLinePosition++
	}
	return charArt[:len(charArt)-1], int(charArtLength - 1)
}
