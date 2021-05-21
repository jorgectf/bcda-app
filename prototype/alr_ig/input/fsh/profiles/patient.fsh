// Your definitions here!
Profile:        ALRPatient
Parent:         http://hl7.org/fhir/us/davinci-atr/StructureDefinition/atr-patient
Id:             alr-Patient
Title:          "ALR Patient Profile"
Description:    "ALR Patient Profile Desc"
* extension contains ServiceCount named ext-serviceCount 0..* MS
* extension contains VoluntaryAlignmentFlag named ext-voluntaryAlignmentFlag 0..1 MS


Mapping:  ALRPatientToALRCSV1-1
Source:   ALRPatient
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report - Table 1-1"
Id:       alr-csv-patient-1-1
* -> "ALR Table 1-1"
* identifier -> "BENE_MBI_ID"
* name.given -> "BENE_1ST_NAME"
* name.family -> "BENE_LAST_NAME"
* gender -> "BENE_SEX_CD"
* birthDate -> "BENE_BRTH_DT"
* deceasedDateTime -> "BENE_DEATH_DT"
* address.district -> "GEO_SSA_CNTY_CD_NAME"
* address.state -> "GEO_SSA_STATE_NAME"


Mapping:  ALRPatientToALRCSV1-2
Source:   ALRPatient
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report - Table 1-2"
Id:       alr-csv-patient-1-2
* -> "ALR Table 1-2"
* identifier -> "BENE_MBI_ID"
* name.given -> "BENE_1ST_NAME"
* name.family -> "BENE_LAST_NAME"
* gender -> "BENE_SEX_CD"
* birthDate -> "BENE_BRTH_DT"
* deceasedDateTime -> "BENE_DEATH_DT"
* extension[ext-serviceCount].extension[participant] -> "MASTER_ID"
* extension[ext-serviceCount].extension[serviceCount] -> "B_EM_LINE_CNT_T"

Mapping:  ALRPatientToALRCSV1-3
Source:   ALRPatient
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report - Table 1-3"
Id:       alr-csv-patient-1-3
* -> "ALR Table 1-3"
* identifier -> "BENE_MBI_ID"
* name.given -> "BENE_1ST_NAME"
* name.family -> "BENE_LAST_NAME"
* gender -> "BENE_SEX_CD"
* birthDate -> "BENE_BRTH_DT"
* deceasedDateTime -> "BENE_DEATH_DT"
* extension[ext-serviceCount].extension[participant] -> "MASTER_ID"
* extension[ext-serviceCount].extension[serviceCount] -> "REV_LINE_CNT"

Mapping:  ALRPatientToALRCSV1-4
Source:   ALRPatient
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report - Table 1-4"
Id:       alr-csv-patient-1-4
* -> "ALR Table 1-4"
* identifier -> "BENE_MBI_ID"
* name.given -> "BENE_1ST_NAME"
* name.family -> "BENE_LAST_NAME"
* gender -> "BENE_SEX_CD"
* birthDate -> "BENE_BRTH_DT"
* deceasedDateTime -> "BENE_DEATH_DT"
* extension[ext-serviceCount].extension[participant] -> "MASTER_ID"
* extension[ext-serviceCount].extension[participant] -> "NPI_USED"
* extension[ext-serviceCount].extension[serviceCount] -> "PCS_COUNT"
