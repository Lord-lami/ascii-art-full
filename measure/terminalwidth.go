// Package measure provides tools for measuring the width
// of art and the terminal.
package measure

import (
	"strings"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// TerminalWidth returns the current width of the terminal.
func TerminalWidth(text string) (width int) {
	ws := &winsize{}
	_, _, errorNbr := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(ws)),
	)
	if errorNbr != 0 {
		if errorNbr == 25 {
			return getLongestLineWidth(text)
		}
		panic("something is wrong with the terminal size. Error No." + errorNbr.Error())
	}
	return int(ws.Col)
}

func getLongestLineWidth(text string) int {
	longestLineWidth := 0
	for line := range strings.SplitSeq(text, "\\n") {
		lineLength := 0
		for _, char := range line {
			lineLength += GetCharArtWidth(byte(char))
		}
		longestLineWidth = max(lineLength, longestLineWidth)
	}
	return longestLineWidth + 1
}
