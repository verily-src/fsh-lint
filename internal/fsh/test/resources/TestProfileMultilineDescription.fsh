Profile: VerilyLocation3PData
Parent: VerilyLocationBaseR4
Id: verily-location-3p-data
Title: "Verily Location 3PD"
Description: 
"""
This is
a multiline
description
"""

* ^contact.name = "system:ingestion"
* ^contact.telecom[0].system = #url
* ^contact.telecom[=].value = $IngestionSystemUrl

// align with US Core Location Profile
* status MS // Must support in US Core
* name MS // Required in US Core
* type MS // Must support in US Core
* telecom MS // Must support in US Core
* address MS // Must support in US Core
* managingOrganization MS // Must support in US Core

// verily requirements
* identifier 1..* MS // at least one identifier for third party source (e.g. Zus resource id)

// Zus identifier
* identifier ^slicing.discriminator.type = #value
* identifier ^slicing.discriminator.path = "system"
* identifier ^slicing.rules = #open
* identifier ^slicing.description = "Identifier from third party source"
* identifier ^slicing.ordered = false