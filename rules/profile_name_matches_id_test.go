package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintProfileNameMatchesID(t *testing.T) {
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
			name: "name matches id",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfile", "example-profile", testLocation),
			},
			want: nil,
		},
		{
			name: "one word match",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("Example", "example", testLocation),
			},
			want: nil,
		},
		{
			name: "id different text",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfileOne", "example-profile-two", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesIDID,
					Message:  rules.ProfileNameMatchesIDMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "example-profile-two",
						Want:      "example-profile-one",
						FieldName: "Profile ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "id incorrect case",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfileName", "Example-Profile-Name", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "Example-Profile-Name",
					Want:      "example-profile-name",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
		{
			name: "id incorrect kebab style",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfileName", "exampleprofile-name", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "exampleprofile-name",
					Want:      "example-profile-name",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
		{
			name: "empty id",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfileName", "", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "",
					Want:      "example-profile-name",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
		{
			name: "one matching, one non matching profile",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfileOne", "example-profile-one", testLocation),
				createTestProfileWithNameAndID("ExampleProfileTwo", "example-profile-one", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "example-profile-one",
					Want:      "example-profile-two",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
		{
			name: "two non matching profiles",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("ExampleProfileOne", "example-profile-three", testLocation),
				createTestProfileWithNameAndID("ExampleProfileTwo", "example-profile-three", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   rules.ProfileNameMatchesIDID,
					Message:  rules.ProfileNameMatchesIDMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "example-profile-three",
						Want:      "example-profile-one",
						FieldName: "Profile ID",
					},
					IsFixable: true,
				},
				{
					RuleID:   rules.ProfileNameMatchesIDID,
					Message:  rules.ProfileNameMatchesIDMessage,
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "example-profile-three",
						Want:      "example-profile-two",
						FieldName: "Profile ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "no hyphen after number",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("123Abc", "123abc", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphen before number",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("Abc123", "abc123", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphens around numbers",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("A123B", "a123b", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphen numbers before or after",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("123Abc456", "123abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one hyphen after none before",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("123Abc456", "123-abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one hyphen before none after",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("123Abc456", "123abc-456", testLocation),
			},
			want: nil,
		},
		{
			name: "hyphens before and after",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("123Abc456", "123-abc-456", testLocation),
			},
			want: nil,
		},
		{
			name: "extra hyphens between numbers",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("1234abc", "1-2-34abc", testLocation),
			},
			want: nil,
		},
		{
			name: "double hyphen should not pass",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("12", "1--2", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "1--2",
					Want:      "12",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
		{
			name: "should not start with a hyphen",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("1Abc", "-1abc", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "-1abc",
					Want:      "1-abc",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
		{
			name: "should not end with a hyphen",
			profiles: []*types.Profile{
				createTestProfileWithNameAndID("Abc1", "Abc1-", testLocation),
			},
			want: []*lint.Problem{{
				RuleID:   rules.ProfileNameMatchesIDID,
				Message:  rules.ProfileNameMatchesIDMessage,
				Location: testLocation,
				Diff: &lint.Diff{
					Got:       "Abc1-",
					Want:      "abc-1",
					FieldName: "Profile ID",
				},
				IsFixable: true,
			}},
		},
	}

	var sut rules.ProfileNameMatchesIDRule
	ruleName := "ProfileNameMatchesIDRule"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					Profiles: tt.profiles},
			}

			got, err := sut.Validate(fc)
			if err != nil {
				t.Fatalf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, nil)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%v.Validate() mismatch (-got +want):\n%s", ruleName, diff)
			}

		})
	}
}

func createTestProfileWithNameAndID(name string, id string, testLocation *types.Location) *types.Profile {
	return &types.Profile{
		Name: &types.ParsedElement[string]{
			Value:    name,
			Location: testLocation,
		},
		ID: &types.ParsedElement[string]{
			Value:    id,
			Location: testLocation,
		},
	}
}
