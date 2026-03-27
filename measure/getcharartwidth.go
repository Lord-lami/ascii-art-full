package measure

import "asciiart/banners"

// GetCharArtWidth finds the width of the art of the passed
// character.
//
// Note: It will panic if a non ASCII character is passed to it.
func GetCharArtWidth(char byte) (width int) {
	if char >= 32 && char < 127 {
		charIndex := int(char - 32)
		bannerLinePosition := 1 + charIndex*9
		width = int(banners.BannerLineIndex[bannerLinePosition+1]-
			banners.BannerLineIndex[bannerLinePosition]) - 1
	}
	return
}
