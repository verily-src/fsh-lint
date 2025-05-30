package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintProfileNameMatchesTitle(t *testing.T) {
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
		name     string
		profiles []*types.Profile
		want     []*lint.Problem
	}{
		{
			name:     "no profiles",
			profiles: []*types.Profile{},
			want:     nil,
		},
		{
			name: "name matches title",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleTest", "Example Test Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "name matches title with acronym",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleTestABC", "Example Test ABC Profile", nil),
			},
			want: nil,
		},
		{
			name: "if profile is already in the name, don't need another 'Profile' suffix",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleProfile", "Example Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "title is the same but missing 'Profile' at the end",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleName", "Example Name", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Name",
						Want:      "Example Name Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "incorrect title spacing",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleName", "Example   Name Profile", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example   Name Profile",
						Want:      "Example Name Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "empty title",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleName", "", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "",
						Want:      "Example Name Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "one matching, one non matching profile",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleOne", "Example Two Profile", testLocation),
				createTestProfileWithNameAndTitle("ExampleTwo", "Example Two Profile", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Two Profile",
						Want:      "Example One Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "two non matching profiles",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ExampleOne", "Example Three", testLocation),
				createTestProfileWithNameAndTitle("ExampleTwo", "Example Three", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Three",
						Want:      "Example One Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example Three",
						Want:      "Example Two Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "different case ignored",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("ProfileName", "Profile nAME Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "no space after number",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("123Abc", "123Abc Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "no space before number",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("Abc123", "Abc123 Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "no spaces around numbers",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("A123B", "A123b Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "no space numbers before or after",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("123Abc456", "123abc456 Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "one space after none before",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("123Abc456", "123 abc456 Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "one space before none after",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("123Abc456", "123abc 456 Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "spaces before and after",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("123Abc456", "123 ABC 456 Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "extra spaces between numbers",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("1234abc", "1 2 34ABC Profile", testLocation),
			},
			want: nil,
		},
		{
			name: "double space should not pass",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("12", "1  2 Profile", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "1  2 Profile",
						Want:      "12 Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not start with a space",
			profiles: []*types.Profile{
				createTestProfileWithNameAndTitle("1Abc", " 1abc Profile", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesTitleID,
					Message:  rules.ProfileNameMatchesTitleMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       " 1abc Profile",
						Want:      "1 Abc Profile",
						FieldName: "Profile Title",
					},
					IsFixable: true,
				},
			},
		},
	}

	var sut rules.ProfileNameMatchesTitleRule
	ruleName := "ProfileNameMatchesTitleRule"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileContext := lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: tt.profiles,
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

func createTestProfileWithNameAndTitle(name string, title string, testLocation *types.Location) *types.Profile {
	return &types.Profile{
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
