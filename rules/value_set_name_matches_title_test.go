package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintValueSetNameMatchesTitle(t *testing.T) {
	sut := rules.ValueSetNameMatchesTitleRule{NameSuffix: "_VS"}
	ruleName := "ValueSetNameMatchesTitleRule"

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
		name      string
		valueSets []*types.ValueSet
		want      []*lint.Problem
	}{
		{
			name:      "no value sets",
			valueSets: []*types.ValueSet{},
			want:      nil,
		},
		{
			name: "name matches title with _VS",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleTest_VS", "Example Test", testLocation),
			},
			want: nil,
		},
		{
			name: "name matches title without _VS",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleTest", "Example Test", testLocation),
			},
			want: nil,
		},
		{
			name: "name matches title with acronym",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleTestABC_VS", "Example Test ABC", nil),
			},
			want: nil,
		},
		{
			name: "title is the same but adds '_VS' at the end",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleName_VS", "Example Name VS", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Name VS",
						Want:      "Example Name",
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "incorrect title spacing",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleName", "Example   Name", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example   Name",
						Want:      "Example Name",
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "empty title",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleName", "", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "",
						Want:      "Example Name",
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "one matching, one non matching",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleOne", "Example Two", testLocation),
				createTestValueSetWithNameAndTitle("ExampleTwo", "Example Two", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Two",
						Want:      "Example One",
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "two non matching",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("ExampleOne", "Example Three", testLocation),
				createTestValueSetWithNameAndTitle("ExampleTwo", "Example Three", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Three",
						Want:      "Example One",
						FieldName: "Value Set Title",
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
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "no space after number",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("123Abc", "123Abc", testLocation),
			},
			want: nil,
		},
		{
			name: "no space before number",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("Abc123", "Abc123", testLocation),
			},
			want: nil,
		},
		{
			name: "no spaces around numbers",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("A123B", "A123b", testLocation),
			},
			want: nil,
		},
		{
			name: "no space numbers before or after",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("123Abc456", "123abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one space after none before",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("123Abc456", "123 abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one space before none after",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("123Abc456", "123abc 456", testLocation),
			},
			want: nil,
		},
		{
			name: "spaces before and after",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("123Abc456", "123 ABC 456", testLocation),
			},
			want: nil,
		},
		{
			name: "extra spaces between numbers",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("1234abc", "1 2 34ABC", testLocation),
			},
			want: nil,
		},
		{
			name: "double space should not pass",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("12", "1  2", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "1  2",
						Want:      "12",
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not start with a space",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("1Abc", " 1abc", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       " 1abc",
						Want:      "1 Abc",
						FieldName: "Value Set Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not end with a space",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndTitle("Abc1", "Abc1 ", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Abc1 ",
						Want:      "Abc 1",
						FieldName: "Value Set Title",
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
					ValueSets: tt.valueSets,
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

func createTestValueSetWithNameAndTitle(name string, title string, testLocation *types.Location) *types.ValueSet {
	return &types.ValueSet{
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
