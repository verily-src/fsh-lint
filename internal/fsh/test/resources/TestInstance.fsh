Instance: VerilySearchParameterInstanceExample
InstanceOf: SearchParameter
Title: "Instance Title"
Usage: #definition
Description: "This is a description."

// assignment rules
* url = "http://fhir.example.com/oneverily/SearchParameter/test-parameter"
* name = "TestParameter"
* status = #active
* expression = """
ResearchStudy.extension('http://fhir.example.com/StructureDefinition/test-parameter').extension('artifact-canonicalReference').valueCanonical |
HealthcareService.extension('http://fhir.example.com/StructureDefinition/test-parameter').extension('artifact-canonicalReference').valueCanonical
"""
* target[0] = #PlanDefinition

// insert rules
* insert RuleSet1
* path insert RuleSet2
* insert RuleSet3 (param1, param2, [[param3]])

// path rule
* somePath
