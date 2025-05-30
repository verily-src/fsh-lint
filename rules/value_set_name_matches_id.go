package rules

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/internal/match"
	"github.com/verily-src/fsh-lint/lint"
)

type ValueSetNameMatchesIDRule struct {
	// NameSuffix is an optional suffix ignored during comparison between the
	// value set name and ID. IDs should not include NameSuffix.
	// When NameSuffix is empty, the ID must match the value set name exactly.
	NameSuffix string
}

// ID() returns the rule ID.
func (*ValueSetNameMatchesIDRule) ID() string {
	return "value-set-name-matches-id"
}

// Message() returns the appropriate lint error message for this rule.
func (r *ValueSetNameMatchesIDRule) Message() string {
	if r.NameSuffix == "" {
		return "Value set name (PascalCase) must match value set ID in kebab-case."
	}
	return fmt.Sprintf("Value set name (PascalCase) must match value set ID in kebab-case without the %s suffix.", r.NameSuffix)
}

// Validate returns a *lint.Problem for each value set name that does not match
// its corresponding ID without the NameSuffix.
func (r *ValueSetNameMatchesIDRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, vs := range fc.ParsedFSH.ValueSets {
		p, err := valueSetNameMatchesIDViolation(vs, r.NameSuffix, r.ID(), r.Message())
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// valueSetNameMatchesIDViolation returns nil if the value set name with the NameSuffix removed
// matches the ID in kebab-case, and a *lint.Problem otherwise.
func valueSetNameMatchesIDViolation(vs *types.ValueSet, nameSuffix string, id string, msg string) (*lint.Problem, error) {
	trimmedName := strings.TrimSuffix(vs.Name.Value, nameSuffix)
	if !match.IsNameKebabMatchWithID(trimmedName, vs.ID.Value, true) {
		// Value Set Name does not match ID
		diff := &lint.Diff{
			Got:       vs.ID.Value,
			Want:      strcase.ToKebab(trimmedName),
			FieldName: "Value Set ID",
		}

		return lint.NewProblem(id, msg, vs.ID.Location, diff, true)
	}
	return nil, nil
}
