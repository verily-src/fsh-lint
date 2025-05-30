package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintValueSetNameMatchesFilename(t *testing.T) {
	sut := rules.ValueSetNameMatchesFilenameRule{NameSuffix: "_VS"}
	ruleName := "ValueSetNameMatchesFilenameRule"

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
			name: "no value sets",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{}},
				Path: "path/to/ValueSet.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "exact match",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSet",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleValueSet.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "correct match no path",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSet_VS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "ExampleValueSet.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "correct match, path in root",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSet_VS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "/ExampleValueSet.fsh",
			},
			want: []*lint.Problem{},
		},
		{
			name: "not match",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSetOne_VS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleValueSetTwo.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleValueSetTwo.fsh",
						Want:      "ExampleValueSetOne.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
		{
			name: "empty name",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleValueSet.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleValueSet.fsh",
						Want:      ".fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
		{
			name: "one matching, one non matching",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSetOne_VS",
								Location: testLocation,
							},
						},
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSetTwo_VS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleValueSetOne.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleValueSetOne.fsh",
						Want:      "ExampleValueSetTwo.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
		{
			name: "two non matching",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: []*types.ValueSet{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSetOne",
								Location: testLocation,
							},
						},
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleValueSetOne_VS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleValueSetTwo.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleValueSetTwo.fsh",
						Want:      "ExampleValueSetOne.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleValueSetTwo.fsh",
						Want:      "ExampleValueSetOne.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
			},
		},
	}

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
