package rules

import (
	"path/filepath"
	"strings"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
)

const ProfileNameMatchesFilenameID = "profile-name-matches-filename"
const ProfileNameMatchesFilenameMessage = "Profile name must match filename."

type ProfileNameMatchesFilenameRule struct{}

// ID() returns the rule ID.
func (*ProfileNameMatchesFilenameRule) ID() string {
	return ProfileNameMatchesFilenameID
}

// Message() returns the appropriate lint error message for this rule.
func (*ProfileNameMatchesFilenameRule) Message() string {
	return ProfileNameMatchesFilenameMessage
}

// Validate returns a *lint.Problem for each profile name
// that does not match the filename.
func (*ProfileNameMatchesFilenameRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, profile := range fc.ParsedFSH.Profiles {
		p, err := profileNameMatchesFilenameViolation(profile, fc.Path)
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// ProfileNameMatchesFilenameViolation returns nil if the profile name
// matches filename exactly, and a *lint.Problem otherwise.
func profileNameMatchesFilenameViolation(p *types.Profile, fp string) (*lint.Problem, error) {
	filename := filepath.Base(fp)
	if strings.TrimSuffix(filename, ".fsh") != p.Name.Value {
		// Profile Name does not match filename
		diff := &lint.Diff{
			Got:       filename,
			Want:      p.Name.Value + ".fsh",
			FieldName: "Filename",
		}

		return lint.NewProblem(ProfileNameMatchesFilenameID, ProfileNameMatchesFilenameMessage, p.Name.Location, diff, false)
	}
	return nil, nil
}
