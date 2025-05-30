package rules_test

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintProfileNameFormat(t *testing.T) {
	regexpString := "^Prefix.*Suffix$"
	profileRegex := regexp.MustCompile(regexpString)
	sut := rules.ProfileNameFormatRule{RegexFormat: profileRegex}
	ruleName := "ProfileNameFormatRule"

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
		fshDocument *types.FSHDocument
		want        []*lint.Problem
	}{
		{
			name:        "no profiles",
			fshDocument: &types.FSHDocument{},
			want:        []*lint.Problem{},
		},
		{
			name: "multiple valid profiles",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value:    "PrefixOneSuffix",
							Location: testLocation,
						},
					},
					{
						Name: &types.ParsedElement[string]{
							Value:    "PrefixTwoSuffix",
							Location: testLocation,
						},
					},
				},
			},
			want: []*lint.Problem{},
		},
		{
			name: "multiple invalid profiles",
			fshDocument: &types.FSHDocument{
				Profiles: []*types.Profile{
					{
						Name: &types.ParsedElement[string]{
							Value:    "NotPrefixOneSuffix",
							Location: testLocation,
						},
					},
					{
						Name: &types.ParsedElement[string]{
							Value:    "PrefixTwoSuffixNot",
							Location: testLocation,
						},
					},
				},
			},
			want: []*lint.Problem{
				{
					RuleID:    sut.ID(),
					Message:   sut.Message(),
					Location:  testLocation,
					IsFixable: false,
				},
				{
					RuleID:    sut.ID(),
					Message:   sut.Message(),
					Location:  testLocation,
					IsFixable: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileContext := lint.FileContext{
				ParsedFSH: tt.fshDocument,
			}
			got, err := sut.Validate(&fileContext)
			if err != nil {
				t.Errorf("%v.Validate(): got error = %v, want error = %v.", ruleName, err, nil)
			}
			if diff := cmp.Diff(got, tt.want, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("%v.Validate() mismatch (-got +want):\n%s", ruleName, diff)
			}
		})
	}
}
