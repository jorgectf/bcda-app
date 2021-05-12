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
* -> "ALR Table 1-1"
* prediction -> "BENE_RSK_R_SCRE_N"
* code -> "BENE_RSK_R_SCRE_N"
* prediction -> "ESRD_SCORE"
* code -> "ESRD_SCORE"
* prediction -> "DIS_SCORE"
* code -> "DIS_SCORE"
* prediction -> "AGDU_SCORE"
* code -> "AGDU_SCORE"
* prediction -> "AGND_SCORE"
* code -> "AGND_SCORE"
* prediction -> "DEM_ESRD_SCORE"
* code -> "DEM_ESRD_SCORE"
* prediction -> "DEM_DIS_SCORE"
* code -> "DEM_DIS_SCORE"
* prediction -> "DEM_AGDU_SCORE"
* code -> "DEM_AGDU_SCORE"
* prediction -> "DEM_AGND_SCORE"
* code -> "DEM_AGND_SCORE"
