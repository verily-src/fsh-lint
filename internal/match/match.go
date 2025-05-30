package match

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

// IsNameMatchWithTitle returns true if a PascalCase name matches a Title Case title,
// and false otherwise. If relaxNums is true, spaces in the Title are optional when
// adjacent to a number. Otherwise, the PascalCase name converted to Title Case must
// match the Title exactly.

// Examples:
//
//	IsNameMatchWithTitle("OneTwo3", "One Two3", false) returns false because relaxNums is
//		false and the name converted to Title Case would be "One Two 3", not "One Two3".
//	IsNameMatchWithTitle("OneTwo3", "One Two3", true) returns true because relaxNums
//		is true so the space adjacent to "3" in the title is optional, so "One Two3" is a valid title.
func IsNameMatchWithTitle(name string, title string, relaxNums bool) bool {
	spaceSeparatedName := strcase.ToDelimited(name, ' ')
	loweredTitle := strings.ToLower(title)

	// exact match
	if loweredTitle == spaceSeparatedName {
		return true
	}

	// title is blank
	if loweredTitle == "" {
		return false
	}

	// non-exact match and we are not relaxing numbers
	if !relaxNums {
		return false
	}

	if isAlphaNumericWords(title) && isNumRelaxedMatch(loweredTitle, spaceSeparatedName, " ") {
		return true
	}

	return false
}

// isNameKebabMatchWithID returns true if a PascalCase name matches a kebab-case ID,
// and false otherwise. If relaxNums is true, hyphens in the ID are optional when
// adjacent to numbers. Otherwise, the PascalCase name converted to kebab-case must
// match the ID exactly.

// Examples:
//
//	isNameKebabMatchWithID("OneTwo3", "one-two3", false) returns false because relaxNums is
//		false and the name converted to kebab-case would be "one-two-3", not "one-two3".
//	IsNameMatchWithTitle("OneTwo3", "one-two3", true) returns true because relaxNums
//		is true so the hyphen adjacent to "3" is optional and "one-two3" becomes a valid title.

func IsNameKebabMatchWithID(name string, id string, relaxNums bool) bool {
	kebabedName := strcase.ToKebab(name)

	// exact match
	if id == kebabedName {
		return true
	}

	// id is blank
	if id == "" {
		return false
	}

	// non-exact match and we are not relaxing numbers
	if !relaxNums {
		return false
	}

	if isKebabCase(id) && isNumRelaxedMatch(id, kebabedName, "-") {
		return true
	}

	return false
}

// isAlphaNumericWords returns true if s is a valid string of alphanumeric words
// with no special characters or incorrect spacing, and false otherwise.
func isAlphaNumericWords(s string) bool {
	// empty string is valid
	if s == "" {
		return true
	}

	// s starts or ends with a space
	if s[0] == ' ' || s[len(s)-1] == ' ' {
		return false
	}

	// s has double spaces
	if strings.Contains(s, "  ") {
		return false
	}

	// s contains characters that are not a-z, A-Z, 0-9, or space
	alphaNumChars := regexp.MustCompile("^[a-zA-Z0-9 ]*$")
	if !alphaNumChars.MatchString(s) {
		return false
	}

	return true
}

// isKebabCase checks if s is in kebab-case.
func isKebabCase(s string) bool {
	// empty string is kebab-case
	if s == "" {
		return true
	}

	// s starts or ends with a hyphen
	if s[0] == '-' || s[len(s)-1] == '-' {
		return false
	}

	// s has double hyphens
	if strings.Contains(s, "--") {
		return false
	}

	// s contains characters that are not a-z, 0-9, or '-'
	kebabChars := regexp.MustCompile("^[a-z0-9-]*$")
	if !kebabChars.MatchString(s) {
		return false
	}

	return true
}

// isNumRelaxedMatch returns true if a and b are matching strings where a delimiter is
// optional if it are found adjacent to a number character, and false otherwise.
func isNumRelaxedMatch(a string, b string, delim string) bool {
	delimAfterNum := regexp.MustCompile("([0-9])" + delim)
	delimBeforeNum := regexp.MustCompile(delim + "([0-9])")

	a = delimAfterNum.ReplaceAllString(a, `$1`)
	a = delimBeforeNum.ReplaceAllString(a, `$1`)

	b = delimAfterNum.ReplaceAllString(b, `$1`)
	b = delimBeforeNum.ReplaceAllString(b, `$1`)

	return a == b
}
