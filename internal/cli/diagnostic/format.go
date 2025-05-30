package diagnostic

import (
	"encoding"
	"fmt"
	"io"
)

// Format represents the format type of the diagnostic output.
type Format string

// ParseFormat parses the given string into a Format.
func ParseFormat(s string) (Format, error) {
	var result Format
	err := result.UnmarshalText([]byte(s))
	return result, err
}

// UnmarshalText unmarshals the given text into a [Format].
func (f *Format) UnmarshalText(text []byte) error {
	s := string(text)
	switch s {
	case string(FormatGitHub), string(FormatText), string(FormatJSON):
		*f = Format(s)
		return nil
	}
	return fmt.Errorf("invalid format: %s", s)
}

var _ encoding.TextUnmarshaler = (*Format)(nil)

const (
	// FormatGitHub represents the GitHub format.
	FormatGitHub Format = "github"

	// FormatText represents the text format.
	FormatText Format = "text"

	// FormatJSON represents the JSON format.
	FormatJSON Format = "json"
)

// Printer represents a printer that can print diagnostic messages.
func (f Format) Printer(w io.Writer) Printer {
	return NewPrinter(w, f)
}

// Reporter returns a new reporter that uses the given format.
func (f Format) Reporter(w io.Writer) *Reporter {
	return NewReporter(f.Printer(w))
}
