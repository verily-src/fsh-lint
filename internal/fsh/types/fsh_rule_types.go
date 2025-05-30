package types

import "fmt"

// CardRule represents a FSH cardinality rule.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#cardinality-rules for details.
type CardRule struct {
	Element     *ParsedElement[string]
	Cardinality *Cardinality
	Flags       *Flags
}

func (cr *CardRule) String() string {
	return fmt.Sprintf("CardRule{\n  Element: %v,\n  Cardinality: %v,\n  Flags: %v\n}", cr.Element, cr.Cardinality, cr.Flags)
}

// FlagRule represents a FSH flag rule.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#flag-rules for details.
type FlagRule struct {
	Elements []*ParsedElement[string]
	Flags    *Flags
}

func (fr *FlagRule) String() string {
	return fmt.Sprintf("FlagRule{\n  Elements: %v,\n  Flags: %v\n}", fr.Elements, fr.Flags)
}

// BindingRule represents a FSH binding rule. Strength is either nil or one
// of the strengths listed here: https://hl7.org/fhir/R5/valueset-binding-strength.html
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#binding-rules for details.
type BindingRule struct {
	Bindable *ParsedElement[string] `json:"bindable"`
	ValueSet *ParsedElement[string] `json:"valueSet"`
	Strength *ParsedElement[string] `json:"strength"`
}

func (br *BindingRule) String() string {
	return fmt.Sprintf("BindingRule{\n  Bindable: %v,\n  ValueSet: %v,\n  Strength: %v\n}", br.Bindable, br.ValueSet, br.Strength)
}

// AssignmentRule represents a FSH assignment rule. For assignment values that have multiple strings, such as
// assignments with quantity data, or with the coding data type, Value will be a concatenation of all the given strings.
// This is a feature of the grammar, since the lexer will not return each string individually.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules for details.
type AssignmentRule struct {
	Element *ParsedElement[string] `json:"element"`
	Value   *ParsedElement[string] `json:"value"`
	Exactly *ParsedElement[bool]   `json:"exactly"`
}

func (ar *AssignmentRule) String() string {
	return fmt.Sprintf("AssignmentRule{\n  Element: %v,\n  Value: %v,\n  Exactly: %v\n}", ar.Element, ar.Value, ar.Exactly)
}

// ContainsRule represents a FSH contains rule.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#contains-rules-for-extensions for details.
type ContainsRule struct {
	Name  *ParsedElement[string] `json:"name"`
	Items []*Item                `json:"items"`
}

func (cr *ContainsRule) String() string {
	return fmt.Sprintf("ContainsRule{\n  Name: %v,\n  Items: %v\n}", cr.Name, cr.Items)
}

// TypeRule represents a FSH type rule.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#type-rules for details.
type TypeRule struct {
	Element *ParsedElement[string] `json:"element"`
	Types   []*DataType            `json:"types"`
}

func (tr *TypeRule) String() string {
	return fmt.Sprintf("TypeRule{\n  Element: %v,\n  Types: %v\n}", tr.Element, tr.Types)
}

// ObeysRule represents a FSH obeys rule where invariant(s) are applied to the whole
// profile or one element of the profile. A nil element field indicates that it applies
// to the whole profile.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#obeys-rules for details.
type ObeysRule struct {
	Element    *ParsedElement[string]   `json:"element"`
	Invariants []*ParsedElement[string] `json:"invariants"`
}

func (or *ObeysRule) String() string {
	return fmt.Sprintf("ObeysRule{\n  Element: %v,\n  Invariants: %v\n}", or.Element, or.Invariants)
}

// CaretValueRule represents a FSH assignment using caret values.
// Element is either:
//  1. an element of the StructureDefinition, (in which case ElementInProfile is nil) or
//  2. an element of an ElementDefinition within a Profile (in which case ElementInProfile is the element within the StructureDefinition)
//
// Note that Value will be a concatenation of all the given strings. This is a feature of the grammar, since the lexer will not return each string individually.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#caret-paths for details
type CaretValueRule struct {
	Element          *ParsedElement[string] `json:"element"`
	ElementInProfile *ParsedElement[string] `json:"elementInProfile"`
	Value            *ParsedElement[string] `json:"value"`
}

func (cvr *CaretValueRule) String() string {
	return fmt.Sprintf("CaretValueRule{\n  Element: %v,\n  ElementInProfile: %v,\n  Value: %v\n}", cvr.Element, cvr.ElementInProfile, cvr.Value)
}

// CodeCaretValueRule represents a FSH assignment using caret values and codes within a value set.
// ConceptCodes is a hierarchy of codes, where each code is the ancestor of all following codes.
// ConcptCode is always included or excluded before a caret rule assignments, and Element is an element
// of the corresponding concept.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules:~:text=repeat%20the%20code%20and%20follow%20it%20with%20a%20caret%20path%20assignment%20instead for details
type CodeCaretValueRule struct {
	ConceptCodes []*ParsedElement[string] `json:"conceptCodes"`
	Element      *ParsedElement[string]   `json:"element"`
	Value        *ParsedElement[string]   `json:"value"`
}

