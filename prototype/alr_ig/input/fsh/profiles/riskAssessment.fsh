// Your definitions here!
Profile:        ALRRiskAssessment
Parent:         RiskAssessment
Id:             alr-RiskAssessment
Title:          "ALR Risk Assessment Profile"
Description:    "ALR Risk Assessment Profile Desc"

// not allowed 
* basedOn 0..0
* parent 0..0
* encounter 0..0
* reasonCode 0..0
* reasonReference 0..0
* prediction.qualitativeRisk 0..0
* prediction.relativeRisk 0..0
* mitigation 0..0
// * extension contains HccRiskFlag named ext-hccRiskFlag 0..* MS


Mapping:  ALRRiskAssessmentToALRCSV
Source:   ALRRiskAssessment
Target:   "http://cms.gov/ALR/Datadictionary"
Title:    "CMS Assignment List Report"
Id:       alr-csv-riskAssessment
* -> "RiskAssessment"
* prediction -> "BENE_RSK_R_SCRE_N"
* code -> "BENE_RSK_R_SCRE_N"
