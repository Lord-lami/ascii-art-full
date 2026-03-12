package paint

import "strings"

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
	case "default":
		return colorBrush + "39m"
	default:
		return ""
	}
}
