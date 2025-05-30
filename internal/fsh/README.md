# FSHParser

This library will parse FSH resources to an internal format. Antlr is used along with
[FSH Grammar](https://github.com/FHIR/sushi/tree/master/antlr/src/main/antlr) to parse the FSH file
content.

## Contributing

Whenever a new parse capability for any type is added, add it to the Parsing Implemented
[List](#parsing-implemented).

### Tips

Read up on [Visitor Pattern](https://refactoring.guru/design-patterns/visitor) Create your own
custom return type in types package.

In fshparser, this part will remain the same.

```go
    stream := antlr.NewInputStream(fshData)
    lexer := g.NewFSHLexer(stream)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    p := g.NewFSHParser(tokens)
    tree := p.[ResourceType]()
    v := &pr.FSHVisitor{}
    o := tree.Accept(v)
    res, ok := o.(t.<CustomValueObject>)
    if !ok {
        return t.<CustomValueObject>{}, fmt.Errorf("got type %T, expected types.ValueSet", o)
    }
    return valueSet, nil
```

View the tree to get more information on how the parsing happens

```go
    st := tree.ToStringTree(p.RuleNames, p)
    fmt.Println(st)
```

In [visitor.go](internal/parser/visitor.go), implement `Visit<ResourceType>` function. Look at
`VisitValueSet` for an example.

## Parsing Implemented

- [ ] Alias

  - [ ] Name
  - [ ] Value

- [x] CodeSystem

  - [x] Name
  - [x] ID
  - [x] Parent
  - [x] Description
  - [x] Concepts
  - [ ] Rules
    - [ ] CodeCaretValueRule
    - [ ] CodeInsertRule

- [x] Extension

  - [x] Name
  - [x] Parent
  - [x] ID
  - [x] Title
  - [x] Description
  - [x] Contexts
  - [x] Rules
    - [x] [Cardinality Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#cardinality-rules)
    - [x] [Flag Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#flag-rules)
    - [x] [Binding Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#binding-rules)
          (called
          [valueSetRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L85)
          in the grammar)
    - [x] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [x] [Contains Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#contains-rules-for-extensions)
    - [x] [Type Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#type-rules)
          (called
          [onlyRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L88)
          in the grammar)
    - [x] [Obeys Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#obeys-rules)
    - [x] [Caret Value Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignments-with-caret-paths)
    - [x] [Insert Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [x] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)

- [x] Instance

  - [x] Name
  - [x] InstanceOf
  - [x] Title
  - [x] Description
  - [x] Usage
  - [x] Rules
    - [x] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [x] [InsertRules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [x] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)

- [ ] Invariant

  - [ ] Name
  - [ ] Description
  - [ ] Expression
  - [ ] XPath
  - [ ] Severity
  - [ ] Rules
    - [ ] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [ ] [InsertRules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [ ] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)

- [ ] Logical

  - [ ] Name
  - [ ] Parent
  - [ ] ID
  - [ ] Title
  - [ ] Description
  - [ ] Characteristics
  - [ ] Rules
    - [ ] [Cardinality Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#cardinality-rules)
    - [ ] [Flag Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#flag-rules)
    - [ ] [Binding Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#binding-rules)
          (called
          [valueSetRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L85)
          in the grammar)
    - [ ] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [ ] [Contains Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#contains-rules-for-extensions)
    - [ ] [Type Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#type-rules)
          (called
          [onlyRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L88)
          in the grammar)
    - [ ] [Obeys Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#obeys-rules)
    - [ ] [Caret Value Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignments-with-caret-paths)
    - [ ] [Insert Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [ ] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)
    - [ ] [AddElementRule](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#add-element-rules)
    - [ ] [AddCRElementRule](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#add-element-rules)
          (Add Content Reference Element)

- [ ] Mapping

  - [ ] Name
  - [ ] ID
  - [ ] Source
  - [ ] Target
  - [ ] Description
  - [ ] Title
  - [ ] Rules
    - [ ] MappingRule
    - [ ] [InsertRules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [ ] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)

- [x] Profile

  - [x] Name
  - [x] Parent
  - [x] ID
  - [x] Title
  - [x] Description
  - [x] Rules
    - [x] [Cardinality Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#cardinality-rules)
    - [x] [Flag Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#flag-rules)
    - [x] [Binding Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#binding-rules)
          (called
          [valueSetRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L85)
          in the grammar)
    - [x] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [x] [Contains Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#contains-rules-for-extensions)
    - [x] [Type Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#type-rules)
          (called
          [onlyRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L88)
          in the grammar)
    - [x] [Obeys Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#obeys-rules)
    - [x] [Caret Value Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignments-with-caret-paths)
    - [x] [Insert Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [x] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)

- [ ] ParamRuleSet (This may become a part of RuleSet)

  - [ ] Parameters
  - [ ] Rules

- [ ] Resource

  - [ ] Name
  - [ ] Parent
  - [ ] ID
  - [ ] Title
  - [ ] Description
  - [ ] Rules
    - [ ] [Cardinality Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#cardinality-rules)
    - [ ] [Flag Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#flag-rules)
    - [ ] [Binding Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#binding-rules)
          (called
          [valueSetRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L85)
          in the grammar)
    - [ ] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [ ] [Contains Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#contains-rules-for-extensions)
    - [ ] [Type Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#type-rules)
          (called
          [onlyRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L88)
          in the grammar)
    - [ ] [Obeys Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#obeys-rules)
    - [ ] [Caret Value Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignments-with-caret-paths)
    - [ ] [Insert Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [ ] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)
    - [ ] [AddElementRule](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#add-element-rules)
    - [ ] [AddCRElementRule](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#add-element-rules)
          (Add Content Reference Element)

- [ ] RuleSet

  - [ ] Name
  - [ ] Rules
    - [ ] [Cardinality Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#cardinality-rules)
    - [ ] [Flag Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#flag-rules)
    - [ ] [Binding Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#binding-rules)
          (called
          [valueSetRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L85)
          in the grammar)
    - [ ] [Assignment Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignment-rules)
          (called
          [fixedValueRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L86)
          in the grammar)
    - [ ] [Contains Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#contains-rules-for-extensions)
    - [ ] [Type Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#type-rules)
          (called
          [onlyRule](https://github.com/verily-src/verily1/blob/main/common/fsh/internal/grammar/FSH.g4#L88)
          in the grammar)
    - [ ] [Obeys Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#obeys-rules)
    - [ ] [Caret Value Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignments-with-caret-paths)
    - [ ] [Insert Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [ ] [Path Rules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#path-rules)
    - [ ] Add Element Rule
    - [ ] Add CR Element Rule
    - [ ] Concepts
    - [ ] CaretValueRules
          ([Assignment Rules with Caret Paths and Coding](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules:~:text=authors%20MAY%20choose%20to%20repeat%20the%20code))
    - [x] CodeInsertRules
          ([Insert Rules with the Concept Code as the context](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#inserting-rule-sets-with-path-context:~:text=inserted%20in%20the%20context%20of%20a%20concept))
    - [ ] Components
    - [ ] MappingRule

- [x] ValueSet
  - [x] Name
  - [x] ID
  - [x] Description
  - [x] [Components](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#defining-value-sets)
  - [x] Rules
    - [x] CaretValueRules
          ([Assignment Rules with Caret Paths](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#assignments-with-caret-paths))
    - [x] CodeCaretValueRules
          ([Assignment Rules with Caret Paths and Coding](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules:~:text=authors%20MAY%20choose%20to%20repeat%20the%20code))
    - [x] [InsertRules](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#insert-rules)
    - [x] CodeInsertRules
          ([Insert Rules with the Concept Code as the context](https://build.fhir.org/ig/HL7/fhir-shorthand/reference.html#inserting-rule-sets-with-path-context:~:text=inserted%20in%20the%20context%20of%20a%20concept))
