package main

import (
	"asciiart/art"
	"asciiart/banners"
	"asciiart/paint"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	colorUsage := `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`
	alignUsage := `Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard`
	color := flag.String("color", "default0", colorUsage)
	align := flag.String("align", "left", alignUsage)
	flag.Parse()
	// Patch for using flags without =
	for i := range flag.NFlag() {
		flagNames := [2]string{"--color", "--align"}
		for _, flagName := range flagNames {
			if os.Args[i+1] == flagName || os.Args[i+1] == flagName[1:] {
				flag.Usage()
				os.Exit(1)
			}
		}
	}

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
	// Set up
	defer banners.BannerFile.Close()
	banners.SetBannerLineIndex()
	str = strings.ReplaceAll(str, "\\n", "\n")

	//Draw, Paint, Wrap and Align
	brush := paint.Brush(*color)
	fmt.Printf("%q", art.DrawPaintWrapAlignTextArt(str, subStr, brush, *align))
}

// func main() {
// 	defer banners.BannerFile.Close()
// 	banners.SetBannerLineIndex()
// 	inputText := "TfeftrfstfthefOran tree"
// 	fmt.Println(inputText)
// 	paintedTextArt := paint.CommissionPainter(inputText, "t", "Blue")(0, len(inputText))
// 	fmt.Println(paintedTextArt)
// 	// fmt.Println(paintTextArt(0, 10))
// 	// fmt.Println(paintTextArt(9, 25))
// }
