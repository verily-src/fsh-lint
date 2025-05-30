ValueSet: TestValueSet_VS
Id: test-value-set
Title: "Test Value Set"
Description: "This is a value set description."

// include/exclude rules (value set component)
* $codeAlias#code "Example code"
* include http://example-url#code "Another example code"
* include codes from system http://system-url|version and valueset http://valueset1-url|version and http://valueset2-url|version where property descendant-of #some-parent
* exclude codes from valueset http://valueset1-url|version and http://valueset2-url|version and system http://system-url|version where property is-a #123 "Some concept"

// caret value rule
* ^slice[0].field = #code "code description"

// code caret value rule
* $code#codeNum #child ^some.field = $a#code "code description"

// insert rules
* insert RuleSet1
* path insert RuleSet2
* insert RuleSet3 (param1, param2, [[param3]])

// code insert rules
* #code-one insert RuleSet4
* #code-one #code-two #code-three insert RuleSet5