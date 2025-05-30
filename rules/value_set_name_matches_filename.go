package rules

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
)

type ValueSetNameMatchesFilenameRule struct {
	// NameSuffix is an optional suffix ignored during comparison between the
	// value set name and filename. Filenames should not end with NameSuffix.
	// When NameSuffix is empty, the filename must match the value set name exactly.
	NameSuffix string
}

// ID() returns the rule ID.
func (*ValueSetNameMatchesFilenameRule) ID() string {
	return "value-set-name-matches-filename"
}

// Message() returns the appropriate lint error message for this rule.
func (r *ValueSetNameMatchesFilenameRule) Message() string {
	if r.NameSuffix == "" {
		return "Value set name must match filename."
	}
	return fmt.Sprintf("Value set name must match filename (where filename does not include %s).", r.NameSuffix)
}

// Validate returns a *lint.Problem for each value set name that does not match the filename.
func (r *ValueSetNameMatchesFilenameRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, vs := range fc.ParsedFSH.ValueSets {
		p, err := ValueSetNameMatchesFilenameViolation(vs, fc.Path, r.NameSuffix, r.ID(), r.Message())
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// ValueSetNameMatchesFilenameViolation returns nil if the value set name with the NameSuffix removed
// matches the filename, and a *lint.Problem otherwise.
func ValueSetNameMatchesFilenameViolation(vs *types.ValueSet, fp string, nameSuffix string, id string, msg string) (*lint.Problem, error) {
	filename := filepath.Base(fp)
	trimmedName := strings.TrimSuffix(vs.Name.Value, nameSuffix)
	if strings.TrimSuffix(filename, ".fsh") != trimmedName {
		// Value set name does not match filename
		diff := &lint.Diff{
			Got:       filename,
			Want:      trimmedName + ".fsh",
			FieldName: "Filename",
		}

		return lint.NewProblem(id, msg, vs.Name.Location, diff, false)
	}
	return nil, nil
}
