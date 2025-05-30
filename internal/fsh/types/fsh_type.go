package types

import "fmt"

// FSHDocument represents a FSH document.
// A FSH document can have any number and type of entries.
// Only implemented entries are added here for now.
type FSHDocument struct {
	ValueSets   []*ValueSet   `json:"valueSets"`
	Profiles    []*Profile    `json:"profiles"`
	CodeSystems []*CodeSystem `json:"codeSystems"`
	Instances   []*Instance   `json:"instances"`
	Extensions  []*Extension  `json:"extensions"`
}

func (doc *FSHDocument) String() string {
	return fmt.Sprintf(
		"FSHDocument{\n  ValueSets: %v,\n  Profiles: %v,\n  CodeSystems: %v\n}",
		doc.ValueSets, doc.Profiles, doc.CodeSystems,
	)
}

// ValueSet represents a FSH ValueSet. This is a custom type, and not defined in the FSH grammar.
type ValueSet struct {
	Name              *ParsedElement[string] `json:"name"`
	ID                *ParsedElement[string] `json:"id"`
	Title             *ParsedElement[string] `json:"title"`
	Description       *ParsedElement[string] `json:"description"`
	IncludeComponents []*ValueSetComponent   `json:"includeComponents"`
	ExcludeComponents []*ValueSetComponent   `json:"excludeComponents"`
	ValueSetRules     *ValueSetRules         `json:"valueSetRules"`
}

func (vs *ValueSet) String() string {
	return fmt.Sprintf(
		"ValueSet{\n  Name: %v,\n  ID: %v,\n  Title: %v,\n  Description: %v,\n  IncludeComponents: %v,\n  ExcludeComponents: %v,\n  ValueSetRules: %v\n}",
		vs.Name, vs.ID, vs.Title, vs.Description, vs.IncludeComponents, vs.ExcludeComponents, vs.ValueSetRules,
	)
}

// ValueSetRules represents an exhaustive list of the rules of a FSH ValueSet.
type ValueSetRules struct {
	CaretValueRules     []*CaretValueRule     `json:"caretValueRules"`
	CodeCaretValueRules []*CodeCaretValueRule `json:"codeCaretValueRules"`
	InsertRules         []*InsertRule         `json:"insertRules"`
	CodeInsertRules     []*CodeInsertRule     `json:"codeInsertRules"`
}

func (vsr *ValueSetRules) String() string {
	return fmt.Sprintf(
		"ValueSetRules{\n  CaretValueRules: %v,\n  CodeCaretValueRules: %v,\n  InsertRules: %v,\n  CodeInsertRules: %v\n}",
		vsr.CaretValueRules, vsr.CodeCaretValueRules, vsr.InsertRules, vsr.CodeInsertRules,
	)
}

// Each component maps to one include statement
// include means the intersection of value sets and the optional code system
type ValueSetComponent struct {
	// CodePath and CodeString will be non-nil when including a single code, but
	// CodePath and CodeString will be nil when when using include from code system or value set
	CodePath   *ParsedElement[string] `json:"codePath"`
	CodeString *ParsedElement[string] `json:"codeString"`

	// there can be at most one code system per include statement
	FromCodeSystem *ValueSetCodesSource `json:"fromCodeSystem"`

	// but there can be multiple value sets
	FromValueSet []*ValueSetCodesSource `json:"fromValueSet"`

	Filters []*ValueSetFilter `json:"filters"`
}

func (vsc *ValueSetComponent) String() string {
	return fmt.Sprintf(
		"ValueSetComponent{\n  CodePath: %v,\n  CodeString: %v,\n  FromCodeSystem: %v,\n  FromValueSet: %v,\n  Filters: %v\n}",
		vsc.CodePath, vsc.CodeString, vsc.FromCodeSystem, vsc.FromValueSet, vsc.Filters,
	)
}

// ValueSetCodesSource is either a ValueSet or a CodeSystem listed as a source in a ValueSetComponent.
type ValueSetCodesSource struct {
	Name    *ParsedElement[string] `json:"name"`
	Version *ParsedElement[string] `json:"version"`
}

func (vscs *ValueSetCodesSource) String() string {
	return fmt.Sprintf(
		"ValueSetCodesSource{\n  Name: %v,\n  Version: %v\n}",
		vscs.Name, vscs.Version,
	)
}

// ValueSetFilter represents a filter in a ValueSetComponent. Note if the value is a code with multiple strings,
// Value will be a concatenated string. This is a feature of the lexer and grammar.
type ValueSetFilter struct {
	Property *ParsedElement[string] `json:"property"`
	Operator *ParsedElement[string] `json:"operator"`
	Value    *ParsedElement[string] `json:"value"`
}

func (vsf *ValueSetFilter) String() string {
	return fmt.Sprintf(
		"ValueSetFilter{\n  Property: %v,\n  Operator: %v,\n  Value: %v\n}",
		vsf.Property, vsf.Operator, vsf.Value,
	)
}

