package rules

import (
	"fmt"

	"github.com/verily-src/fsh-lint/lint"
)

// ProfileAssignmentPresentRule will check that there exists an assignment rule
// (caret value rule) that sets the value of ProfileAssignmentPresentRule.Element.
type ProfileAssignmentPresentRule struct {
	// Element is the name of the property that is required to be set as an assignment
	// rule (caret value rule). Element is case sensitive. Required.
	Element string

	// AssignmentExample will be appended to the lint message to guide the user on
	// how to add a caret value rule that sets the value of Element. Optional.
	AssignmentExample string
}

// ID() returns the rule ID.
func (*ProfileAssignmentPresentRule) ID() string {
	return "profile-assignment-present"
}

// Message() returns the appropriate lint error message for this rule.
func (r *ProfileAssignmentPresentRule) Message() string {
	if r.AssignmentExample == "" {
		return fmt.Sprintf("Profile field '%s' must be set. Example: * ^%s = <value>", r.Element, r.Element)
	}
	return fmt.Sprintf("Profile field '%s' must be set. Example: %s", r.Element, r.AssignmentExample)
}

// Validate returns a *lint.Problem for each profile found that does not contain an assignment rule
// (caret value rule) that sets the value of ProfileAssignmentPresentRule.Element.
func (r *ProfileAssignmentPresentRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	// No issue if nothing needs to be set
	if r.Element == "" {
		return nil, nil
	}

	var problems []*lint.Problem
	for _, p := range fc.ParsedFSH.Profiles {
		hasElement := false
		for _, rule := range p.ProfileRules.CaretValueRules {
			if rule.Element != nil && rule.Element.Value == r.Element && rule.Value != nil && rule.Value.Value != "" {
				hasElement = true
				break
			}
		}

		if !hasElement {
			p, err := lint.NewProblem(r.ID(), r.Message(), nil, nil, false)

			if err != nil {
				return nil, err
			}
			problems = append(problems, p)
		}
	}

	return problems, nil
}
