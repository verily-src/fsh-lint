package diagnostictest

import "github.com/verily-src/fsh-lint/internal/cli/diagnostic"

// FakePrinter is a fake implementation of diagnostic.Printer.
type FakePrinter struct {
	// Messages is the list of messages that have been printed.
	Messages []*diagnostic.Message
}

// Print appends the given message to the list of messages.
func (p *FakePrinter) Print(message *diagnostic.Message) {
	p.Messages = append(p.Messages, message)
}

// NewFakeReporter creates a new diagnostic.Reporter and a FakePrinter.
func NewFakeReporter() (*diagnostic.Reporter, *FakePrinter) {
	printer := &FakePrinter{}
	return diagnostic.NewReporter(printer), printer
}
