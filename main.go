package main

import (
	"asciiart/art"
	"asciiart/banners"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	// If a style is specified
	if len(args) > 1 {
		bannerFilePath := "./banners/" + args[1] + ".txt"

		// Check that a file for it exists and set it.
		var err error
		if banners.BannerFile, err = os.Open(bannerFilePath); err != nil {
			log.Fatal("there is no " + args[1] + " banner style file")
		}
	}

	if len(args) < 1 {
		log.Fatal("there is no text to draw")
	}

	if args[0] == "" {
		return
	}

	if args[0] == "\\n" {
		fmt.Println()
		return
	}

	defer banners.BannerFile.Close()
	banners.SetBannerLineIndex()

	text := strings.Split(args[0], "\\n")
	for _, line := range text {
		if line != "" {
			fmt.Print(art.GetLineArt(&line))
		} else {
			fmt.Println()
		}
	}
}
