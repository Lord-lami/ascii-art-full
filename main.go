package main

import (
	"asciiart/banners"
	"asciiart/paint"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	usage := `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`
	// Patch for --color usage
	if len(os.Args) > 1 && (os.Args[1] == "--color" || os.Args[1] == "-color") {
		log.Fatal(usage)
	}
	color := flag.String("color", "default0", usage)
	flag.Parse()
	args := flag.Args()

	subStr := ""
	str := ""
	switch {
	// If a style is specified
	case len(args) > 2 || (*color == "default0" && len(args) == 2):
		bannerFileName := args[len(args)-1]
		bannerFilePath := "./banners/" + bannerFileName + ".txt"
		// Check that a file for it exists and set it.
		var err error
		if banners.BannerFile, err = os.Open(bannerFilePath); err != nil {
			log.Fatal("there is no " + bannerFileName + " banner style file")
		}
		fallthrough
	// Substring was provided
	case len(args) == 2:
		if *color == "default0" {
			str = args[0]
			break
		}
		subStr = args[0]
		str = args[1]
	// Substring was NOT provided
	case len(args) == 1:
		str = args[0]
	// No arguments
	case len(args) < 1:
		log.Fatal("there is no text to draw")
	}

	if str == "" {
		return
	}

	if str == "\\n" {
		fmt.Println()
		return
	}

	defer banners.BannerFile.Close()
	banners.SetBannerLineIndex()
	str = strings.ReplaceAll(str, "\\n", "\n")
	coloredTextArt := paint.PaintSubstring(str, subStr, *color, 0, len(str))
	if coloredTextArt == "" {
		log.Fatal("color must be one of: black, red, green, yellow, blue, magenta, cyan, white, default")
	}
	fmt.Println(coloredTextArt)
}

// func main() {
// 	defer banners.BannerFile.Close()
// 	banners.SetBannerLineIndex()
// 	fmt.Println(justify.Center("Tetst\nTest\n\nTest2", "t\n", "Blue"))
// }
