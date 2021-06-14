Alias:   CT = http://hl7.org/fhir/us/davinci-atr/STU1/ValueSet-atr-changetype

Instance:       ExampleGroup
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-Group
Title:          "Group: General Example"
Usage:          #example
* id = "example-id-group"
* active = true
* text.div = "<div xmlns=\"http://www.w3.org/1999/xhtml\"><p>General example of a Group resource</p></div>"
* text.status = #additional
* type = #person
* actual = true
* name = "General Example of ALR Group"
* managingEntity.reference = "Organization/123"
* managingEntity.display = "Sample Organization"
* meta.lastUpdated = "2021-04-01T03:04:12.348-04:00"
* extension[contractValidityPeriod].valuePeriod.start = "2020-04-01" 
* extension[contractValidityPeriod].valuePeriod.end = "2021-03-31"
* member[0].extension[changeType].valueCode = #excluded
* member[0].extension[coverageReference].valueReference.reference = "Coverage/example-id-coverage"
* member[0].extension[attributedProvider].valueReference.reference = "Practitioner/example-id-prac"
* member[0].extension[ext-changeReason].extension[reasonCode][0].valueCoding.code = #GHP_EXCLUDED
* member[0].extension[ext-changeReason].extension[reasonCode][0].valueCoding.system = "https://bluebutton.cms.gov/resources/variables/alr/changeReason"
* member[0].extension[ext-changeReason].extension[reasonCode][0].valueCoding.display = "Beneficiary had at least one month of Part-A only or Part-B only coverage"
* member[0].extension[ext-changeReason].extension[reasonCode][1].valueCoding.code = #MISSING_ID_EXCLUDED
* member[0].extension[ext-changeReason].extension[reasonCode][1].valueCoding.system = "https://bluebutton.cms.gov/resources/variables/alr/changeReason"
* member[0].extension[ext-changeReason].extension[reasonCode][1].valueCoding.display = "Beneficiary identifier is missing"
* member[0].extension[ext-claimsBasedAssignmentFlag].valueBoolean = true
* member[0].extension[ext-claimsBasedAssignmentStep].valueInteger = 3
* member[0].extension[ext-prevAssignedBeneficiaryFlag].valueBoolean = true
* member[0].extension[ext-newlyAssignedBeneficiaryFlag].valueBoolean = false
* member[0].extension[ext-voluntaryAlignmentFlag].valueBoolean = false
* member[0].entity.reference = "Patient/example-id-patient"
//
* member[1].extension[changeType].valueCode = #nochange
* member[1].extension[coverageReference].valueReference.reference = "Coverage/example-id-coverage"
* member[1].extension[attributedProvider].valueReference.reference = "Practitioner/example-id-prac"
* member[1].extension[ext-claimsBasedAssignmentFlag].valueBoolean = false
* member[1].extension[ext-claimsBasedAssignmentStep].valueInteger = 2
* member[1].extension[ext-prevAssignedBeneficiaryFlag].valueBoolean = false
* member[1].extension[ext-newlyAssignedBeneficiaryFlag].valueBoolean = true
* member[1].extension[ext-voluntaryAlignmentFlag].valueBoolean = true
* member[1].entity.reference = "Patient/example-id-2-patient"
//
* member[2].extension[changeType].valueCode = #new
* member[2].extension[coverageReference].valueReference.reference = "Coverage/example-id-coverage"
* member[2].extension[attributedProvider].valueReference.reference = "Practitioner/example-id-prac"
* member[2].extension[ext-claimsBasedAssignmentFlag].valueBoolean = true
* member[2].extension[ext-claimsBasedAssignmentStep].valueInteger = 1
* member[2].extension[ext-prevAssignedBeneficiaryFlag].valueBoolean = false
* member[2].extension[ext-newlyAssignedBeneficiaryFlag].valueBoolean = true
* member[2].extension[ext-voluntaryAlignmentFlag].valueBoolean = true
* member[2].entity.reference = "Patient/example-id-3-patient"
//
* member[3].extension[changeType].valueCode = #dropped
* member[3].extension[coverageReference].valueReference.reference = "Coverage/example-id-coverage"
* member[3].extension[attributedProvider].valueReference.reference = "Practitioner/example-id-prac"
* member[3].extension[ext-changeReason].extension[reasonCode][0].valueCoding.code = ##PLUR_R05
* member[3].extension[ext-changeReason].extension[reasonCode][0].valueCoding.system = "https://bluebutton.cms.gov/resources/variables/alr/changeReason"
* member[3].extension[ext-changeReason].extension[reasonCode][0].valueCoding.display = "Beneficiary did not receive the plurality of their primary care services from the ACO"
* member[3].extension[ext-changeReason].extension[reasonCode][1].valueCoding.code = #NO_US_R02
* member[3].extension[ext-changeReason].extension[reasonCode][1].valueCoding.system = "https://bluebutton.cms.gov/resources/variables/alr/changeReason"
* member[3].extension[ext-changeReason].extension[reasonCode][1].valueCoding.display = "Beneficiary does not reside in the United States"
* member[3].extension[ext-claimsBasedAssignmentFlag].valueBoolean = false
* member[3].extension[ext-claimsBasedAssignmentStep].valueInteger = 0
* member[3].extension[ext-prevAssignedBeneficiaryFlag].valueBoolean = true
* member[3].extension[ext-newlyAssignedBeneficiaryFlag].valueBoolean = false
* member[3].extension[ext-voluntaryAlignmentFlag].valueBoolean = false
* member[3].entity.reference = "Patient/example-id-patient"

