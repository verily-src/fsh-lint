package rules

import (
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/internal/match"
	"github.com/verily-src/fsh-lint/lint"
)

const ProfileNameMatchesTitleID = "profile-name-matches-title"
const ProfileNameMatchesTitleMessage = `Profile name (PascalCase) must match profile title in Title Case (space seperated) and ending with "Profile".`

type ProfileNameMatchesTitleRule struct{}

// ID() returns the rule ID.
func (*ProfileNameMatchesTitleRule) ID() string {
	return ProfileNameMatchesTitleID
}

// Message() returns the appropriate lint error message for this rule.
func (*ProfileNameMatchesTitleRule) Message() string {
	return ProfileNameMatchesTitleMessage
}

// Validate returns a *lint.Problem for each profile name
// that does not match its corresponding title.
func (*ProfileNameMatchesTitleRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, profile := range fc.ParsedFSH.Profiles {
		p, err := profileNameMatchesTitleViolation(profile)
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// ProfileNameMatchesTitleViolation returns nil if the profile name matches title in
// Title Case with "Profile" at the end, and a *lint.Problem otherwise.
func profileNameMatchesTitleViolation(p *types.Profile) (*lint.Problem, error) {
	const loweredProfileSuffix = " profile"

	spaceSeparatedName := strcase.ToDelimited(p.Name.Value, ' ')
	loweredTitle := strings.ToLower(p.Title.Value)
	trimmedTitle := strings.TrimSuffix(loweredTitle, loweredProfileSuffix)

	// when the name ends in profile, we don't want title to end in "Profile Profile".
	trimmedName := strings.TrimSuffix(spaceSeparatedName, loweredProfileSuffix)

	if !match.IsNameMatchWithTitle(trimmedName, trimmedTitle, true) || !strings.HasSuffix(loweredTitle, " profile") {
		// Profile Name does not match Title or missing "Profile" at the end
		caser := cases.Title(language.AmericanEnglish)
		diff := &lint.Diff{
			Got:       p.Title.Value,
			Want:      caser.String(spaceSeparatedName),
			FieldName: "Profile Title",
		}

		if !strings.HasSuffix(diff.Want, " Profile") {
			diff.Want += " Profile"
		}

		return lint.NewProblem(ProfileNameMatchesTitleID, ProfileNameMatchesTitleMessage, p.Title.Location, diff, true)
	}
	return nil, nil
}
