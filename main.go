package main

import (
	"asciiart/art"
	"asciiart/banners"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	colorUsage := `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --color=<color> <substring to be colored> "something"`
	alignUsage := `Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard`
	outputFileNameUsage := `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`
	color := flag.String("color", "default0", colorUsage)
	align := flag.String("align", "left", alignUsage)
	outputFileName := flag.String("output", "", outputFileNameUsage)
	flag.Parse()
	// Patch for using flags without =
	for i := range flag.NFlag() {
		flagNames := [3]string{"--color", "--align", "--output"}
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

	// Set up
	defer banners.BannerFile.Close()
	banners.SetBannerLineIndex()

	outputFile := os.Stdout
	if *outputFileName != "" {
		var err error
		outputFile, err = os.Create(*outputFileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Draw, Paint, Wrap and Align
	fmt.Fprintf(outputFile, "%s", art.DrawPaintWrapAlignTextArt(str, subStr, *color, *align))
}
