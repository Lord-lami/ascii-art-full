package formatforfile

import (
	"fmt"
	"regexp"
	"strings"
)

func splitByRe(str *string, re *regexp.Regexp) (splits []string) {
	rePos := re.FindAllStringIndex(*str, -1)
	previousIdx := 0
	for _, pos := range rePos {
		splits = append(splits, (*str)[previousIdx:pos[0]])
		previousIdx = pos[1]
	}
	splits = append(splits, (*str)[previousIdx:])
	return
}

func FormatForFile(art *string) {
	lineArts = strings.Split(*art, "\n\n\n\n\n\n\n")
	moveUp7Re := regexp.MustCompile(regexp.QuoteMeta("\033[7A"))
	charArts := splitByRe(art, moveUp7Re)
	
	charArts = charArts[1:]
	// fmt.Println(moveUpPos)
	// fmt.Printf("%q\n", charArts)

	moveDownExp := regexp.QuoteMeta("\033[B")
	escapeExp := regexp.QuoteMeta("\033[")
	moveBackCALExp := escapeExp + `\d+D`

	rowSepRe := regexp.MustCompile(moveDownExp + moveBackCALExp)
	for _, charArt := range charArts {

	}

}
