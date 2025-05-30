package ansi

import (
	"bytes"
	"os"
	"regexp"
	"strconv"
)

var enabled bool

var (
	stripCodes = regexp.MustCompile("\033\\[([0-9;]+)m")

	buildFormatString func(attrs ...Attribute) string
	format            func(string, ...any) string
)

// IsEnabled returns true if ANSI color codes are enabled.
func IsEnabled() bool {
	return enabled
}

func init() {
	// Note: NO_COLOR is a defacto-standard for disabling color in bash environments,
	// NOCOLOR is also used by some tools. Both are supported here to disable
	// color-coding.
	if boolEnv("NO_COLOR") || boolEnv("NOCOLOR") {
		enabled = false
	} else {
		enabled = true
	}
}

func strip(data []byte) []byte {
	if !bytes.Contains(data, []byte(ControlSequenceIntroducer)) {
		return data
	}
	return stripCodes.ReplaceAll(data, nil)
}

func boolEnv(key string) bool {
	if v, ok := os.LookupEnv(key); ok {
		ok, _ := strconv.ParseBool(v)
		return ok
	}
	return false
}
