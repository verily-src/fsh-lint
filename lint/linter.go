package lint

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/verily-src/fsh-lint/internal/cli/diagnostic"
)

// Linter orchestrates the linting process.
type Linter struct {
	// requiredRules is the set of rules that will validate that required rules
	// are present. Any lint problems found in a file from this rule set stops
	// the linter from running the rest of the rules. This is to simplify nil
	// error checks.
	requiredRules []Rule

	// rules is the set of rules to run after requiredRules has been run.
	rules []Rule

	// Formatter is responsible for formatting the generated lint Problems.
	Formatter Formatter

	// Reporter is responsible for reporting the problem.
	Reporter *diagnostic.Reporter

	// Fix is a flag that indicates whether the linter should attempt to fix
	// the problems found.
	Fix bool

	// HasErrors is a flag that indicates whether the linter has found any
	// error-level lint problems.
	HasErrors bool
}

// NewLinter initializes a new linter with the given required rules and rules.
// Sets the default formatter and reporter to nil to make it clear when it is
// intentionally set.
func NewLinter(requiredRules []Rule, rules []Rule) *Linter {
	return &Linter{
		requiredRules: requiredRules,
		rules:         rules,
		Formatter:     nil,
		Reporter:      nil,
		Fix:           false,
		HasErrors:     false,
	}
}

// Lint reads, parses, and validates the file at the given path reporting any
// issues found to Linter.reporter. If the fix flag is set, the linter will
// attempt to fix the problems found.
func (l *Linter) Lint(path string) {
	// Use default reporter and formatter if not set
	if l.Reporter == nil {
		l.Reporter = diagnostic.NewReporter(diagnostic.DefaultPrinter)
	}

	if l.Formatter == nil {
		l.Formatter = &DefaultFormatter{}
	}

	// read and parse file
	fileContext, err := NewFileContext(path)
	if err != nil {
		// return early if file can't be read/parsed
		l.Reporter.Errorf("%v", err)
		return
	}

	// validate that required rules are present
	var problems []*Problem
	missingFieldProblems := lintWithRules(fileContext, l.requiredRules, l.Reporter)
	problems = append(problems, missingFieldProblems...)

	// validate the rule set only if there are no missing fields
	if len(problems) == 0 {
		ruleProblems := lintWithRules(fileContext, l.rules, l.Reporter)
		problems = append(problems, ruleProblems...)
	}

	writeToFile := false
	for _, problem := range problems {
		message := makeMessage(problem, l.Formatter, path)
		l.Reporter.Report(message)

		if l.Fix {
			fixed, err := fixProblem(problem, fileContext)
			if err != nil {
				l.Reporter.Errorf("Error fixing problem: %v", err)
			}
			if fixed {
				log.Printf("[%s] Fixed %s in %s", problem.RuleID, problem.Diff.FieldName, fileContext.Path)
				writeToFile = true
			}
		}
	}

	if writeToFile {
		err = os.WriteFile(fileContext.Path, fileContext.Data, 0644)
		if err != nil {
			l.Reporter.Errorf("Error writing to %s: %v", fileContext.Path, err)
		}
	}

	if l.Reporter.ErrorCount() > 0 {
		l.HasErrors = true
	}
}

// lintWithRules runs the given rules on the given fileContext and returns the
// problems found. reporter is used to report any errors that occur while running.
func lintWithRules(fc *FileContext, rules []Rule, reporter *diagnostic.Reporter) []*Problem {
	var problems []*Problem
	for _, rule := range rules {
		p, err := rule.Validate(fc)
		if err != nil {
			if errors.Is(err, ProblemIsMisconfigured) {
				reporter.Debugf("Rule %s returned a misconfigured lint Problem: %v", rule.ID(), err)
			} else {
				reporter.Errorf("%v", err)
			}
		}
		problems = append(problems, p...)
	}
	return problems
}

// fixProblem updates fc.Data by fixing the problem in the file. Returns true
// if the problem was fixed, false otherwise.
func fixProblem(problem *Problem, fc *FileContext) (bool, error) {
	if !problem.IsFixable {
		return false, nil
	}

	startLine := problem.StartPosition().LineNumber
	endLine := problem.EndPosition().LineNumber
	want := problem.Diff.Want
	got := problem.Diff.Got

	fileLines := strings.Split(string(fc.Data), "\n")
	if startLine < 1 || endLine > len(fileLines) {
		return false, fmt.Errorf("Invalid line range for fix: %d-%d", startLine, endLine)
	}

	linesToFix := fileLines[startLine-1 : endLine]
	contentToFix := strings.Join(linesToFix, "\n")
	if strings.Count(contentToFix, got) != 1 {
		return false, fmt.Errorf("Cannot fix issue in file: %s, '%s' not found exactly once", fc.Path, got)
	}

	fixedContent := strings.Replace(contentToFix, got, want, 1)
	// Replace the lines in the file with the fixed content
	fileLines = append(
		fileLines[:startLine-1], // Lines before the start of the fix
		append(strings.Split(fixedContent, "\n"), // Fixed content split into lines
			fileLines[endLine:]...)..., // Lines after the end of the fix
	)
	fc.Data = []byte(strings.Join(fileLines, "\n"))

	return true, nil
}

// makeMessage creates a diagnostic message from the given problem using the formatter.
func makeMessage(problem *Problem, formatter Formatter, path string) *diagnostic.Message {
	var attachments []diagnostic.Attachment

	// Add available location data
	attachments = append(attachments, diagnostic.File(path))
	if start := problem.StartPosition(); start != nil {
		if end := problem.EndPosition(); end != nil {
			attachments = append(attachments, diagnostic.LineRange(start.LineNumber, end.LineNumber))
			attachments = append(attachments, diagnostic.ColumnRange(start.ColumnNumber, end.ColumnNumber))
		} else {
			attachments = append(attachments, diagnostic.Line(start.LineNumber))
			attachments = append(attachments, diagnostic.Column(start.ColumnNumber))
		}
	}

	msg := formatter.Format(problem)
	message := diagnostic.Noticef("%s", msg)
	return message.With(
		attachments...,
	)
}
