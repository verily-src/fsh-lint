package ansi

import (
	"io"

	"golang.org/x/term"
)

// CanFormat returns true if the given writer supports color and xterm output.
func CanFormat(w io.Writer) bool {
	if !IsEnabled() {
		return false
	}

	if fd, ok := w.(interface{ Fd() uintptr }); ok {
		return term.IsTerminal(int(fd.Fd()))
	}
	return false
}

// NoFormat strips format codes from the output.
func NoFormat(w io.Writer) io.Writer {
	return writeFunc(func(p []byte) (n int, err error) {
		return w.Write(strip(p))
	})
}

// DetectFormat detects if the given writer supports color output, and configures
// the writer to output color if it does.
func DetectFormat(w io.Writer) io.Writer {
	if CanFormat(w) {
		return w
	}
	return NoFormat(w)
}

type writeFunc func(p []byte) (n int, err error)

func (f writeFunc) Write(p []byte) (n int, err error) {
	return f(p)
}
