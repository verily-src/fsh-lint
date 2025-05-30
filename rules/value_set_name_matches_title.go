package rules

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/internal/match"
	"github.com/verily-src/fsh-lint/lint"
)

type ValueSetNameMatchesTitleRule struct {
	// NameSuffix is an optional suffix ignored during comparison between the
	// value set name and title. Titles should not include NameSuffix.
	// When NameSuffix is empty, the title must match the value set name exactly.
	NameSuffix string
}

// ID() returns the rule ID.
func (*ValueSetNameMatchesTitleRule) ID() string {
	return "value-set-name-matches-title"
}

// Message() returns the appropriate lint error message for this rule.
func (r *ValueSetNameMatchesTitleRule) Message() string {
	if r.NameSuffix == "" {
		return "Value set name (PascalCase) must match value set title in Title Case."
	}
	return fmt.Sprintf("Value set name (PascalCase) must match value set title in Title Case without the %s suffix.", r.NameSuffix)
}

// Validate returns a *lint.Problem for each value set name without the NameSuffix
// that does not match its corresponding title.
func (r *ValueSetNameMatchesTitleRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, vs := range fc.ParsedFSH.ValueSets {
		p, err := ValueSetNameMatchesTitleViolation(vs, r.NameSuffix, r.ID(), r.Message())
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// ValueSetNameMatchesTitleViolation returns nil if the value set name without the NameSuffix matches title in
// Title Case, and a *lint.Problem otherwise.
func ValueSetNameMatchesTitleViolation(vs *types.ValueSet, nameSuffix string, id string, msg string) (*lint.Problem, error) {
	trimmedName := strings.TrimSuffix(vs.Name.Value, nameSuffix)
	spaceSeparatedName := strcase.ToDelimited(trimmedName, ' ')

	if !match.IsNameMatchWithTitle(trimmedName, vs.Title.Value, true) {
		caser := cases.Title(language.AmericanEnglish)
		diff := &lint.Diff{
			Got:       vs.Title.Value,
			Want:      caser.String(spaceSeparatedName),
			FieldName: "Value Set Title",
		}

		return lint.NewProblem(id, msg, vs.Title.Location, diff, true)
	}
	return nil, nil
}
