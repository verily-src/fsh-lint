package diagnostic

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/verily-src/fsh-lint/internal/cli/format/ansi"
	"github.com/verily-src/fsh-lint/internal/cli/format/wrap"
	"golang.org/x/term"
)

// Printer represents a Printer that can print diagnostic messages.
type Printer interface {
	// Print prints the given message.
	Print(message *Message)
}

// DefaultPrinter is the default printer to use if no other printer is specified.
var DefaultPrinter Printer = &TextPrinter{W: os.Stdout}

// NewPrinter creates a new printer that writes to the given writer using the
// given formats.
// If no formats are provided, the printer defaults to the text format.
func NewPrinter(w io.Writer, format Format) Printer {
	switch format {
	case FormatGitHub:
		return &GitHubPrinter{W: w}
	case FormatText:
		return &ANSIPrinter{TextPrinter{W: w}}
	case FormatJSON:
		return &JSONPrinter{W: w}
	default:
		return DefaultPrinter
	}
}

// GitHubPrinter prints diagnostic messages in the format expected by GitHub
// Actions.
// See: https://docs.github.com/en/actions/reference/workflow-commands-for-github-actions#setting-an-error-message
type GitHubPrinter struct {
	// W is the writer to write the JSON output to. If not set, it defaults to
	// os.Stdout.
	W io.Writer
}

// Print prints the given message in the format expected by GitHub Actions.
func (p *GitHubPrinter) Print(message *Message) {
	keys := strings.Join(p.keys(message), ",")
	out := writerOrDefault(p.W)
	_, _ = fmt.Fprintf(out, "::%s %s::%s\n",
		message.Severity,
		keys,
		// See this comment in the supporting thread about multiline messages:
		// https://github.com/actions/toolkit/issues/193#issuecomment-605394935
		strings.ReplaceAll(message.Body, "\n", "%0A"),
	)
}

func writerOrDefault(w io.Writer) io.Writer {
	if w == nil {
		return os.Stdout
	}
	return w
}

// keys returns the keys for the message all joined by commas, which is the format
// expected by GitHub Actions.
func (p *GitHubPrinter) keys(message *Message) []string {
	var result []string
	if message.File != "" {
		result = append(result, fmt.Sprintf("file=%s", message.File))
	}
	if message.Line > 0 {
		result = append(result, fmt.Sprintf("line=%d", message.Line))
	}
	if message.Column > 0 {
		result = append(result, fmt.Sprintf("col=%d", message.Column))
	}
	if message.LineEnd > 0 {
		result = append(result, fmt.Sprintf("endLine=%d", message.LineEnd))
	}
	if message.ColumnEnd > 0 {
		result = append(result, fmt.Sprintf("endColumn=%d", message.ColumnEnd))
	}
	if message.Title != "" {
		result = append(result, fmt.Sprintf("title=%s", message.Title))
	}
	return result
}

// JSONPrinter prints diagnostic messages in JSON format.
type JSONPrinter struct {
	// W is the writer to write the JSON output to. If not set, it defaults to
	// os.Stdout.
	W io.Writer

	// Indent specifies whether the JSON output should be indented.
	Indent bool
}

// Print prints the given message in JSON format.
func (p *JSONPrinter) Print(message *Message) {
	out := writerOrDefault(p.W)
	var data []byte
	if p.Indent {
		data, _ = json.MarshalIndent(message, "", "  ")
	} else {
		data, _ = json.Marshal(message)
	}
	_, _ = fmt.Fprintf(out, "%s\n", string(data))
}

// TextPrinter prints diagnostic messages in plain text format.
// This will ignore any source-location information from the message and only
// print the severity and body.
type TextPrinter struct {
	// W is the writer to write the JSON output to. If not set, it defaults to
	// os.Stdout.
	W io.Writer
}

// Print prints the given message in plain text format.
func (p *TextPrinter) Print(message *Message) {
	out := writerOrDefault(p.W)

	var sb strings.Builder
	_, _ = fmt.Fprintf(&sb, "%s: ", message.Severity)
	buildMessageSuffix(&sb, message)
	_, _ = fmt.Fprintln(out, sb.String())
}

var _ Printer = (*TextPrinter)(nil)

// ANSIPrinter prints diagnostic messages in plain text format with ANSI color
// codes, if the underlying writer is a terminal. If it is not a terminal, this
// type behaves like [TextPrinter].
type ANSIPrinter struct {
	TextPrinter
}

// Print prints the given message in plain text format with ANSI color codes.
func (p *ANSIPrinter) Print(message *Message) {
	const (
		errorColor  = ansi.FGRed
		warnColor   = ansi.FGYellow
		noticeColor = ansi.FGCyan
		debugColor  = ansi.FGGreen
	)

	out := writerOrDefault(p.W)

	if fd, ok := out.(interface{ Fd() int64 }); ok {
		if !term.IsTerminal(int(fd.Fd())) {
			p.TextPrinter.Print(message)
			return
		}
	}

	var color ansi.Attribute
	switch message.Severity {
	case SeverityError:
		color = errorColor
	case SeverityWarning:
		color = warnColor
	case SeverityNotice:
		color = noticeColor
	case SeverityDebug:
		color = debugColor
	}
	var msg strings.Builder
	_, _ = fmt.Fprintf(&msg, "%v%s:%v ", color, message.Severity, ansi.Reset)
	buildMessageSuffix(&msg, message)
	_, _ = fmt.Fprintln(out, msg.String())
}

func buildMessageSuffix(sb *strings.Builder, message *Message) {
	prefix := strings.Repeat(" ", len(message.Severity)+2)
	if message.File != "" {
		_, _ = fmt.Fprintf(sb, "%v%s", ansi.Format{ansi.Underline, ansi.FGBrightWhite}, message.File)
		if message.Line > 0 {
			_, _ = fmt.Fprintf(sb, ":%d", message.Line)
			if message.Column > 0 {
				_, _ = fmt.Fprintf(sb, ":%d", message.Column)
			}
		}

		// If a file is specified, the error message prints on the next line, aligned
		// with the start of the filename (after the severity).
		_, _ = fmt.Fprintf(sb, "%v:\n", ansi.Reset)
		indent(sb, prefix, message.Body)
	} else {
		indent(sb, prefix, message.Body)
	}
}

func indent(sb *strings.Builder, prefix, text string) {
	wrapper := wrap.NewWrapper(80 - len(prefix))
	for i, line := range wrapper.Lines(strings.Split(text, "\n")...) {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(prefix)
		sb.WriteString(line)
	}
}

var _ Printer = (*ANSIPrinter)(nil)
