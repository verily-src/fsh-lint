# Rules

## binding-strength-present

### Description

Binding strength (example, preferred, extensible, or required) must be present for every binding
rule. Binding strength is specified by appending the strength in parentheses to the end of the
binding rule.

### Examples

Below is a binding rule where the strength is set to `extensible`.

```fsh
* code from CancerConditionVS (extensible)
```

### Scope

This rule applies to all binding rules which are found only in:

- Profiles
- Extensions

### Resources

- [Binding Strengths](https://hl7.org/fhir/R5/valueset-binding-strength.html)
- [Binding Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html##binding-rules)

## code-system-name-matches-filename

### Description

Code system name must match filename (where filename does not include `_CS`). This match is case
sensitive, and the file extension must be `.fsh`.

### Examples

Code system name `ExampleOne_CS` **matches** filename `ExampleOne.fsh`, but **does not match**
`ExampleOne_CS.fsh`, `exampleone.fsh`, or `ExampleOne.txt`.

### Scope

This rule applies to all code systems.

## code-system-name-matches-id

### Description

Code system name (PascalCase) must match code system id in kebab-case without `_CS` suffix.

If there are numbers in the code system name, **any grouping with its adjacent alphabetic characters
is acceptable**. In other words, hyphens in the id are optional when adjacent to a number. This
flexible hyphenation applies only to numbers.

If there are acronyms in the code system name, the acronym should be grouped as one word. Note, that
the first letter of a word following an acronym should be capitalized in the PascalCase name.

Note, the "Want" value that is generated, is one of the correct configurations where each group of
numbers is grouped as their own separate word. This is the behavior of
[strcase.ToDelimited](https://pkg.go.dev/github.com/iancoleman/strcase#ToDelimited).

### Examples

- Code system name `ExampleOne_CS` **matches** id `example-one`, but **does not match**
  `example-one-cs` or `ExampleOne`.
- Code system name `A123Example_CS` **matches** id `a-123-example`, `a123example`, `a1-23example`,
  and several others.
- Code system name `ABCExample_CS` **matches** id `abc-example`, but **does not match**
  `a-b-c-example` or `abce-xample`.

### Scope

This rule applies to all code systems.

## code-system-name-matches-title

### Description

Code System name without the `_CS` (PascalCase) must match Code System title in Title Case (space
separated). This check is **case insensitive**.

If there are numbers in the code system name, **any grouping with its adjacent alphabetic characters
is acceptable**. In other words, the space is optional in the title when adjacent to a number. This
flexibility applies only to numbers.

If there are acronyms in the code system name, the acronym should be grouped as one word. Note, that
the first letter of a word following an acronym should be capitalized in the PascalCase name.

Note, the "Want" value that is generated, is one of the correct configurations where each group of
numbers is grouped as their own separate word. This is the behavior of
[strcase.ToDelimited](https://pkg.go.dev/github.com/iancoleman/strcase#ToDelimited).

### Examples

- Code system name `Example_One_CS` **matches** title `Example One` and `example one`, but **does
  not match** `Example One CS` or `ExampleOne`.
- Code system name `A123Example_CS` **matches** title `A 123 Example`, `A123example`,
  `A1 23example`, and several others.
- Code system name `ABCExample_CS` **matches** title `ABC Example`, but **does not match**
  `A B C Example` or `Abce Xample`

### Scope

This rule applies to all code systems.

## profile-assignment-present

### Description

All profiles should have the fields listed below set in an assignment rule (caret value rule).

- status
- abstract

### Examples

A profile with both status and abstract fields correctly set:

```FSH
Profile: Example

* ^status = #retired
* ^abstract = false
```

### Scope

This rule applies to all profiles.

### Resources

- `abstract` must be set to `true` or `false`.
- `status` must be set to one of the statuses defined
  [here](https://build.fhir.org/structuredefinition-definitions.html#:~:text=the%20root%20element.-,StructureDefinition.status,-Element%20Id).

## profile-name-matches-filename

### Description

Profile name must match filename. This match is case-sensitive, and the file extension must be
`.fsh`.

### Examples

Profile name `ExampleProfile` profile name **matches** `ExampleProfile.fsh`, but **does not match**
`exampleprofile.fsh` or `ExampleProfile.txt`.

### Scope

This rule applies to all profiles.

## profile-name-matches-id

### Description

Profile name (PascalCase) must match profile id in kebab-case.

If there are numbers in the profile name, **any grouping with its adjacent alphabetic characters is
acceptable**. In other words, hyphens in the id are optional when adjacent to a number. This
flexible hyphenation applies only to numbers.

If there are acronyms in the profile name, the acronym should be grouped as one word. Note, that the
first letter of a word following an acronym should be capitalized in the PascalCase name.

Note, the "Want" value that is generated, is one of the correct configurations where each group of
numbers is grouped as their own separate word. This is the behavior of
[strcase.ToDelimited](https://pkg.go.dev/github.com/iancoleman/strcase#ToDelimited).

### Examples

- Profile name `ExampleProfile` **matches** `example-profile`, but **does not match**
  `exampleprofile`.
- Profile name `A123Example` **matches** `a-123-example`, `a123example`, `a1-23example`, and several
  others.
- Profile name `ABCExample` **matches** id `abc-example`, but **does not match** `a-b-c-example` or
  `abce-xample`.

### Scope

This rule applies to all profiles.

## profile-name-matches-title

### Description

Profile name (PascalCase) must match profile title in Title Case (space separated) and ending with
"Profile". This check is **case insensitive**.

If there are numbers in the profile name, **any grouping with its adjacent alphabetic characters is
acceptable**. In other words, the space is optional in the title when adjacent to a number. This
flexibility applies only to numbers.

If there are acronyms in the profile name, the acronym should be grouped as one word. Note, that the
first letter of a word following an acronym should be capitalized in the PascalCase name.

Note, the "Want" value that is generated, is one of the correct configurations where each group of
numbers is grouped as their own separate word. This is the behavior of
[strcase.ToDelimited](https://pkg.go.dev/github.com/iancoleman/strcase#ToDelimited).

### Examples

- Profile name `ExampleOne` **matches** title `Example One Profile`, but **does not match**
  `Example One`.
- Profile name `A123Example` **matches** `A 123 Example`, `A123example`, `A1 23example`, and several
  others.
- Profile name `ABCExample` **matches** `ABC Example`, but **does not match** `A B C Example` or
  `Abce Xample`.

### Scope

This rule applies to all profiles.

## required-field-present

### Description

Required fields should not be missing. If this rule fails, **all other rules WILL NOT run**. This
was made to simplify nil checks in all other rules that depend on the fact that required fields will
not be nil. Required fields so far are:

- Code System
  - Name
  - ID
  - Title
- Profile
  - Name
  - ID
  - Title
- Value Set
  - Name
  - ID
  - Title

### Scope

This rule can run against all declaration types, depending on the required fields set.

## value-set-name-matches-filename

### Description

Value Set name must match filename (where filename does not include `_VS`). This match is case
sensitive, and the file extension must be `.fsh`.

### Examples

Value set name `Example_VS` **matches** `Example.fsh`, but **does not match** `Example_VS.fsh`, or
`example.fsh`, or `example.txt`

### Scope

This rule applies to all value sets.

## value-set-name-matches-id

### Description

Value set name (PascalCase) must match value set id in kebab-case without `_VS` suffix.

If there are numbers in the value set name, **any grouping with its adjacent alphabetic characters
is acceptable**. In other words, hyphens in the id are optional when adjacent to a number. This
flexible hyphenation applies only to numbers.

If there are acronyms in the value set name, the acronym should be grouped as one word. Note, that
the first letter of a word following an acronym should be capitalized in the PascalCase name.

Note, the "Want" value that is generated, is one of the correct configurations where each group of
numbers is grouped as their own separate word. This is the behavior of
[strcase.ToDelimited](https://pkg.go.dev/github.com/iancoleman/strcase#ToDelimited).

### Examples

- Value set name `Example_VS` **matches** id `example`, but **does not match** `example-vs` or
  `Example`.
- Value set name `A123Example_VS` **matches** id `a-123-example`, `a123example`, `a1-23example`, and
  several others.
- Value set name `ABCExample_VS` **matches** id `abc-example`, but **does not match**
  `a-b-c-example` or `abce-xample`.

### Scope

This rule applies to all value sets.

## value-set-name-matches-title

### Description

Value Set name without the `_VS` (PascalCase) must match Value Set title in Title Case (space
separated). This check is **case insensitive**.

If there are numbers in the value set name, **any grouping with its adjacent alphabetic characters
is acceptable**. In other words, the space is optional in the title when adjacent to a number. This
flexibility applies only to numbers.

If there are acronyms in the value set name, the acronym should be grouped as one word. Note, that
the first letter of a word following an acronym should be capitalized in the PascalCase name.

Note, the "Want" value that is generated, is one of the correct configurations where each group of
numbers is grouped as their own separate word. This is the behavior of
[strcase.ToDelimited](https://pkg.go.dev/github.com/iancoleman/strcase#ToDelimited).

### Examples

- Value set name `Example_One_VS` **matches** title `Example One` and `example one`, but **does not
  match** `Example One VS` or `ExampleOne`.
- Value set name `A123Example_VS` **matches** title `A 123 Example`, `A123example`, `A1 23example`,
  and several others.
- Value set name `ABCExample_VS` **matches** title `ABC Example`, but **does not match**
  `A B C Example` or `Abce Xample`

### Scope

This rule applies to all value sets.
