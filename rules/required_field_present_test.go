package rules_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestRequiredFieldPresent(t *testing.T) {
	tests := []struct {
		name        string
		fieldPath   string
		fshDocument *types.FSHDocument
		wantProblem bool
	}{
		{
			name:      "Path points to a *ParsedElement[string], and set to nil",
			fieldPath: "Profiles.ID",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value: "Profile",
						},
					},
				},
			},
			wantProblem: true,
		},
		{
			name:      "Path points to a *ParsedElement[string], and not nil",
			fieldPath: "Profiles.ID",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value: "Profile",
						},
						ID: &types.ParsedElement[string]{
							Value: "profile",
						},
					},
				},
			},
			wantProblem: false,
		},

		{
			name:      "Path points to a *ParsedElement[string], and value is set to empty string",
			fieldPath: "Profiles.ID",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value: "Profile",
						},
						ID: &types.ParsedElement[string]{
							Value: "",
						},
					},
				},
			},
			wantProblem: true,
		},
		{
			name:      "Path points to a *ParsedElement[string], and value is set to non-empty string",
			fieldPath: "Profiles.ID",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value: "Profile",
						},
						ID: &types.ParsedElement[string]{
							Value: "profile",
						},
					},
				},
			},
			wantProblem: false,
		},
		{
			name:      "Path points to a string set to empty",
			fieldPath: "ValueSets.Description.Value",
			fshDocument: &types.FSHDocument{
				ValueSets: []*types.ValueSet{
					{
						Description: &types.ParsedElement[string]{
							Value: "",
						},
					},
				},
			},
			wantProblem: true,
		},
		{
			name:      "Path points to a string set to non-empty",
			fieldPath: "ValueSets.Description.Value",
			fshDocument: &types.FSHDocument{
				ValueSets: []*types.ValueSet{
					{
						Description: &types.ParsedElement[string]{
							Value: "Description.",
						},
					},
				},
			},
			wantProblem: false,
		},
		{
			name:      "Path points to an empty slice",
			fieldPath: "ValueSets.IncludeComponents",
			fshDocument: &types.FSHDocument{
				ValueSets: []*types.ValueSet{
					{
						IncludeComponents: []*types.ValueSetComponent{},
					},
				},
			},
			wantProblem: true,
		},
		{
			name:      "Path points to a non empty slice",
			fieldPath: "ValueSets.IncludeComponents",
			fshDocument: &types.FSHDocument{
				ValueSets: []*types.ValueSet{
					{
						IncludeComponents: []*types.ValueSetComponent{
							{},
						},
					},
				},
			},
			wantProblem: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ruleName := "RequiredFieldPresentRule"

			sut := rules.RequiredFieldPresentRule{FieldPath: tt.fieldPath}
			fileContext := &lint.FileContext{ParsedFSH: tt.fshDocument}
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

func TestRequiredFieldPresentConfigurationErrors(t *testing.T) {
	tests := []struct {
		name        string
		fieldPath   string
		fshDocument *types.FSHDocument
		wantError   error
	}{
		{
			name:      "empty field path",
			fieldPath: "",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value: "Profile",
						},
					},
				},
			},
			wantError: rules.ErrFieldPathEmpty,
		},
		{
			name:      "invalid field",
			fieldPath: "Profiles.NotARealField",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{},
				},
			},
			wantError: rules.ErrInvalidField,
		},
		{
			name:      "field not recursable anymore",
			fieldPath: "Profiles.Name.Value.TooFarIntoRecursion",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value: "Profile",
						},
					},
				},
			},
			wantError: rules.ErrNotRecursable,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ruleName := "RequiredFieldPresentRule"
			sut := rules.RequiredFieldPresentRule{FieldPath: tt.fieldPath}
			fileContext := &lint.FileContext{ParsedFSH: tt.fshDocument}

			_, err := sut.Validate(fileContext)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, tt.wantError)
			}
		})
	}
}
