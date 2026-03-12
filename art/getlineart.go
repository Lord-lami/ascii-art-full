package art

import (
	"strconv"
	"strings"
)

// getLineArt returns the art of the first character in the string, line, 
// and a function that can print the art of the rest of the line if an 
// empty string ("") is passed to the function or another string 
// on the same line if an additional string, addition, is passed to the function.
func getLineArt(line string) (firstArt string, getTheRest func(addition string) string) {
	charArt, lineLength := fetchCharArt(line[0])
	firstArt = string(charArt)

	artLength := lineLength
	const (
		moveUp8  = "\033[8A"
		moveDown = "\033[B"
	)
	getTheRest = func(addition string) string {
		if addition == "" {
			addition = line[1:]
		}
		continuedArt := []byte{}
		for i := range addition {
			moveForwardX := "\033[" + strconv.Itoa(artLength) + "C"
			charArt, lineLength = fetchCharArt(addition[i])
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
