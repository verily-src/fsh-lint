package rules

import (
	"github.com/iancoleman/strcase"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/internal/match"
	"github.com/verily-src/fsh-lint/lint"
)

const ProfileNameMatchesIDID = "profile-name-matches-id"
const ProfileNameMatchesIDMessage = "Profile name (PascalCase) must match profile id in kebab-case."

type ProfileNameMatchesIDRule struct{}

// ID() returns the rule ID.
func (*ProfileNameMatchesIDRule) ID() string {
	return ProfileNameMatchesIDID
}

// Message() returns the appropriate lint error message for this rule.
func (*ProfileNameMatchesIDRule) Message() string {
	return ProfileNameMatchesIDMessage
}

// Validate returns a *lint.Problem for each profile name
// that does not match its corresponding ID.
func (*ProfileNameMatchesIDRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, profile := range fc.ParsedFSH.Profiles {
		p, err := profileNameMatchesIDViolation(profile)
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// ProfileNameMatchesIDViolation returns nil if the profile name
// matches id exactly, and a *lint.Problem otherwise.
func profileNameMatchesIDViolation(p *types.Profile) (*lint.Problem, error) {
	if !match.IsNameKebabMatchWithID(p.Name.Value, p.ID.Value, true) {
		// Profile Name does not match ID
		diff := &lint.Diff{
			Got:       p.ID.Value,
			Want:      strcase.ToKebab(p.Name.Value),
			FieldName: "Profile ID",
		}

		return lint.NewProblem(ProfileNameMatchesIDID, ProfileNameMatchesIDMessage, p.ID.Location, diff, true)
	}
	return nil, nil
}
