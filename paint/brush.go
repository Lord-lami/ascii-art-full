package paint

// Brush accepts a color string that is either rgba or Hex color notation
// It returns a span that colors the string after it in html
func Brush(color string) (string, string) {
	if color == "default0" {
		return "", ""
	}
	return `<span style="color: ` + color + `;">`, `</span>`
}
