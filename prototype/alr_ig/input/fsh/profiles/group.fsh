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
* member.extension contains VaSelectionOnlyFlag named ext-vaSelectionOnlyFlag 0..1 MS
* characteristic 0..0

Mapping:  ALRGroupToALRCSV
Source:   ALRGroup
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report"
Id:       alr-csv-group
* -> "ALR Table 1-1"
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
* -> "ALR Table 1-5"
* member.extension[ext-changeReason] -> "PLUR_R05"
* member.extension[ext-changeReason] -> "AB_R01"
* member.extension[ext-changeReason] -> "HMO_R03"
* member.extension[ext-changeReason] -> "NO_US_R02"
* member.extension[ext-changeReason] -> "MDM_R04"
* member.extension[ext-changeReason] -> "NOFND_R06"
* -> "ALR Table 1-6"
* member.extension[ext-vaSelectionOnlyFlag] -> "VA_SELECTION_ONLY"
