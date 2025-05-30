Profile: VerilyTestProfile
Parent: TestParent
Id: verily-test-profile
Title: "Verily Test Profile"
Description: "This is a profile description."


// cardinality rules
* field 1..123 MS
* field2 ..123 MS
* field2 123..

// flag rule
* someElement MS SU ?! N TU D

// binding rules
* bindable from http://hl7.org/fhir/ValueSet (example)
* bindable from http://hl7.org/fhir/ValueSet (required)
* noStrength from $Alias

// assignment rules
* code = https://loinc.org#69548-6
* code = $LNC#69548-6 (exactly)
* status = #arrived
* active = true
* onsetDateTime = "2019-04-02"
* recordedDate = "2013-06-08T09:57:34.2112Z"
* extension[my-extension].valueInteger64 = 1234567890
* myCoding = http://hl7.org/fhir/CodeSystem/example-supplement|201801103#chol-mmol
* type = urn:iso-astm:E1762-95:2013#1.2.840.10065.1.12.1.2 "DisplayString"
* <CodeableConcept>.text = "{string}"
* myCodeableConcept.coding[0] = $SCT#363346000 "Malignant neoplastic disease (disorder)"
* valueQuantity = 55.0 'mm'
* valueQuantity = 55.0 'mm' "millimeter"
* subject = Reference(EveAnyperson)
* entry[0].resource = EveAnyperson

// contains rules
* extension contains
      $Disability named disability ..2 MS D and
      $GenderIdentity named genderIdentity 1.. MS
* component contains systolicBP 1..1 MS SU

// type rules
* valueQuantity only SimpleQuantity
* onset[x] only Period or Range
* value[x] only integer64
* performer only Reference(Practitioner or PractitionerRole)
* performer[Practitioner] only Reference(PrimaryCareProvider)
* action.definition[x] only Canonical(ActivityDefinition or PlanDefinition)
* activity.performedActivity only CodeableReference(Encounter or Procedure)

// obeys rules
* obeys us-core-9 and us-core-8
* name obeys us-core-8 and us-core-7

// caret value (assginment) rule
* testElement ^experimental = true
* ^url = "http://example.org/custom/myextension"

// insert rules
* insert RuleSet1
* path insert RuleSet2
* insert RuleSet3 (param1, param2, [[param3]])

// path rule
* pathRule