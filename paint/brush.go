package paint

import "strings"

// brush accepts a color string.
// The color string must be one of: black, red, green, yellow, blue, magenta, cyan, white, default.
//
// It returns an ANSI Control Sequence that changes the color of the following text to the color
// specified by the string.
func brush(color string) string {
	color = strings.ToLower(color)
	colorBrush := "\033["
	switch color {
	case "black":
		return colorBrush + "30m"
	case "red":
		return colorBrush + "31m"
	case "green":
		return colorBrush + "32m"
	case "yellow":
		return colorBrush + "33m"
	case "blue":
		return colorBrush + "34m"
	case "magenta":
		return colorBrush + "35m"
	case "cyan":
		return colorBrush + "36m"
	case "white":
		return colorBrush + "37m"
	case "default", "default0":
		return colorBrush + "39m"
	default:
		panic("color must be one of: black, red, green, yellow, blue, magenta, cyan, white, default")
	}
}
