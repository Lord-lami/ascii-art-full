package art

import (
	"strconv"
	"strings"
)

// getLineArt returns the art of a line of text.
func getLineArt(line string) (firstArt string, getTheRest func(remaining string) string) {
	charArt, lineLength := fetchCharArt(line[0])
	firstArt = string(charArt)

	artLength := lineLength
	const (
		moveUp8  = "\033[8A"
		moveDown = "\033[B"
	)
	getTheRest = func(remaining string) string {
		if remaining == "" {
			remaining = line[1:]
		}
		continuedArt := []byte{}
		for i := range remaining {
			moveForwardX := "\033[" + strconv.Itoa(artLength) + "C"
			charArt, lineLength = fetchCharArt(remaining[i])
			moveBackL := "\033[" + strconv.Itoa(lineLength) + "D"
			charArtStr := strings.Replace(string(charArt), "\n", moveDown+moveBackL, 7)
			continuedArt = append(continuedArt, []byte(moveUp8+moveForwardX)...)
			continuedArt = append(continuedArt, []byte(charArtStr)...)
			artLength += lineLength
		}
		return string(continuedArt)
	}
	return
}
