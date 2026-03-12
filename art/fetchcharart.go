// Package art provides tools to print the ascii art of a text.
package art

import (
	"asciiart/banners"
)

// fetchCharArt fetches the art of a character, char
// in banner file:
//
//	banners.BannerFile
func fetchCharArt(char byte) ([]byte, int) {
	charIndex := int(char - 32)
	bannerLinePosition := 1 + charIndex*9
	charArt := []byte{}
	lineLength := banners.BannerLineIndex[bannerLinePosition+1] -
		banners.BannerLineIndex[bannerLinePosition]
	for range 8 {
		offset := banners.BannerLineIndex[bannerLinePosition]
		rowArt := make([]byte, lineLength)
		_, err := banners.BannerFile.ReadAt(rowArt, offset)
		if err != nil {
			panic(err)
		}
		charArt = append(charArt, rowArt...)
		bannerLinePosition++
	}
	return charArt, int(lineLength - 1)
}
