// Your definitions here!
Profile:        ALRPatient
Parent:         http://hl7.org/fhir/us/davinci-atr/StructureDefinition/atr-patient
Id:             alr-Patient
Title:          "ALR Patient Profile"
Description:    "ALR Patient Profile Desc"


Mapping:  ALRPatientToALRCSV
Source:   ALRPatient
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "Legacy CSV Assignment List Report"
Id:       alr-csv-patient
* -> "Patient"
* identifier -> "BENE_MBI_ID"
* name.given -> "BENE_1ST_NAME"
* name.family -> "BENE_LAST_NAME"
* gender -> "BENE_SEX_CD"
* birthDate -> "BENE_BRTH_DT"
* deceasedDateTime -> "BENE_DEATH_DT"
* address.district -> "GEO_SSA_CNTY_CD_NAME"
* address.state -> "GEO_SSA_STATE_NAME"
