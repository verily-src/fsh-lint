package diagnostic

import (
	"io"
	"os"

	"github.com/spf13/pflag"
	"github.com/verily-src/fsh-lint/internal/cli/format/ansi"
)

// MustInstallFlags installs the flags for the diagnostic package.
func MustInstallFlags(fs *pflag.FlagSet) {
	// Conditionally make the default to be 'github' if running from a GitHub
	// context, and always enables debug text since runners suppress it by default.
	defaultFormat := FormatText
	defaultDebug := false
	if _, ok := os.LookupEnv("GITHUB_EVENT_PATH"); ok {
		defaultFormat = FormatGitHub
	}
	if os.Getenv("RUNNER_DEBUG") == "1" {
		defaultDebug = true
	}
	fs.String("output-format", string(defaultFormat), "The format to use for printing diagnostic messages")
	fs.Bool("debug", defaultDebug, "Print debug information")
}

// ReporterFromFlags returns the printer based on the flags.
// The reporter will always use stderr by default, since reports need to be
// unbuffered messages.
func ReporterFromFlags(fs *pflag.FlagSet) (*Reporter, error) {
	return ReporterFromFlagsWithWriter(fs, os.Stderr)
}

// ReporterFromFlagsWithWriter returns the printer based on the flags with the
// given writer.
func ReporterFromFlagsWithWriter(fs *pflag.FlagSet, w io.Writer) (*Reporter, error) {
	raw, err := fs.GetString("output-format")
	if err != nil {
		return nil, err
	}
	w = ansi.DetectFormat(w)
	format, err := ParseFormat(raw)
	if err != nil {
		return nil, err
	}
	debug, err := fs.GetBool("debug")
	if err != nil {
		return nil, err
	}

	return format.Reporter(w).ShowDebug(debug), nil
}