func (ccvr *CodeCaretValueRule) String() string {
	return fmt.Sprintf("CodeCaretValueRule{\n  ConceptCodes: %v,\n  Element: %v,\n  Value: %v\n}", ccvr.ConceptCodes, ccvr.Element, ccvr.Value)
}

// InsertRule represents a FSH insert rule. Path is a path to an element that the rule set will apply to.
// A nil path indicates that the rule set will apply to the whole context that it is found in.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#inserting-rule-sets-with-path-context for details.
type InsertRule struct {
	Path        *ParsedElement[string]   `json:"path"`
	RuleSetName *ParsedElement[string]   `json:"ruleSetName"`
	Parameters  []*ParsedElement[string] `json:"parameters"`
}

func (ir *InsertRule) String() string {
	return fmt.Sprintf("InsertRule{\n  Path: %v,\n  RuleSetName: %v,\n  Parameters: %v\n}", ir.Path, ir.RuleSetName, ir.Parameters)
}

// CodeInsertRule represents a FSH insert rule with a concept code as the context. This format can be found in
// code systems and value sets, where the ConceptCodes is a hierarchy of concepts. That is, there are multiple codes
// when the context is a child concept code.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#inserting-rule-sets-with-path-context:~:text=inserted%20in%20the%20context%20of%20a%20concept for details.
type CodeInsertRule struct {
	ConceptCodes []*ParsedElement[string] `json:"conceptCodes"`
	RuleSetName  *ParsedElement[string]   `json:"ruleSetName"`
	Parameters   []*ParsedElement[string] `json:"parameters"`
}

func (cir *CodeInsertRule) String() string {
	return fmt.Sprintf("CodeInsertRule{\n  ConceptCodes: %v,\n  RuleSetName: %v,\n  Parameters: %v\n}", cir.ConceptCodes, cir.RuleSetName, cir.Parameters)
}

// PathRule represents a FSH path rule.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules for details.
type PathRule struct {
	Path *ParsedElement[string] `json:"path"`
}

func (pr *PathRule) String() string {
	return fmt.Sprintf("PathRule{\n  Path: %v\n}", pr.Path)
}

// DataType represents a FSH data type. Only one of "name", "referenceType", "canonical", or
// "codeableReferenceType" will be set depending on the type that is parsed. The other
// fields will be nil. When the format is not name, the string will be in brackets. For example:
// "Canonical({Resource/Profile1} or {Resource/Profile2} or {Resource/Profile3}...)"
type DataType struct {
	Name                  *ParsedElement[string] `json:"name"`
	ReferenceType         *ParsedElement[string] `json:"referenceType"`
	Canonical             *ParsedElement[string] `json:"canonical"`
	CodeableReferenceType *ParsedElement[string] `json:"codeableReferenceType"`
}

func (dt *DataType) String() string {
	return fmt.Sprintf("DataType{\n  Name: %v,\n  ReferenceType: %v,\n  Canonical: %v,\n  CodeableReferenceType: %v\n}", dt.Name, dt.ReferenceType, dt.Canonical, dt.CodeableReferenceType)
}

// Cardinality represents a FSH cardinality.
type Cardinality struct {
	Min *ParsedElement[int] `json:"min"`
	Max *ParsedElement[int] `json:"max"`
}

func (c *Cardinality) String() string {
	return fmt.Sprintf("Cardinality{\n  Min: %v,\n  Max: %v\n}", c.Min, c.Max)
}

// Flags represent FSH flags.
type Flags struct {
	MustSupport      *ParsedElement[bool] `json:"mustSupport"`      // MS
	IncludeInSummary *ParsedElement[bool] `json:"includeInSummary"` // SU
	Modifier         *ParsedElement[bool] `json:"modifier"`         // ?!
	Normative        *ParsedElement[bool] `json:"normative"`        // N
	TrialUse         *ParsedElement[bool] `json:"trialUse"`         // TU
	Draft            *ParsedElement[bool] `json:"draft"`            // D
}

// Constructor for Flags initializes all flags to false with nil location
func NewFlags() *Flags {
	return &Flags{
		MustSupport:      NewParsedElementWithoutLocation(false),
		IncludeInSummary: NewParsedElementWithoutLocation(false),
		Modifier:         NewParsedElementWithoutLocation(false),
		Normative:        NewParsedElementWithoutLocation(false),
		TrialUse:         NewParsedElementWithoutLocation(false),
		Draft:            NewParsedElementWithoutLocation(false),
	}
}

func (f *Flags) String() string {
	return fmt.Sprintf("Flags{\n  MustSupport: %v,\n  IncludeInSummary: %v,\n  Modifier: %v,\n  Normative: %v,\n  TrialUse: %v,\n  Draft: %v\n}", f.MustSupport, f.IncludeInSummary, f.Modifier, f.Normative, f.TrialUse, f.Draft)
}

// Item represents a FSH item.
type Item struct {
	Name        *ParsedElement[string] `json:"name"`
	LocalName   *ParsedElement[string] `json:"localName"`
	Cardinality *Cardinality           `json:"cardinality"`
	Flags       *Flags                 `json:"flags"`
}

func (i *Item) String() string {
	return fmt.Sprintf("Item{\n  Name: %v,\n  LocalName: %v,\n  Cardinality: %v,\n  Flags: %v\n}", i.Name, i.LocalName, i.Cardinality, i.Flags)
}
