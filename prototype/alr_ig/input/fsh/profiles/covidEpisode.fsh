// Your definitions here!
Profile:        ALRCovidEpisode
Parent:         EpisodeOfCare
Id:             alr-covidEpisode
Title:          "ALR Covid Episode"
Description:    "ALR Covid Episode Desc"

// not allowed 
// period 1..*
* statusHistory 0..0
* type 0..0
* referralRequest 0..0
* managingOrganization 0..0
* careManager 0..0
* team 0..0
* account 0..0
* extension contains CovidEpisode named ext-covidEpisode 0..1
* extension contains CovidFlag named ext-covidFlag 0..* MS



/*
Mapping:  ALRRiskFlagToALRCSV
Source:   ALRHccRiskFlag
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report"
Id:       alr-csv-group
* -> "ALR Table 1-1"
* component -> "HCC_COL_N"
*/
