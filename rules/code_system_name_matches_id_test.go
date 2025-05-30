package rules_test

import (
	"testing"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestLintCodeSystemNameMatchesID(t *testing.T) {
	sut := rules.CodeSystemNameMatchesIDRule{NameSuffix: "_CS"}
	ruleName := "CodeSystemNameMatchesIDRule"

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
			codeSystems: nil,
			want:        nil,
		},
		{
			name: "name with _CS matches id",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("CodeSystem_CS", "code-system", testLocation),
			},
			want: nil,
		},
		{
			name: "name without _CS matches id",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("CodeSystem", "code-system", testLocation),
			},
			want: nil,
		},
		{
			name: "one word match",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("Example", "example", testLocation),
			},
			want: nil,
		},
		{
			name: "id different text",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("CodeSystemOne", "code-system-two", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "code-system-two",
						Want:      "code-system-one",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "id incorrect case",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("ExampleCodeSystem", "Example-Code-System", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Example-Code-System",
						Want:      "example-code-system",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "id incorrect kebab style",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("ExampleCodeSystem", "examplecode-system", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "examplecode-system",
						Want:      "example-code-system",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "empty id",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("ExampleCodeSystem_CS", "", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "",
						Want:      "example-code-system",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "one matching, one non matching",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("CodeSystemOne_CS", "code-system-one", testLocation),
				createTestCodeSystemWithNameAndID("CodeSystemOne_CS", "code-system-one-cs", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "code-system-one-cs",
						Want:      "code-system-one",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "two non matching",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("CodeSystemOne_CS", "code-system-three", testLocation),
				createTestCodeSystemWithNameAndID("CodeSystemTwo_CS", "code-system-three", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "code-system-three",
						Want:      "code-system-one",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "code-system-three",
						Want:      "code-system-two",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "no hyphen after number",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("123Abc", "123abc", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphen before number",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("Abc123", "abc123", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphens around numbers",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("A123B", "a123b", testLocation),
			},
			want: nil,
		},
		{
			name: "no hyphen numbers before or after",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("123Abc456", "123abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one hyphen after none before",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("123Abc456", "123-abc456", testLocation),
			},
			want: nil,
		},
		{
			name: "one hyphen before none after",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("123Abc456", "123abc-456", testLocation),
			},
			want: nil,
		},
		{
			name: "hyphens before and after",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("123Abc456", "123-abc-456", testLocation),
			},
			want: nil,
		},
		{
			name: "extra hyphens between numbers",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("1234abc", "1-2-34abc", testLocation),
			},
			want: nil,
		},
		{
			name: "double hyphen should not pass",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("12", "1--2", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "1--2",
						Want:      "12",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not start with a hyphen",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("1Abc", "-1abc", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "-1abc",
						Want:      "1-abc",
						FieldName: "Code System ID",
					},
					IsFixable: true,
				},
			},
		},
		{
			name: "should not end with a hyphen",
			codeSystems: []*types.CodeSystem{
				createTestCodeSystemWithNameAndID("Abc1", "Abc1-", testLocation),
			},
			want: []*lint.Problem{
				{
					RuleID:   sut.ID(),
					Message:  sut.Message(),
					Location: testLocation,
					Diff: &lint.Diff{
						Got:       "Abc1-",
						Want:      "abc-1",
						FieldName: "Code System ID",
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
					CodeSystems: tt.codeSystems},
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

func createTestCodeSystemWithNameAndID(name string, id string, testLocation *types.Location) *types.CodeSystem {
	return &types.CodeSystem{
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
