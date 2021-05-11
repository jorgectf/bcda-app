// Your definitions here!
Profile:        ALRHccRiskFlag
Parent:         Observation
Id:             alr-HccRiskFlag
Title:          "ALR HCC Flag Profile"
Description:    "ALR HCC Flag Profile Desc"


Mapping:  ALRRiskFlagToALRCSV
Source:   ALRHccRiskFlag
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report"
Id:       alr-csv-group
* -> "Observation"
* component -> "HCC_COL_N"
