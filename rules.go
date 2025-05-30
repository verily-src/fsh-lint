package main

import (
	"github.com/verily-src/fsh-lint/lint"
	"github.com/verily-src/fsh-lint/rules"
)

var (
	// Linter is the linter that will be used to lint the files.
	Linter *lint.Linter
)

func init() {
	requiredFieldsRules := []lint.Rule{
		&rules.RequiredFieldPresentRule{FieldPath: "Profiles.Name", FieldName: "Profile Name"},
		&rules.RequiredFieldPresentRule{FieldPath: "Profiles.ID", FieldName: "Profile ID"},
		&rules.RequiredFieldPresentRule{FieldPath: "Profiles.Title", FieldName: "Profile Title"},
		&rules.ProfileAssignmentPresentRule{Element: "status"},
		&rules.ProfileAssignmentPresentRule{
			Element:           "abstract",
			AssignmentExample: "* ^abstract = true or * ^abstract = false",
		},

		&rules.RequiredFieldPresentRule{FieldPath: "ValueSets.Name", FieldName: "Value Set Name"},
		&rules.RequiredFieldPresentRule{FieldPath: "ValueSets.ID", FieldName: "Value Set ID"},
		&rules.RequiredFieldPresentRule{FieldPath: "ValueSets.Title", FieldName: "Value Set Title"},

		&rules.RequiredFieldPresentRule{FieldPath: "CodeSystems.Name", FieldName: "Code System Name"},
		&rules.RequiredFieldPresentRule{FieldPath: "CodeSystems.ID", FieldName: "Code System ID"},
		&rules.RequiredFieldPresentRule{FieldPath: "CodeSystems.Title", FieldName: "Code System Title"},
	}
	rules := []lint.Rule{
		&rules.ProfileNameFormatRule{},
		&rules.ProfileNameMatchesFilenameRule{},
		&rules.ProfileNameMatchesIDRule{},
		&rules.ProfileNameMatchesTitleRule{},

		&rules.ValueSetNameMatchesFilenameRule{},
		&rules.ValueSetNameMatchesIDRule{},
		&rules.ValueSetNameMatchesTitleRule{},

		&rules.CodeSystemNameMatchesFilenameRule{},
		&rules.CodeSystemNameMatchesIDRule{},
		&rules.CodeSystemNameMatchesTitleRule{},
	}

	Linter = lint.NewLinter(requiredFieldsRules, rules)
}
