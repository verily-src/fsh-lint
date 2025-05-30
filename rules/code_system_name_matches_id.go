package rules

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/internal/match"
	"github.com/verily-src/fsh-lint/lint"
)

type CodeSystemNameMatchesIDRule struct {
	// NameSuffix is an optional suffix ignored during comparison between the
	// code system name and ID. IDs should not include NameSuffix.
	// When NameSuffix is empty, the ID must match the code system name exactly.
	NameSuffix string
}

// ID() returns the rule ID.
func (*CodeSystemNameMatchesIDRule) ID() string {
	return "code-system-name-matches-id"
}

// Message() returns the appropriate lint error message for this rule.
func (r *CodeSystemNameMatchesIDRule) Message() string {
	if r.NameSuffix == "" {
		return "Code system name (PascalCase) must match code system ID in kebab-case."
	}
	return fmt.Sprintf("Code system name (PascalCase) must match code system ID in kebab-case without the %s suffix.", r.NameSuffix)
}

// Validate returns a *lint.Problem for each code system name that does not match
// its corresponding ID without the NameSuffix.
func (r *CodeSystemNameMatchesIDRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, cs := range fc.ParsedFSH.CodeSystems {
		p, err := codeSystemNameMatchesIDViolation(cs, r.NameSuffix, r.ID(), r.Message())
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// codeSystemNameMatchesIDViolation returns nil if the code system name
// matches ID in kebab-case without the NameSuffix, and a *lint.Problem otherwise.
func codeSystemNameMatchesIDViolation(cs *types.CodeSystem, nameSuffix string, id string, msg string) (*lint.Problem, error) {
	trimmedName := strings.TrimSuffix(cs.Name.Value, nameSuffix)
	if !match.IsNameKebabMatchWithID(trimmedName, cs.ID.Value, true) {
		// Code System Name does not match ID
		diff := &lint.Diff{
			Got:       cs.ID.Value,
			Want:      strcase.ToKebab(trimmedName),
			FieldName: "Code System ID",
		}

		return lint.NewProblem(id, msg, cs.ID.Location, diff, true)
	}
	return nil, nil
}
