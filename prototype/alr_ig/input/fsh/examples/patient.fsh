Instance:       ExamplePatient
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-Patient
Title:          "Patient: General Example"
Usage:          #example
* id = "example-id-patient"
* text.status = #additional
* text.div = "<div xmlns=\"http://www.w3.org/1999/xhtml\"><p>General example of a patient, including extensions for the service counts found in the ALR data dictionary: table 1-2, 1-3 and 1-4</p></div>"
* identifier[0].system = "https://bluebutton.cms.gov/resources/variables/bene_id"
* identifier[0].value = "example-ccw-id"
* identifier[1].system = "https://bluebutton.cms.gov/resources/identifier/mbi-hash"
* identifier[1].value = "412be343c343343324f661527671abb5307d5e89031db9a9f"
* identifier[2].system = "http://hl7.org/fhir/sid/us-mbi"
* identifier[2].value = "example-mbi-patient"
// TODO: there is a BB extension that goes with us-mbi that dictates whether mbi is current or historical
// but I"m basing this Patient off of CARIN, not BB so....
* name.given[0] = "John"
* name.family = "Doe"
* gender = #male
* birthDate = "1957-01-17"
* deceasedDateTime = "2020-03-18"
* address.line[0] = "1000 Main Street"
* address.line[1] = "Apt #2"
* address.city = "Anytown"
* address.state = "UT"
* address.district = "SL County"
/*
// table 1-2
* extension[ext-serviceCount][0].extension[serviceCountCode][0].valueCoding.system = "https://alr.cms.gov/resources/variables/service-count-codes"
* extension[ext-serviceCount][0].extension[serviceCountCode][0].valueCoding.code = #B_EM_LINE_CNT_T
* extension[ext-serviceCount][0].extension[serviceCountCode][0].valueCoding.display = "Count of Primary Care Services - ACO Participant TIN"
* extension[ext-serviceCount][0].extension[participant][0].valueIdentifier.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCount][0].extension[participant][0].valueIdentifier.value = "example-TIN"
* extension[ext-serviceCount][0].extension[participant][0].valueIdentifier.type.coding.code = #TAX
* extension[ext-serviceCount][0].extension[participant][0].valueIdentifier.type.coding.display = "Tax ID number"
* extension[ext-serviceCount][0].extension[participant][0].valueIdentifier.type.coding.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCount][0].extension[serviceCount].valueInteger = 5
// table 1-3
* extension[ext-serviceCount][1].extension[serviceCountCode][0].valueCoding.system = "https://alr.cms.gov/resources/variables/service-count-codes"
* extension[ext-serviceCount][1].extension[serviceCountCode][0].valueCoding.code = #REV_LINE_CNT
* extension[ext-serviceCount][1].extension[serviceCountCode][0].valueCoding.display = "Count of Primary Care Services - ACO Participant CCN"
* extension[ext-serviceCount][1].extension[participant][0].valueIdentifier.system = "https://bluebutton.cms.gov/resources/variables/prvdr_num"
* extension[ext-serviceCount][1].extension[participant][0].valueIdentifier.value = "example-CCN"
* extension[ext-serviceCount][1].extension[participant][0].valueIdentifier.type.coding.code = #CCN
* extension[ext-serviceCount][1].extension[participant][0].valueIdentifier.type.coding.display = "CCN Number"
* extension[ext-serviceCount][1].extension[participant][0].valueIdentifier.type.coding.system = "https://bluebutton.cms.gov/resources/variables/prvdr_num"
* extension[ext-serviceCount][1].extension[serviceCount].valueInteger = 15
// table 1-4
* extension[ext-serviceCount][2].extension[serviceCountCode][0].valueCoding.system = "https://alr.cms.gov/resources/variables/service-count-codes"
* extension[ext-serviceCount][2].extension[serviceCountCode][0].valueCoding.code = #PCS_COUNT
* extension[ext-serviceCount][2].extension[serviceCountCode][0].valueCoding.display = "Top ACO Participant TIN-Individual NPI"
* extension[ext-serviceCount][2].extension[participant][0].valueIdentifier.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCount][2].extension[participant][0].valueIdentifier.value = "example-TIN"
* extension[ext-serviceCount][2].extension[participant][0].valueIdentifier.type.coding.code = #TAX
* extension[ext-serviceCount][2].extension[participant][0].valueIdentifier.type.coding.display = "Tax ID number"
* extension[ext-serviceCount][2].extension[participant][0].valueIdentifier.type.coding.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCount][2].extension[participant][1].valueIdentifier.system = "http://hl7.org/fhir/sid/us-npi"
* extension[ext-serviceCount][2].extension[participant][1].valueIdentifier.value = "example-NPI"
* extension[ext-serviceCount][2].extension[participant][1].valueIdentifier.type.coding.code = #NPI
* extension[ext-serviceCount][2].extension[participant][1].valueIdentifier.type.coding.display = "NPI Number"
* extension[ext-serviceCount][2].extension[participant][1].valueIdentifier.type.coding.system = "http://hl7.org/fhir/sid/us-npi"
* extension[ext-serviceCount][2].extension[serviceCount].valueInteger = 7
// table 1-6
// TODO: table 1-5 and table 1-6 may be "different" in that they may be a patient population "not currently assigned"
* extension[ext-voluntaryAlignmentFlag].valueBoolean = false
*/
/*
// table 1-2
* extension[ext-serviceCount2][0].extension[participant][0].valueIdentifier.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCount2][0].extension[participant][0].valueIdentifier.value = "example-TIN"
* extension[ext-serviceCount2][0].extension[participant][0].valueIdentifier.type.coding.code = #TAX
* extension[ext-serviceCount2][0].extension[participant][0].valueIdentifier.type.coding.display = "Tax ID number"
* extension[ext-serviceCount2][0].extension[participant][0].valueIdentifier.type.coding.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCount2][0].extension[serviceCount].valueQuantity.value = 5
* extension[ext-serviceCount2][0].extension[serviceCount].valueQuantity.system = "https://alr.cms.gov/resources/variables/service-count-codes"
* extension[ext-serviceCount2][0].extension[serviceCount].valueQuantity.code = #B_EM_LINE_CNT_T
* extension[ext-serviceCount2][0].extension[serviceCount].valueQuantity.unit = "Count of Primary Care Services by ACO Participant TIN"
*/

