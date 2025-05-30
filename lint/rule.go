package lint

// Rule represents a single FSH linting rule.
type Rule interface {
	// ID is the unique identifier of the rule
	// IDs should be unique across all rules in this package.
	// IDs should be in kebab-case and include characters a-z, 0-9, and hyphens.
	ID() string

	// Message is the error message that is given when the rules is violated.
	Message() string

	// Validate is the function that will be called to lint a file.
	// This method should return a list of problems found in the file and an error if one occurs.
	Validate(*FileContext) ([]*Problem, error)
}
