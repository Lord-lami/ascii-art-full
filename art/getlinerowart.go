package art

import (
	"bytes"
)

// getLineRowArt returns the (0 index) yth row arts
// of each byte in a line of text.
func getLineRowArt(line *string, y byte) []byte {
	ch := make(chan charRowArt)
	lineRowArt := make([][]byte, len(*line))

	for x := range *line {
		go fetchCharRowArtConc((*line)[x], [2]int{x, int(y)}, ch)
	}

	for range len(*line) {
		charRowArt := <-ch
		lineRowArt[charRowArt.CharIndex] = charRowArt.RowArt
	}
	close(ch)

	return bytes.Join(lineRowArt, []byte{})
}
