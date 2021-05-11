Instance:       ExampleHccRiskFlags
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-HccRiskFlag
Title:          "HCC Risk Flags: General Example"
Usage:          #example
* id = "example-id-hcc-risk-flags"
* status = #final
* code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/hcc-risk-flags"
* code.coding[0].code = #hccRiskFlags
* code.coding[0].version = "V12"
* code.text = "HCC Risk Flags"
* subject.identifier[0].system = "https://bluebutton.cms.gov/resources/variables/bene_id"
* subject.identifier[0].value = "19990000000140"
* component[0].code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/hcc-risk-flags"
* component[0].code.coding[0].code = #HCC1
* component[0].code.coding[0].version = "V12"
* component[0].code.coding[0].display = "HIV/AIDS"
* component[0].valueString = "0"
* component[1].code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/hcc-risk-flags"
* component[1].code.coding[0].code = #HCC2
* component[1].code.coding[0].version = "V12"
* component[1].code.coding[0].display = "Septicemia/Shock"
* component[1].valueString = "1"
* component[2].code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/hcc-risk-flags"
* component[2].code.coding[0].code = #HCC5
* component[2].code.coding[0].version = "V12"
* component[2].code.coding[0].display = "Opportunistic Infections"
* component[2].valueString = "N/A"
* component[3].code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/hcc-risk-flags"
* component[3].code.coding[0].code = #HCC7
* component[3].code.coding[0].version = "V12"
* component[3].code.coding[0].display = "Metastatic Cancer and Acute Leukemia"
* component[3].valueString = "N/A"














/*
* subscriberId = "-19999000381"
* beneficiary.reference = "Patient/123"
* period.start = "2020-01-01"
* period.end = "2020-12-31"
* payor.reference = "Organization/123"
* extension[ext-enrollmentPeriod].extension[period][0].valuePeriod.start = "2020-01-01"
* extension[ext-enrollmentPeriod].extension[period][0].valuePeriod.end = "2020-06-30"
* extension[ext-enrollmentPeriod].extension[period][1].valuePeriod.start = "2020-10-01"
* extension[ext-enrollmentPeriod].extension[period][1].valuePeriod.end = "2020-12-31"


* managingEntity.reference = "Organization/123"
* managingEntity.display = "Sample Organization"
* meta.lastUpdated = "2021-04-01T03:04:12.348-04:00"
* extension[contractValidityPeriod].valuePeriod.start = "2020-04-01" 
* extension[contractValidityPeriod].valuePeriod.end = "2021-03-31"
* member.extension[changeType].valueCode = #dropped 
* member.extension[coverageReference].valueReference.reference = "Coverage/123"
* member.extension[attributedProvider].valueReference.reference = "Practitioner/XYZ"
* member.extension[ext-changeReason].extension[reasonCode][0].valueCoding.code = #GHP_EXCLUDED
* member.extension[ext-changeReason].extension[reasonCode][0].valueCoding.system = "http://cms.gov/ig/valueset-changeReason"
* member.extension[ext-changeReason].extension[reasonCode][0].valueCoding.display = "Beneficiary had at least one month of Part-A only or Part-B only coverage"
* member.extension[ext-changeReason].extension[reasonCode][1].valueCoding.code = #MISSING_ID_EXCLUDED
* member.extension[ext-changeReason].extension[reasonCode][1].valueCoding.system = "http://cms.gov/ig/valueset-changeReason"
* member.extension[ext-changeReason].extension[reasonCode][1].valueCoding.display = "Beneficiary identifier is missing"
* member.extension[ext-claimsBasedAssignmentFlag].valueBoolean = true
* member.extension[ext-claimsBasedAssignmentStep].valueInteger = 3
* member.extension[ext-prevAssignedBeneficiaryFlag].valueBoolean = true
* member.extension[ext-newlyAssignedBeneficiaryFlag].valueBoolean = false
* member.extension[ext-voluntaryAlignmentFlag].valueBoolean = false
* member[0].entity.reference = "Patient/123"

*/
