/*
Package wrap provides string wrap formatting functionality, which enables
wrapping text to a specified width.
*/
package wrap

import "strings"

// Splitter is an abstraction for splitting a string into a series of tokens.
type Splitter interface {
	Split(string) []string
}

// SplitFunc is a function type that implements the Splitter interface.
type SplitFunc func(string) []string

// Split splits text into a series of tokens.
func (f SplitFunc) Split(text string) []string {
	return f(text)
}

var _ Splitter = (*SplitFunc)(nil)

// SpaceSplitter is a Splitter that splits text by unicode-identified space
// code points.
var SpaceSplitter Splitter = SplitFunc(strings.Fields)

// Wrapper is a text wrapper that wraps text to a specified width.
type Wrapper struct {
	Width     int
	Splitter  Splitter
	Separator string
}

// NewWrapper creates a new Wrapper with the specified width.
func NewWrapper(width int) *Wrapper {
	return &Wrapper{
		Width: width,
	}
}

// WithSplitter sets the Splitter for the Wrapper.
func (w *Wrapper) WithSplitter(splitter Splitter) *Wrapper {
	w.Splitter = splitter
	return w
}

// WithSeparator sets the separator for the Wrapper.
func (w *Wrapper) WithSeparator(separator string) *Wrapper {
	w.Separator = separator
	return w
}

// String returns a wrapped string with the specified width.
func (w *Wrapper) String(s string) string {
	return w.Strings(strings.Split(s, "\n")...)
}

// Strings returns a wrapped string with the specified width.
func (w *Wrapper) Strings(lines ...string) string {
	return strings.Join(w.Lines(lines...), "\n")
}

// Lines returns a wrapped slice of strings with the specified width.
func (w *Wrapper) Lines(lines ...string) []string {
	var tokens []string
	for _, line := range lines {
		if line == "" {
			tokens = append(tokens, "\n")
			continue
		}
		tokens = append(tokens, w.splitter().Split(line)...)
	}
	return w.reconstructLines(tokens)
}

func (w *Wrapper) splitter() Splitter {
	if w.Splitter == nil {
		return SpaceSplitter
	}
	return w.Splitter
}

func (w *Wrapper) separator() string {
	if w.Separator == "" {
		return " "
	}
	return w.Separator
}

func (w *Wrapper) reconstructLines(tokens []string) []string {
	separator := w.separator()

	if w.Width <= 0 {
		return []string{strings.Join(tokens, separator)}
	}

	var sb strings.Builder
	var lines []string
	for _, token := range tokens {
		if token == "\n" {
			lines = append(lines, sb.String())
			lines = append(lines, "")
			sb.Reset()
			continue
		}
		if sb.Len() == 0 {
			sb.WriteString(token)
			continue
		}
		if sb.Len()+len(token)+len(separator) < w.Width {
			sb.WriteString(separator)
		} else {
			lines = append(lines, sb.String())
			sb.Reset()
		}
		if len(token) >= w.Width {
			lines = append(lines, token)
			continue
		}
		sb.WriteString(token)
	}
	if sb.Len() > 0 {
		lines = append(lines, sb.String())
	}
	return lines
}
