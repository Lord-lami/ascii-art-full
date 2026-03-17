// Package justify provides tools for justifying/arranging/aligning 
// the art of a text on the terminal.
package justify

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

// terminalWidth returns the current width of the terminal.
func terminalWidth() (width int) {
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
