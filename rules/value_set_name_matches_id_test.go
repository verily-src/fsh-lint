package rules_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

func TestLintValueSetNameMatchesID(t *testing.T) {
	sut := rules.ValueSetNameMatchesIDRule{NameSuffix: "_VS"}
	ruleName := "ValueSetNameMatchesIDRule"

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
			name: "name with _VS matches id",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ValueSet_VS", "value-set", testLocation),
			},
			want: nil,
		},
		{
			name: "name without _VS matches id",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ValueSet", "value-set", testLocation),
			},
			want: nil,
		},
		{
			name: "one word match",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("Example", "example", testLocation),
			},
			want: nil,
		},
		{
			name: "id different text",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ValueSetOne", "value-set-two", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "value-set-two",
						Want:      "value-set-one",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "id incorrect case",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ExampleValueSet", "Example-Value-Set", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example-Value-Set",
						Want:      "example-value-set",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "id incorrect kebab style",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ExampleValueSet", "examplevalue-set", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "examplevalue-set",
						Want:      "example-value-set",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "empty id",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ExampleValueSet_VS", "", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "",
						Want:      "example-value-set",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "one matching, one non matching",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ValueSetOne_VS", "value-set-one", testLocation),
				createTestValueSetWithNameAndID("ValueSetOne_VS", "value-set-one-vs", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "value-set-one-vs",
						Want:      "value-set-one",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "two non matching",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("ValueSetOne_VS", "value-set-three", testLocation),
				createTestValueSetWithNameAndID("ValueSetTwo_VS", "value-set-three", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "value-set-three",
						Want:      "value-set-one",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "value-set-three",
						Want:      "value-set-two",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "no hyphen after number",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("123Abc", "123abc", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphen before number",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("Abc123", "abc123", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphens around numbers",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("A123B", "a123b", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphen numbers before or after",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("123Abc456", "123abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one hyphen after none before",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("123Abc456", "123-abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one hyphen before none after",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("123Abc456", "123abc-456", testLocation),
			},
			want: nil,
		},
		{
			name: "hyphens before and after",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("123Abc456", "123-abc-456", testLocation),
			},
			want: nil,
		},
		{
			name: "extra hyphens between numbers",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("1234abc", "1-2-34abc", testLocation),
			},
			want: nil,
		},
		{
			name: "double hyphen should not pass",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("12", "1--2", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "1--2",
						Want:      "12",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not start with a hyphen",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("1Abc", "-1abc", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "-1abc",
						Want:      "1-abc",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not end with a hyphen",
			valueSets: []*types.ValueSet{
				createTestValueSetWithNameAndID("Abc1", "Abc1-", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Abc1-",
						Want:      "abc-1",
						FieldName: "Value Set ID",
					},
					IsFixable: true,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &lint.FileContext{
				ParsedFSH: &types.FSHDocument{
					ValueSets: tt.valueSets},
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

func createTestValueSetWithNameAndID(name string, id string, testLocation *types.Location) *types.ValueSet {
	return &types.ValueSet{
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
