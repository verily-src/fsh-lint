package diagnostic

import "os"

// Reporter represents an emitter that can emit diagnostic messages.
type Reporter struct {
	printer Printer

	// count of each emitted message type
	errors, warnings, notices, debug int

	enableDebug bool
}

// NewReporter creates a new reporter with the given printer.
func NewReporter(printer Printer) *Reporter {
	return &Reporter{printer: printer}
}

// ShowDebug sets whether debug messages should be emitted.
func (r *Reporter) ShowDebug(enable bool) *Reporter {
	r.enableDebug = enable
	return r
}

// Report emits the given message.
func (r *Reporter) Report(message *Message) {
	switch message.Severity {
	case SeverityError:
		r.errors++
	case SeverityWarning:
		r.warnings++
	case SeverityNotice:
		r.notices++
	case SeverityDebug:
		r.debug++
		if !r.enableDebug {
			return
		}
	}
	r.getPrinter().Print(message)
}

// ReportFatal emits a fatal error and exits the program.
func (r *Reporter) ReportFatal(message *Message) {
	msg := *message
	msg.Severity = SeverityError
	r.Report(&msg)
	os.Exit(1)
}

// Fatalf emits a fatal error and exits the program.
func (r *Reporter) Fatalf(format string, args ...any) {
	r.Report(Errorf(format, args...))
	os.Exit(1)
}

// Errorf emits an error message.
func (r *Reporter) Errorf(format string, args ...any) {
	r.Report(Errorf(format, args...))
}

// Warningf emits a warning message.
func (r *Reporter) Warningf(format string, args ...any) {
	r.Report(Warningf(format, args...))
}

// Noticef emits a notice message.
func (r *Reporter) Noticef(format string, args ...any) {
	r.Report(Noticef(format, args...))
}

// Debugf emits a debug message.
func (r *Reporter) Debugf(format string, args ...any) {
	r.Report(Debugf(format, args...))
}

func (r *Reporter) getPrinter() Printer {
	if r.printer == nil {
		return DefaultPrinter
	}
	return r.printer
}

// ErrorCount returns the number of errors emitted.
func (r *Reporter) ErrorCount() int {
	return r.errors
}

// WarningCount returns the number of warnings emitted.
func (r *Reporter) WarningCount() int {
	return r.warnings
}

// NoticeCount returns the number of notices emitted.
func (r *Reporter) NoticeCount() int {
	return r.notices
}