// table 1-2
* extension[ext-serviceCountTIN][0].extension[participant][0].valueIdentifier.value = "example-TIN"
* extension[ext-serviceCountTIN][0].extension[participant][0].valueIdentifier.type.coding.code = #TAX
* extension[ext-serviceCountTIN][0].extension[participant][0].valueIdentifier.type.coding.display = "Tax ID number"
* extension[ext-serviceCountTIN][0].extension[participant][0].valueIdentifier.type.coding.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCountTIN][0].extension[serviceCount].valueInteger = 5

// table 1-3
* extension[ext-serviceCountCCN][0].extension[participant][0].valueIdentifier.value = "example-CCN"
* extension[ext-serviceCountCCN][0].extension[participant][0].valueIdentifier.type.coding.code = #CCN
* extension[ext-serviceCountCCN][0].extension[participant][0].valueIdentifier.type.coding.display = "CCN number"
* extension[ext-serviceCountCCN][0].extension[participant][0].valueIdentifier.type.coding.system = "https://bluebutton.cms.gov/resources/variables/prvdr_num"
* extension[ext-serviceCountCCN][0].extension[serviceCount].valueInteger = 15



// table 1-4
* extension[ext-serviceCountTIN-NPI][0].extension[participant][0].valueIdentifier.value = "example-TIN"
* extension[ext-serviceCountTIN-NPI][0].extension[participant][0].valueIdentifier.type.coding.code = #TAX
* extension[ext-serviceCountTIN-NPI][0].extension[participant][0].valueIdentifier.type.coding.display = "Tax ID number"
* extension[ext-serviceCountTIN-NPI][0].extension[participant][0].valueIdentifier.type.coding.system = "http://terminology.hl7.org/CodeSystem/v2-0203"
* extension[ext-serviceCountTIN-NPI][0].extension[participant][1].valueIdentifier.value = "example-NPI"
* extension[ext-serviceCountTIN-NPI][0].extension[participant][1].valueIdentifier.type.coding.code = #NPI
* extension[ext-serviceCountTIN-NPI][0].extension[participant][1].valueIdentifier.type.coding.display = "NPI number"
* extension[ext-serviceCountTIN-NPI][0].extension[participant][1].valueIdentifier.type.coding.system = "http://hl7.org/fhir/sid/us-npi"
* extension[ext-serviceCountTIN-NPI][0].extension[serviceCount].valueInteger = 7



