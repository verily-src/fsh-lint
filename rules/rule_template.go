package rules

import "github.com/verily-src/fsh-lint/lint"

type TemplateRule struct {
	// Add parameters to configure your rule here
}

// ID() returns the rule ID.
func (*TemplateRule) ID() string {
	return "rule-id-in-kebab-case"
}

// Message() returns the appropriate lint error message for this rule.
func (r *TemplateRule) Message() string {
	return "Message displayed when this rule is violated."
}

// Validate returns a *lint.Problem for each rule violation found.
func (r *TemplateRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	var problems []*lint.Problem

	for _, profile := range fc.ParsedFSH.Profiles { // Change this to the collection you want to iterate over
		if profile.Name.Value == "" { // Change this to the condition you want to check
			p, err := lint.NewProblem(r.ID(), r.Message(), nil, nil, false)
			if err != nil {
				return nil, err
			}
			problems = append(problems, p)
		}
	}

	return problems, nil
}
