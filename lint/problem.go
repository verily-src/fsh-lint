package lint

import (
	"fmt"
	"strings"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
)

var ProblemIsMisconfigured = fmt.Errorf("problem is misconfigured")

// Problem represents a single linting problem.
type Problem struct {
	// Message provides a description of the problem and how to fix it. Required.
	Message string

	// Location provides the location of the problem. Required.
	Location *types.Location

	// Diff provides the expected and found values of a field. Optional.
	Diff *Diff

	// RuleID provides the id of the rule this problem is violation of. Required.
	RuleID string

	// IsFixable indicates whether the problem can be automatically fixed. Required.
	IsFixable bool
}

// NewProblem creates a new Problem instance with the given parameters, and returns
// an error if it is misconfigured (e.g., if IsFixable is true but Diff or Location is nil).
func NewProblem(ruleID string, message string, location *types.Location, diff *Diff, isFixable bool) (*Problem, error) {
	if isFixable == true {
		if diff == nil {
			return nil, fmt.Errorf("diff must not be nil if problem is fixable: %w", ProblemIsMisconfigured)
		}
		if location == nil || location.Start == nil || location.End == nil {
			return nil, fmt.Errorf("location start and end must be provided if problem is fixable: %w", ProblemIsMisconfigured)
		}
	}

	return &Problem{
		Message:   message,
		Location:  location,
		Diff:      diff,
		RuleID:    ruleID,
		IsFixable: isFixable,
	}, nil
}

// StartPosition returns the start position of the problem, or nil if not available.
func (p *Problem) StartPosition() *types.Position {
	if p.Location == nil {
		return nil
	}
	return p.Location.Start
}

// EndPosition returns the end position of the problem, or nil if not available.
func (p *Problem) EndPosition() *types.Position {
	if p.Location == nil {
		return nil
	}
	return p.Location.End
}

// Diff holds the expected and found values of a field.
type Diff struct {
	Got       string
	Want      string
	FieldName string
}

// String returns the best string representation of the diff depending on which
// fields are set.
func (d *Diff) String() string {
	var parts []string

	if d.Got != "" {
		if d.FieldName != "" {
			parts = append(parts, fmt.Sprintf("Got %s: '%s'", d.FieldName, d.Got))
		} else {
			parts = append(parts, fmt.Sprintf("Got: '%s'", d.Got))
		}
	}

	if d.Want != "" {
		parts = append(parts, fmt.Sprintf("Want: '%s'", d.Want))
	}

	return strings.Join(parts, ", ")
}
