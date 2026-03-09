package art

// GetLineArt returns the art of a line of text.
func GetLineArt(line *string) string {
	lineBytes := []byte{}
	for i := range byte(8) {
		lineBytes = append(lineBytes, getLineRowArt(line, i)...)
		lineBytes = append(lineBytes, '\n')
	}
	return string(lineBytes)
}
