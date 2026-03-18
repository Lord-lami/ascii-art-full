package measure

import "asciiart/banners"

// GetStringArtWidth finds the width of the art of the passed
// string.
//
// Note: It will panic if a non ASCII character is passed to it.
func GetStringArtWidth(str string) (width int) {
	for _, char := range str {
		if char >= 32 && char < 127 || char == '\n' {
			if char == '\n' {
				continue
			}
			charIndex := int(char - 32)
			bannerLinePosition := 1 + charIndex*9
			width += int(banners.BannerLineIndex[bannerLinePosition+1]-
				banners.BannerLineIndex[bannerLinePosition]) - 1
		} else {
			panic("invalid ASCII character in GetStringArtWidth" + string(char))
		}
	}
	return
}
