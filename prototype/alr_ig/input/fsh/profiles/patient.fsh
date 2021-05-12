// Your definitions here!
Profile:        ALRPatient
Parent:         http://hl7.org/fhir/us/davinci-atr/StructureDefinition/atr-patient
Id:             alr-Patient
Title:          "ALR Patient Profile"
Description:    "ALR Patient Profile Desc"
* extension contains ServiceCount named ext-serviceCount 0..* MS
* extension contains VoluntaryAlignmentFlag named ext-voluntaryAlignmentFlag 0..1 MS


Mapping:  ALRPatientToALRCSV
Source:   ALRPatient
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report"
Id:       alr-csv-patient
* -> "ALR Table 1-1"
* identifier -> "BENE_MBI_ID"
* name.given -> "BENE_1ST_NAME"
* name.family -> "BENE_LAST_NAME"
* gender -> "BENE_SEX_CD"
* birthDate -> "BENE_BRTH_DT"
* deceasedDateTime -> "BENE_DEATH_DT"
* address.district -> "GEO_SSA_CNTY_CD_NAME"
* address.state -> "GEO_SSA_STATE_NAME"
