package rules

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
)

type CodeSystemNameMatchesFilenameRule struct {
	// NameSuffix is an optional suffix ignored during comparison between the
	// code system name and filename. Filenames should not include NameSuffix.
	// When NameSuffix is empty, the filename must match the code system name exactly.
	NameSuffix string
}

// ID() returns the rule ID.
func (*CodeSystemNameMatchesFilenameRule) ID() string {
	return "code-system-name-matches-filename"
}

// Message() returns the appropriate lint error message for this rule.
func (r *CodeSystemNameMatchesFilenameRule) Message() string {
	if r.NameSuffix == "" {
		return "Code system name must match filename."
	}
	return fmt.Sprintf("Code system name must match filename (where filename does not include %s).", r.NameSuffix)
}

// Validate returns a *lint.Problem for each code system name that does not match the filename.
func (r *CodeSystemNameMatchesFilenameRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, cs := range fc.ParsedFSH.CodeSystems {
		p, err := CodeSystemNameMatchesFilenameViolation(cs, fc.Path, r.NameSuffix, r.ID(), r.Message())
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// CodeSystemNameMatchesFilenameViolation returns nil if the code system name with the NameSuffix removed
// matches the filename, and a *lint.Problem otherwise.
func CodeSystemNameMatchesFilenameViolation(cs *types.CodeSystem, fp string, nameSuffix string, id string, msg string) (*lint.Problem, error) {
	filename := filepath.Base(fp)
	trimmedName := strings.TrimSuffix(cs.Name.Value, nameSuffix)
	if strings.TrimSuffix(filename, ".fsh") != trimmedName {
		// Code system name does not match filename
		diff := &lint.Diff{
			Got:       filename,
			Want:      trimmedName + ".fsh",
			FieldName: "Filename",
		}
		return lint.NewProblem(id, msg, cs.Name.Location, diff, false)
	}
	return nil, nil
}
