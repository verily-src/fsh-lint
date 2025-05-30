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

type CodeSystemNameMatchesTitleRule struct {
	// NameSuffix is an optional suffix ignored during comparison between the
	// code system name and title. Titles should not include NameSuffix.
	// When NameSuffix is empty, the title must match the code system name exactly.
	NameSuffix string
}

// ID() returns the rule ID.
func (*CodeSystemNameMatchesTitleRule) ID() string {
	return "code-system-name-matches-title"
}

// Message() returns the appropriate lint error message for this rule.
func (r *CodeSystemNameMatchesTitleRule) Message() string {
	if r.NameSuffix == "" {
		return "Code system name (PascalCase) must match code system title in Title Case."
	}
	return fmt.Sprintf("Code system name (PascalCase) must match code system title in Title Case without the %s suffix.", r.NameSuffix)
}

// Validate returns a *lint.Problem for each code system name without the NameSuffix
// that does not match its corresponding title.
func (r *CodeSystemNameMatchesTitleRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, cs := range fc.ParsedFSH.CodeSystems {
		p, err := codeSystemNameMatchesTitleViolation(cs, r.NameSuffix, r.ID(), r.Message())
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// codeSystemNameMatchesTitleViolation returns nil if the code system name without the NameSuffix matches title in
// Title Case, and a *lint.Problem otherwise.
func codeSystemNameMatchesTitleViolation(cs *types.CodeSystem, nameSuffix string, id string, msg string) (*lint.Problem, error) {
	trimmedName := strings.TrimSuffix(cs.Name.Value, nameSuffix)
	spaceSeparatedName := strcase.ToDelimited(trimmedName, ' ')

	if !match.IsNameMatchWithTitle(spaceSeparatedName, cs.Title.Value, true) {
		caser := cases.Title(language.AmericanEnglish)
		diff := &lint.Diff{
			Got:       cs.Title.Value,
			Want:      caser.String(spaceSeparatedName),
			FieldName: "Code System Title",
		}

		return lint.NewProblem(id, msg, cs.Title.Location, diff, true)
	}
	return nil, nil
}
