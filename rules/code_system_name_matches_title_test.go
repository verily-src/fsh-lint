package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintCodeSystemNameMatchesTitle(t *testing.T) {
	sut := rules.CodeSystemNameMatchesTitleRule{NameSuffix: "_CS"}
	ruleName := "CodeSystemNameMatchesTitleRule"

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
		codeSystems []*types.CodeSystem
		want        []*lint.Problem
	}{
		{
			name:        "no code systems",
			codeSystems: []*types.CodeSystem{},
			want:        nil,
		},
		{
			name: "name matches title with _CS",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleTest_CS", "Example Test", testLocation),
			},
			want: nil,
		},
		{
			name: "name matches title without _CS",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleTest", "Example Test", testLocation),
			},
			want: nil,
		},
		{
			name: "name matches title with acronym",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleTestABC_CS", "Example Test ABC", nil),
			},
			want: nil,
		},
		{
			name: "title is the same but adds '_CS' at the end",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleName_CS", "Example Name CS", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Name CS",
						Want:      "Example Name",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "incorrect title spacing",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleName", "Example   Name", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example   Name",
						Want:      "Example Name",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "empty title",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleName", "", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "",
						Want:      "Example Name",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "one matching, one non matching",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleOne", "Example Two", testLocation),
				createTestCodeSystemWithNameAndTitle("ExampleTwo", "Example Two", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Two",
						Want:      "Example One",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "two non matching",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("ExampleOne", "Example Three", testLocation),
				createTestCodeSystemWithNameAndTitle("ExampleTwo", "Example Three", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Three",
						Want:      "Example One",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Three",
						Want:      "Example Two",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "different case ignored",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("CodeSystemName", "Code system nAME", testLocation),
			},
			want: nil,
		},
		{
			name: "no space after number",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("123Abc", "123Abc", testLocation),
			},
			want: nil,
		},
		{
			name: "no space before number",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("Abc123", "Abc123", testLocation),
			},
			want: nil,
		},
		{
			name: "no spaces around numbers",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("A123B", "A123b", testLocation),
			},
			want: nil,
		},
		{
			name: "no space numbers before or after",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("123Abc456", "123abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one space after none before",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("123Abc456", "123 abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one space before none after",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("123Abc456", "123abc 456", testLocation),
			},
			want: nil,
		},
		{
			name: "spaces before and after",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("123Abc456", "123 ABC 456", testLocation),
			},
			want: nil,
		},
		{
			name: "extra spaces between numbers",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("1234abc", "1 2 34ABC", testLocation),
			},
			want: nil,
		},
		{
			name: "double space should not pass",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("12", "1  2", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "1  2",
						Want:      "12",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not start with a space",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("1Abc", " 1abc", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       " 1abc",
						Want:      "1 Abc",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not end with a space",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndTitle("Abc1", "Abc1 ", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Abc1 ",
						Want:      "Abc 1",
						FieldName: "Code System Title",
					},
					IsFixable: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileContext := lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					CodeSystems: tt.codeSystems,
				},
			}
			got, err := sut.Validate(&fileContext)
			if err != nil {
				t.Fatalf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, nil)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%v.Validate() mismatch (-got +want):\n%s", ruleName, diff)
			}
		})
	}
}

func createTestCodeSystemWithNameAndTitle(name string, title string, testLocation *types.Location) *types.CodeSystem {
	return &types.CodeSystem{
		Name: &types.ParsedElement[string]{
			Value:    name,
			Location: testLocation,
		},
		Title: &types.ParsedElement[string]{
			Value:    title,
			Location: testLocation,
		},
	}
}
