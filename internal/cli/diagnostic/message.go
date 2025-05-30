package diagnostic

import "fmt"

// Severity represents the severity of a diagnostic message.
type Severity string

const (
	// SeverityError represents an error message.
	SeverityError Severity = "error"

	// SeverityWarning represents a warning message.
	SeverityWarning Severity = "warning"

	// SeverityNotice represents a notice message.
	SeverityNotice Severity = "notice"

	// SeverityDebug represents a debug message.
	SeverityDebug Severity = "debug"
)

// Message represents a diagnostic message.
// The fields in this message map to the message fields as defined in
// https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/workflow-commands-for-github-actions
type Message struct {
	// Severity is the severity of the message (required).
	Severity Severity `json:"severity"`

	// Body is the body of the message (required).
	Body string `json:"body"`

	// Title is the title of the message (optional).
	Title string `json:"title,omitempty"`

	// File is the file name where the message originated (optional).
	File string `json:"file,omitempty"`

	// Line is the line number where the message originated (optional).
	Line int `json:"line,omitempty"`

	// Column is the column number where the message originated (optional).
	Column int `json:"column,omitempty"`

	// LineEnd is the end line number where the message originated (optional).
	LineEnd int `json:"line-end,omitempty"`

	// ColumnEnd is the end column number where the message originated (optional).
	ColumnEnd int `json:"column-end,omitempty"`
}

// Errorf creates a new error message with the given severity and body.
func Errorf(format string, args ...any) *Message {
	return NewMessage(SeverityError, format, args...)
}

// Warningf creates a new warning message with the given severity and body.
func Warningf(format string, args ...any) *Message {
	return NewMessage(SeverityWarning, format, args...)
}

// Noticef creates a new notice message with the given severity and body.
func Noticef(format string, args ...any) *Message {
	return NewMessage(SeverityNotice, format, args...)
}

// Debugf creates a new debug message with the given severity and body.
func Debugf(format string, args ...any) *Message {
	return NewMessage(SeverityDebug, format, args...)
}

// NewMessage creates a new message with the given severity and body.
func NewMessage(severity Severity, format string, args ...any) *Message {
	return &Message{
		Severity: severity,
		Body:     fmt.Sprintf(format, args...),
	}
}

// With applies the attachments to the message.
func (m *Message) With(attachments ...Attachment) *Message {
	for _, attachment := range attachments {
		attachment.set(m)
	}
	return m
}

// Attachment represents an attachment that can be applied to a diagnostic
// message.
//
// This follows the Go "options" pattern just to make it variadic and easy
// to set in a clean way.
type Attachment interface {
	set(*Message)
}

// Title returns an attachment that sets the title of the message.
func Title(title string) Attachment {
	return messageOption(func(m *Message) {
		m.Title = title
	})
}

// File returns an attachment that sets the file of the message.
func File(file string) Attachment {
	return messageOption(func(m *Message) {
		m.File = file
	})
}

// Line returns an attachment that sets the line of the message.
func Line(line int) Attachment {
	return messageOption(func(m *Message) {
		m.Line = line
	})
}

// LineRange returns an attachment that sets the line and end line of the message.
func LineRange(line, lineEnd int) Attachment {
	return messageOption(func(m *Message) {
		m.Line = line
		m.LineEnd = lineEnd
	})
}

// Column returns an attachment that sets the column of the message.
func Column(column int) Attachment {
	return messageOption(func(m *Message) {
		m.Column = column
	})
}

// ColumnRange returns an attachment that sets the column and end column of the message.
func ColumnRange(column, columnEnd int) Attachment {
	return messageOption(func(m *Message) {
		m.Column = column
		m.ColumnEnd = columnEnd
	})
}

type messageOption func(*Message)

func (o messageOption) set(m *Message) {
	o(m)
}
