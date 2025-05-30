package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintCodeSystemNameMatchesFilename(t *testing.T) {
	sut := rules.CodeSystemNameMatchesFilenameRule{NameSuffix: "_CS"}
	ruleName := "CodeSystemNameMatchesFilenameRule"

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
			name: "no code systems",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					CodeSystems: []*types.CodeSystem{}},
				Path: "path/to/CodeSystem.fsh",
			},
			want: nil,
		},
		{
			name: "exact match",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystem",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleCodeSystem.fsh",
			},
			want: nil,
		},
		{
			name: "correct match no path",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystem_CS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "ExampleCodeSystem.fsh",
			},
			want: nil,
		},
		{
			name: "correct match, path in root",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystem_CS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "/ExampleCodeSystem.fsh",
			},
			want: nil,
		},
		{
			name: "not match",
			fileContext: lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystemOne_CS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleCodeSystemTwo.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleCodeSystemTwo.fsh",
						Want:      "ExampleCodeSystemOne.fsh",
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
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleCodeSystem.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleCodeSystem.fsh",
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
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystemOne_CS",
								Location: testLocation,
							},
						},
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystemTwo_CS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleCodeSystemOne.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleCodeSystemOne.fsh",
						Want:      "ExampleCodeSystemTwo.fsh",
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
					CodeSystems: []*types.CodeSystem{
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystemOne",
								Location: testLocation,
							},
						},
						{
							Name: &types.ParsedElement[string]{
								Value:    "ExampleCodeSystemOne_CS",
								Location: testLocation,
							},
						},
					},
				},
				Path: "path/to/ExampleCodeSystemTwo.fsh",
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleCodeSystemTwo.fsh",
						Want:      "ExampleCodeSystemOne.fsh",
						FieldName: "Filename",
					},
					IsFixable: false,
				},
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "ExampleCodeSystemTwo.fsh",
						Want:      "ExampleCodeSystemOne.fsh",
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
				t.Fatalf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, nil)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("%v.Validate() mismatch (-got +want):\n%s", ruleName, diff)
			}
		})
	}
}