// Profile represents a FSH Profile. This is a custom type, and not defined in the FSH grammar.
type Profile struct {
	Name         *ParsedElement[string] `json:"name"`
	Parent       *ParsedElement[string] `json:"parent"`
	ID           *ParsedElement[string] `json:"id"`
	Title        *ParsedElement[string] `json:"title"`
	Description  *ParsedElement[string] `json:"description"`
	ProfileRules *StructureDefRules     `json:"profileRules"`
}

func (p *Profile) String() string {
	return fmt.Sprintf(
		"Profile{\n  Name: %v,\n  Parent: %v,\n  ID: %v,\n  Title: %v,\n  Description: %v,\n  ProfileRules: %v\n}",
		p.Name, p.Parent, p.ID, p.Title, p.Description, p.ProfileRules,
	)
}

// StructureDefRules is an exhaustive list of rules in a Profile or Extension.
type StructureDefRules struct {
	CardRules       []*CardRule       `json:"cardRules"`
	FlagRules       []*FlagRule       `json:"flagRules"`
	BindingRules    []*BindingRule    `json:"bindingRules"`
	AssignmentRules []*AssignmentRule `json:"assignmentRules"`
	ContainsRules   []*ContainsRule   `json:"containsRules"`
	TypeRules       []*TypeRule       `json:"typeRules"`
	ObeysRules      []*ObeysRule      `json:"obeysRules"`
	CaretValueRules []*CaretValueRule `json:"caretValueRules"` // Assignment using caret values
	InsertRules     []*InsertRule     `json:"insertRules"`
	PathRules       []*PathRule       `json:"pathRules"`
}

func (pr *StructureDefRules) String() string {
	return fmt.Sprintf(
		"StructureDefRules{\n  CardRules: %v,\n  FlagRules: %v,\n  BindingRules: %v,\n  AssignmentRules: %v,\n  ContainsRules: %v,\n  TypeRules: %v,\n  ObeysRules: %v,\n  CaretValueRules: %v,\n  InsertRules: %v,\n  PathRules: %v\n}",
		pr.CardRules, pr.FlagRules, pr.BindingRules, pr.AssignmentRules, pr.ContainsRules, pr.TypeRules, pr.ObeysRules, pr.CaretValueRules, pr.InsertRules, pr.PathRules,
	)
}

// CodeSystem represents a FSH CodeSystem. This is a custom type, and not defined in the FSH grammar.
type CodeSystem struct {
	Name        *ParsedElement[string] `json:"name"`
	ID          *ParsedElement[string] `json:"id"`
	Title       *ParsedElement[string] `json:"title"`
	Description *ParsedElement[string] `json:"description"`
	Concepts    []*Concept             `json:"concepts"`
}

func (cs *CodeSystem) String() string {
	return fmt.Sprintf(
		"CodeSystem{\n  Name: %v,\n  ID: %v,\n  Title: %v,\n  Description: %v,\n  Concepts: %v\n}",
		cs.Name, cs.ID, cs.Title, cs.Description, cs.Concepts,
	)
}

// Concept represents a FSH Concept.
type Concept struct {
	Name        *ParsedElement[string] `json:"name"`
	Display     *ParsedElement[string] `json:"display"`
	Definition  *ParsedElement[string] `json:"definition"`
	SubConcepts []*Concept             `json:"subConcepts"`
}

func (c *Concept) String() string {
	return fmt.Sprintf("Concept{\n  Name: %v,\n  Display: %v,\n  Definition: %v,\n  SubConcepts: %v\n}", c.Name, c.Display, c.Definition, c.SubConcepts)
}

// Instance represents a FSH Instance. This is a custom type, and not defined in the FSH grammar.
type Instance struct {
	Name          *ParsedElement[string] `json:"name"`
	InstanceOf    *ParsedElement[string] `json:"instanceOf"`
	Title         *ParsedElement[string] `json:"title"`
	Description   *ParsedElement[string] `json:"description"`
	Usage         *ParsedElement[string] `json:"usage"`
	InstanceRules *InstanceRules         `json:"instanceRules"`
}

type InstanceRules struct {
	AssignmentRules []*AssignmentRule `json:"assignmentRules"`
	InsertRules     []*InsertRule     `json:"insertRules"`
	PathRules       []*PathRule       `json:"pathRules"`
}

// Extension represents a FSH Extension. This is a custom type, and not defined in the FSH grammar.
// Also note that, the Contexts list contains all contexts defined using the Context keyword, and
// any contexts set using caret value rules, will be found in the ExtensionRules.CaretValueRules list.
// See https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#defining-extensions for details.
type Extension struct {
	Name           *ParsedElement[string]   `json:"name"`
	ID             *ParsedElement[string]   `json:"id"`
	Title          *ParsedElement[string]   `json:"title"`
	Description    *ParsedElement[string]   `json:"description"`
	Parent         *ParsedElement[string]   `json:"parent"`
	Contexts       []*ParsedElement[string] `json:"contexts"`
	ExtensionRules *StructureDefRules       `json:"profileRules"`
}
