package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/verily-src/fsh-lint/internal/fsh"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestProfileAssignmentPresent(t *testing.T) {
	tests := []struct {
		name        string
		element     string
		fsh         string
		wantProblem bool
	}{
		{
			name:    "element is present and set",
			element: "status",
			fsh: `Profile: Example
			Title: "Example"
			* ^status = #retired`,
			wantProblem: false,
		},
		{
			name:    "element is not present",
			element: "status",
			fsh: `Profile: Example
			Title: "Example"
			* ^notStatus = #retired`,
			wantProblem: true,
		},
		{
			name:    "element is present but not set",
			element: "status",
			fsh: `Profile: Example
			Title: "Example"
			* ^status = ""`,
			wantProblem: true,
		},
		{
			name:    "no element provided give no issues",
			element: "",
			fsh: `Profile: Example
			Title: "Example"
			* ^status = #retired`,
			wantProblem: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ruleName := "ProfileAssignmentPresentRule"

			parsedFSH, err := fsh.Parse(tt.fsh)
			if err != nil {
				t.Fatalf("error parsing fsh: %s", err)
			}

			sut := rules.ProfileAssignmentPresentRule{Element: tt.element}
			fileContext := &lint.FileContext{ParsedFSH: parsedFSH}
			problems, err := sut.Validate(fileContext)
			if err != nil {
				t.Fatalf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, nil)
			}

			got := len(problems) > 0
			if diff := cmp.Diff(got, tt.wantProblem); diff != "" {
				t.Errorf("%v.Validate() mismatch (-got +want):\n%s", ruleName, diff)
			}
		})
	}
}
