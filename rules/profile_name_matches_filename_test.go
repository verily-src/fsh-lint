package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintProfileNameMatchesFilename(t *testing.T) {
	testLocation := &types.Location{
		Start: &types.Position{
			LineNumber:   1,
			ColumnNumber: 2,
		},
		End: &types.Position{
			LineNumber:   1,
			ColumnNumber: 4,
		},
	}

	tests := []struct {
		name        string
		fileContext lint.FileContext
		want        []*lint.Problem
	}{
		{
			name: "no profiles",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{}},
				Path: "path/to/Profile.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "exact match",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfile",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleProfile.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "exact match no path",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfile",
								Location: testLocation,
							},
						},
					},
				},
				Path: "ExampleProfile.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "exact match, path in root",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfile",
								Location: testLocation,
							},
						},
					},
				},
				Path: "/ExampleProfile.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "not match",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfileOne",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleProfileTwo.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesFilenameID,
					Message:  rules.ProfileNameMatchesFilenameMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleProfileTwo.fsh",
						Want:      "ExampleProfileOne.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
		{
			name: "empty profile name",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleProfileTwo.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesFilenameID,
					Message:  rules.ProfileNameMatchesFilenameMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleProfileTwo.fsh",
						Want:      ".fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
		{
			name: "one matching, one non matching profile",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfileOne",
								Location: testLocation,
							},
						},
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfileTwo",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleProfileOne.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesFilenameID,
					Message:  rules.ProfileNameMatchesFilenameMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleProfileOne.fsh",
						Want:      "ExampleProfileTwo.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
		{
			name: "two non matching profiles",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: []*types.Profile{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfileOne",
								Location: testLocation,
							},
						},
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleProfileTwo",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleProfileThree.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesFilenameID,
					Message:  rules.ProfileNameMatchesFilenameMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleProfileThree.fsh",
						Want:      "ExampleProfileOne.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
				{
					RuleID:   rules.ProfileNameMatchesFilenameID,
					Message:  rules.ProfileNameMatchesFilenameMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleProfileThree.fsh",
						Want:      "ExampleProfileTwo.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
	}

	var sut rules.ProfileNameMatchesFilenameRule
	ruleName := "ProfileNameMatchesFilenameRule"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sut.Validate(&tt.fileContext)
			if err != nil {
				t.Errorf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, nil)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%v.Validate() mismatch (-got +want):\n%s", ruleName, diff)
			}
		})
	}
}
