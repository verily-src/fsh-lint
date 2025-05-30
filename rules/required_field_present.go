package rules

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/verily-src/fsh-lint/internal/fsh/types"
	"github.com/verily-src/fsh-lint/lint"
)

var (
	ErrFieldPathEmpty = fmt.Errorf("FieldPath cannot be empty")
	ErrInvalidField   = fmt.Errorf("invalid field")
	ErrInvalidValue   = fmt.Errorf("invalid value")
	ErrNotRecursable  = fmt.Errorf("not recursable")
	ErrNotSupported   = fmt.Errorf("kind not supported")

	parsedElementStringType = reflect.TypeFor[*types.ParsedElement[string]]()
)

// RequiredFieldPresentRule will check that the field given by the FieldPath value is non-nil
// and non empty. The field paths supports primitive types, pointers, slices, and arrays:
// * Primitive types (int, bool, string, etc.) are valid if they are a non-zero value (val.IsZero() is false)
// * Pointers are valid if they are not a nil pointer
// * Slices are valid if they are non-nil and not empty
// * Arrays are valid if they are non-empty
type RequiredFieldPresentRule struct {
	// FieldPath is the path to the field that should be non-nil and non-empty.
	// The path should be dot-separated. This rule will check that for each direct parent of
	// that field that exists, the field should be non-nil and non-empty.
	//
	// Example: FieldPath "Profiles.ProfileRules.BindingRules.Strength" means that
	// for each Profiles.ProfileRules.BindingRules that is found/exists, the Strength field
	// must exist. If BindingRules slice is nil, there will be no problems, because there are no
	// binding rules at all, so none are missing the strength field.
	FieldPath string

	// FieldName is the desired display name for the field in the lint message.
	FieldName string
}

// ID() returns the rule ID.
func (*RequiredFieldPresentRule) ID() string {
	return "required-field-present"
}

// Message() returns the appropriate lint error message for this rule.
func (r *RequiredFieldPresentRule) Message() string {
	if r.FieldName == "" {
		return fmt.Sprintf("%s is missing.", r.FieldPath)
	}
	return fmt.Sprintf("%s is missing.", r.FieldName)
}

// Validate returns a *lint.Problem if the value at FieldPath is missing. An error will
// be returned if the FieldPath is an invalid path.
func (r *RequiredFieldPresentRule) Validate(fc *lint.FileContext) ([]*lint.Problem, error) {
	// check that FieldPath is valid
	if r.FieldPath == "" {
		return nil, ErrFieldPathEmpty
	}

	// Split the fieldPath into components.
	pathParts := strings.Split(r.FieldPath, ".")
	fshDoc := *fc.ParsedFSH
	val := reflect.ValueOf(fshDoc)
	valid, err := checkFieldValidity(pathParts, val)
	if err != nil {
		return nil, err
	}

	if valid {
		return nil, nil
	}
	p, err := lint.NewProblem(r.ID(), r.Message(), nil, nil, false)
	if err != nil {
		return nil, err
	}
	return []*lint.Problem{p}, nil
}

// checkFieldValidity returns true if the given field in pathParts, that is, the last
// element in pathParts is not nil or empty. pathParts[0] should be the next field
// to be accessed in val. If pathParts[i] is a slice/array or pointer, then checkFieldValidity
// will check every non nil value's sub-values when i is not the last element of pathParts.
// That is, for each pathParts[:len(pathParts)-1] that exists, the pathParts[:len(pathParts)] can't be nil.
// Examples:
// If val is Profile{Name: "John"}, and pathParts is ["Name"], checkFieldValidity(pathParts, val) is true,
// since the Name field is "John", so it is not nil or empty.
func checkFieldValidity(pathParts []string, val reflect.Value) (bool, error) {
	if !val.IsValid() {
		return false, fmt.Errorf("%w.", ErrInvalidValue)
	}

	// base case: reached the end of the pathParts, so val is the value to check
	if len(pathParts) == 0 {
		return !isNilOrEmpty(val), nil
	}

	// recursive case: keep traversing the path
	switch {
	case val.Kind() == reflect.Struct:
		field := val.FieldByName(pathParts[0])
		if !field.IsValid() {
			// Field does not exist in the struct
			return false, fmt.Errorf("%w: Field '%s' is not a field of struct '%s'", ErrInvalidField, pathParts[0], val.Type())
		}

		return checkFieldValidity(pathParts[1:], field)

	case val.Kind() == reflect.Ptr:
		if !val.IsNil() {
			return checkFieldValidity(pathParts, val.Elem())
		}

		// Since we have not arrived at the desired field (not the base case),
		// this means that val does not contain any occurrences of the parent of
		// the desired field, so val is considered valid.
		return true, nil
	case val.Kind() == reflect.Slice || val.Kind() == reflect.Array:
		for i := range val.Len() {
			elem := val.Index(i)

			valid, err := checkFieldValidity(pathParts, elem)
			if err != nil {
				return false, err
			}

			// found an element that is invalid in the slice/array
			if !valid {
				return false, nil
			}
		}

		// all the elements in the slice/array are valid
		return true, nil
	case isTerminalKind(val.Kind()):
		// This case will only happen if the pathParts provided points to an
		// extraneous subfield that does not exist. When pathParts is a valid field,
		// the base case will catch the field at the right time.
		return false, fmt.Errorf(
			"%w: Field %s is a terminal kind (type=%s) and cannot be further recursed upon."+
				" Please ensure FieldPath does not have extraneous subfields.",
			ErrNotRecursable, val.String(), val.Kind())
	default:
		return false, fmt.Errorf("%w: val.Kind() of %s is not yet supported.", ErrNotSupported, val.Kind())
	}
}

// isNilOrEmpty returns true if the given val is nil or empty, and false otherwise.
// If val is specifically a ParsedElement[string], it will return true if val.Value
// is an empty string.
func isNilOrEmpty(val reflect.Value) bool {
	// Check for invalid values (e.g., reflect.ValueOf(nil))
	if !val.IsValid() {
		return true
	}

	// Ptrs and slices can't be nil
	if (val.Kind() == reflect.Ptr || val.Kind() == reflect.Slice) && val.IsNil() {
		return true
	}

	// Supported values that can be nil are ptr and slice
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Len() == 0
	}

	// Special case, if it's type *ParsedElement[string], check that the value field
	// is not set to an empty string. And we know from above that it is not a nil ptr.
	// Note: This case is for convenience for the user when they make the rules,
	// they can simply do for example:
	//    &lint.RequiredFieldPresentRule{FieldPath: "Profiles.Description"}
	// instead of doing both:
	//    &lint.RequiredFieldPresentRule{FieldPath: "Profiles.Description"} and
	//    &lint.RequiredFieldPresentRule{FieldPath: "Profiles.Description.Value"}
	if val.Type() == parsedElementStringType {
		parsedElement := val.Interface().(*types.ParsedElement[string])
		return parsedElement.Value == ""
	}

	return val.IsZero()
}

// isTerminalKind returns true if kind is a value that cannot be recursed further into,
// and false otherwise. Terminal types are all primitive types listed in the list
// of all reflect.Kind values: https://pkg.go.dev/reflect#Kind
func isTerminalKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return true
	default:
		return false
	}
}
