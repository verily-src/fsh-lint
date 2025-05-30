package ansi

import (
	"fmt"
	"strings"
)

// Format represents an ANSI format for text which is a collection of attributes.
type Format []Attribute

// String returns the ANSI format as a string.
func (f Format) String() string {
	return buildANSIFormat(f...)
}

// Attribute represents an ANSI Attribute formatting
type Attribute uint8

func (a Attribute) String() string {
	return buildANSIFormat(a)
}

// Format Attributes
const (
	// Reset is an ANSI attribute for resetting formatting
	Reset Attribute = 0

	// Bold is an ANSI attribute for setting a bold format
	Bold Attribute = 1

	// Faint is an ANSI attribute for setting a faint format
	Faint Attribute = 2

	// Italic is an ANSI attribute for setting an italic format
	Italic Attribute = 3

	// Underline is an ANSI attribute for setting an underline format
	Underline Attribute = 4

	// FGBlack is an ANSI attribute for the foreground color black
	FGBlack Attribute = 30

	// FGRed is an ANSI attribute for the foreground color red
	FGRed Attribute = 31

	// FGGreen is an ANSI attribute for the foreground color green
	FGGreen Attribute = 32

	// FGYellow is an ANSI attribute for the foreground color yellow
	FGYellow Attribute = 33

	// FGBlue is an ANSI attribute for the foreground color blue
	FGBlue Attribute = 34

	// FGMagenta is an ANSI attribute for the foreground color magenta
	FGMagenta Attribute = 35

	// FGCyan is an ANSI attribute for the foreground color cyan
	FGCyan Attribute = 36

	// FGWhite is an ANSI attribute for the foreground color white
	FGWhite Attribute = 37

	// FGGray is an ANSI attribute for the foreground color gray
	FGGray Attribute = 90

	// FGBrightRed is an ANSI attribute for the foreground color brightred
	FGBrightRed Attribute = 91

	// FGBrightGreen is an ANSI attribute for the foreground color brightgreen
	FGBrightGreen Attribute = 92

	// FGBrightYellow is an ANSI attribute for the foreground color brightyellow
	FGBrightYellow Attribute = 93

	// FGBrightBlue is an ANSI attribute for the foreground color brightblue
	FGBrightBlue Attribute = 94

	// FGBrightMagenta is an ANSI attribute for the foreground color brightmagenta
	FGBrightMagenta Attribute = 95

	// FGBrightCyan is an ANSI attribute for the foreground color brightcyan
	FGBrightCyan Attribute = 96

	// FGBrightWhite is an ANSI attribute for the foreground color brightwhite
	FGBrightWhite Attribute = 97

	// FGDefault is an ANSI attribute for setting the default foreground color
	FGDefault Attribute = 39

	// BGBlack is an ANSI attribute for the background color black
	BGBlack Attribute = 40

	// BGRed is an ANSI attribute for the background color red
	BGRed Attribute = 41

	// BGGreen is an ANSI attribute for the background color green
	BGGreen Attribute = 42

	// BGYellow is an ANSI attribute for the background color yellow
	BGYellow Attribute = 43

	// BGBlue is an ANSI attribute for the background color blue
	BGBlue Attribute = 44

	// BGMagenta is an ANSI attribute for the background color magenta
	BGMagenta Attribute = 45

	// BGCyan is an ANSI attribute for the background color cyan
	BGCyan Attribute = 46

	// BGWhite is an ANSI attribute for the background color white
	BGWhite Attribute = 47

	// BGGray is an ANSI attribute for the background color gray
	BGGray Attribute = 100

	// BGBrightRed is an ANSI attribute for the background color brightred
	BGBrightRed Attribute = 101

	// BGBrightGreen is an ANSI attribute for the background color brightgreen
	BGBrightGreen Attribute = 102

	// BGBrightYellow is an ANSI attribute for the background color brightyellow
	BGBrightYellow Attribute = 103

	// BGBrightBlue is an ANSI attribute for the background color brightblue
	BGBrightBlue Attribute = 104

	// BGBrightMagenta is an ANSI attribute for the background color brightmagenta
	BGBrightMagenta Attribute = 105

	// BGBrightCyan is an ANSI attribute for the background color brightcyan
	BGBrightCyan Attribute = 106

	// BGBrightWhite is an ANSI attribute for the background color brightwhite
	BGBrightWhite Attribute = 107

	// BGDefault is an ANSI attribute for setting the default background color
	BGDefault Attribute = 49
)

const (
	// ControlSequenceIntroducer is the prefix for ANSI escape commands.
	ControlSequenceIntroducer = "\033["

	// ControlCode is a rune representing the ANSI control code (0x1B, or 033).
	ControlCode rune = 033

	// Separator is the separator between different control codes.
	Separator = ';'

	// SGRSuffix is the suffix code for the Select Graphics Renderer control codes.
	SGRSuffix = 'm'

	// CursorUpSuffix is the suffix code for moving the cursor up.
	CursorUpSuffix = 'A'
	// CursorDownSuffix is the suffix code for moving the cursor down.
	CursorDownSuffix = 'B'
	// CursorRightSuffix is the suffix code for moving the cursor right.
	CursorRightSuffix = 'C'
	// CursorLeftSuffix is the suffix code for moving the cursor left.
	CursorLeftSuffix = 'D'
	// CursorNextLineSuffix is the suffix code for moving the cursor to the next line.
	CursorNextLineSuffix = 'E'
	// CursorPrevLineSuffix is the suffix code for moving the cursor to the previous line.
	CursorPrevLineSuffix = 'F'
)

func buildANSIFormat(attrs ...Attribute) string {
	sb := strings.Builder{}
	sb.WriteString(ControlSequenceIntroducer)
	toStr := func(v Attribute) string { return fmt.Sprintf("%d", v) }
	for _, attr := range attrs {
		sb.WriteString(toStr(attr))
		sb.WriteRune(Separator)
	}
	sb.WriteRune(SGRSuffix)
	return sb.String()
}
