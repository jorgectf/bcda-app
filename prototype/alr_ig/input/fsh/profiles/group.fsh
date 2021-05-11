// Your definitions here!
Profile:        ALRGroup
Parent:         http://hl7.org/fhir/us/davinci-atr/StructureDefinition/atr-group
Id:             alr-Group
Title:          "ALR Group Profile"
Description:    "ALR Group Profile Desc"
* member.extension contains ChangeReason named ext-changeReason 0..1 MS
* member.extension contains ClaimsBasedAssignmentFlag named ext-claimsBasedAssignmentFlag 0..1 MS
* member.extension contains ClaimsBasedAssignmentStep named ext-claimsBasedAssignmentStep 0..1 MS
* member.extension contains NewlyAssignedBeneficiaryFlag named ext-newlyAssignedBeneficiaryFlag 0..1 MS
* member.extension contains PrevAssignedBeneficiaryFlag named ext-pervAssignedBeneficiaryFlag 0..1 MS
* member.extension contains VoluntaryAlignmentFlag named ext-voluntaryAlignmentFlag 0..1 MS
* characteristic 0..0

Mapping:  ALRGroupToALRCSV
Source:   ALRGroup
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report"
Id:       alr-csv-group
* -> "Group"
* member.extension[changeType].valueCode -> "EXCLUDED"
* member.extension[ext-changeReason] -> "DECEASED_EXCLUDED"
* member.extension[ext-changeReason] -> "MISSING_ID_EXCLUDED"
* member.extension[ext-changeReason] -> "PART_A_B_ONLY_EXCLUDED"
* member.extension[ext-changeReason] -> "GHP_EXCLUDED"
* member.extension[ext-changeReason] -> "OUTSIDE_US_EXCLUDED"
* member.extension[ext-changeReason] -> "OTHER_SHARED_SAV_INT"

* member.extension[ext-claimsBasedAssignmentFlag] -> "CBA_FLAG"
* member.extension[ext-claimsBasedAssignmentStep] -> "ASSIGNMENT_TYPE"
* member.extension[ext-prevAssignedBeneficiaryFlag] -> "ASSIGNED_BEFORE"
* member.extension[ext-newlyAssignedBeneficiaryFlag] -> "ASG_STATUS"
* member.extension[ext-voluntaryAlignmentFlag] -> "IN_VA_MAX"
