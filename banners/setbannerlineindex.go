package banners

import (
	"bufio"
	"os"
)

// BannerLineIndex is a slice of the offsets for each new line in a
// banner file. It is set by SetBannerLineIndex.
var BannerLineIndex []int64

// BannerFIle is the [os.File] that holds read access to the
// standard banner file or the specified banner
// file if it exists.
var BannerFile, _ = os.Open("./banners/standard.txt")

// SetBannerLineIndex runs through the
//
//	banners.BannerFile *os.File
//
// and populates
//
//	banners.BannerLineIndex []int64
//
// with the offsets for each line.
func SetBannerLineIndex() {
	var offset int64
	bannerScanner := bufio.NewScanner(BannerFile)
	for bannerScanner.Scan() {
		BannerLineIndex = append(BannerLineIndex, offset)
		offset += int64(len(bannerScanner.Bytes()) + 1)
	}
}
