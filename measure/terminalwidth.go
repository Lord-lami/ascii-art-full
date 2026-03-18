// Package measure provides tools for measuring the width
// of art and the terminal.
package measure

import (
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
func TerminalWidth() (width int) {
	ws := &winsize{}
	_, _, errorNbr := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		syscall.TIOCGWINSZ,
		uintptr(unsafe.Pointer(ws)),
	)
	if errorNbr != 0 {
		panic("something is wrong with the terminal size")
	}
	return int(ws.Col)
}
