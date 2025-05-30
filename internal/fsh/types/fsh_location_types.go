package types

import (
	"fmt"
)

// ParsedElement represents a parsed element that is tagged with
// location information (line and column start and end).
type ParsedElement[T any] struct {
	Value    T         `json:"value"`
	Location *Location `json:"location"`
}

// NewParsedElement returns a ParsedElement[T] with the given parameters and location data
func NewParsedElement[T any](value T, startLine int, startCol int, endLine int, endCol int) *ParsedElement[T] {
	return &ParsedElement[T]{
		Value: value,
		Location: &Location{
			Start: &Position{
				LineNumber:   startLine,
				ColumnNumber: startCol,
			},
			End: &Position{
				LineNumber:   endLine,
				ColumnNumber: endCol,
			},
		},
	}
}

// NewParsedElementWithoutLocation returns a ParsedElement[T] with the given parameters and location as nil
func NewParsedElementWithoutLocation[T any](value T) *ParsedElement[T] {
	return &ParsedElement[T]{
		Value: value,
	}
}

func (pe *ParsedElement[T]) String() string {
	return fmt.Sprintf(
		"ParsedElement{\n  Value: %v,\n  Location: %v\n}",
		pe.Value, pe.Location,
	)
}

// Location represents the range of position that an element spans.
type Location struct {
	Start *Position `json:"start"`
	End   *Position `json:"end"`
}

func (l *Location) String() string {
	return fmt.Sprintf(
		"Location{\n  Start: %v,\n  End: %v\n}",
		l.Start, l.End,
	)
}

// Position represents a position in a file.
type Position struct {
	LineNumber   int `json:"lineNumber"`
	ColumnNumber int `json:"columnNumber"`
}

func (p *Position) String() string {
	return fmt.Sprintf(
		"Position{\n  LineNumber: %d,\n  ColumnNumber: %d\n}",
		p.LineNumber, p.ColumnNumber,
	)
}
