package rules

import (
	"fmt"
	"regexp"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
)

type ProfileNameFormatRule struct {
	// RegexFormat is used to validate the the profile name matches the given regex.
	// Note, when not set, or set to "", the rule will consider all names valid.
	RegexFormat *regexp.Regexp

	// FormatDescription is used to describe the regexFormat in the lint message. When not set,
	// or set to "", the regexFormat will be used.
	FormatDescription string
}

// ID() returns the rule ID.
func (*ProfileNameFormatRule) ID() string {
	return "profile-name-format"
}

// Message() returns the appropriate lint error message for this rule.
func (r *ProfileNameFormatRule) Message() string {
	if r.FormatDescription == "" {
		return fmt.Sprintf("Profile name must satisfy RegEx: %s", r.RegexFormat)
	}
	return fmt.Sprintf("Profile name must have format: %s", r.FormatDescription)
}

// Validate returns a *lint.Problem for each profile name that does not match the configured RegexFormat.
func (r *ProfileNameFormatRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem
	for _, profile := range fc.ParsedFSH.Profiles {
		p, err := profileNameFormatViolation(profile, r.ID(), r.Message(), r.RegexFormat)
		if err != nil {
			return nil, err
		}
		if p != nil {
			problems = append(problems, p)
		}
	}
	return problems, nil
}

// profileNameFormatViolation returns nil if the profile name matches the configured RegexFormat,
// and a *lint.Problem otherwise.
func profileNameFormatViolation(p *types.Profile, id string, msg string, regex *regexp.Regexp) (*lint.Problem, error) {
	if regex != nil && !regex.MatchString(p.Name.Value) {
		return lint.NewProblem(id, msg, p.Name.Location, nil, false)
	} else {
		return nil, nil
	}
}
