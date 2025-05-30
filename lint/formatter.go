package lint

import "fmt"

// Formatter is responsible for formatting the generated lint Problems to strings.
type Formatter interface {
	// Format takes in a problem and returns a formatted string of that problem.
	Format(*Problem) string
}

type DefaultFormatter struct{}

// Format formats the given problem into a human-readable string.
func (d *DefaultFormatter) Format(p *Problem) string {
	if p.Diff != nil {
		return fmt.Sprintf("[%s] %s. %s", p.RuleID, p.Diff, p.Message)
	}
	return fmt.Sprintf("[%s] %s", p.RuleID, p.Message)
}
