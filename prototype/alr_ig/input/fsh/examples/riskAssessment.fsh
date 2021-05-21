Instance:       ExampleMonthlyHccRiskScore
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-RiskAssessment
Title:          "Risk Score: CMS-HCC Monthly Risk Score"
Usage:          #example
* id = "example-id-monthly-risk-score"
* status = #final
* text.div = "<div xmlns=\"http://www.w3.org/1999/xhtml\"><p>General example of HCC Risk Scores, using the Risk Assessment resource</p></div>"
* text.status = #additional
* code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/bene_rsk_r_scre"
* code.coding[0].code = #BENE_RSK_R_SCRE
* code.coding[0].display = "CMS-HCC Monthly Risk Scores"
* basis[0].reference = "Observation/example-id-hcc-risk-flags"
* subject.reference = "Patient/example-id-patient"
* prediction[0].probabilityDecimal = 0.083132
* prediction[0].whenPeriod.start = "2020-01-01"
* prediction[0].whenPeriod.end = "2020-01-31"
* prediction[1].probabilityDecimal = 1.061543
* prediction[1].whenPeriod.start = "2020-02-01"
* prediction[1].whenPeriod.end = "2020-02-28"
* prediction[2].probabilityDecimal = 0.077255
* prediction[2].whenPeriod.start = "2020-03-01"
* prediction[2].whenPeriod.end = "2020-03-31"
* prediction[3].probabilityDecimal = 0.077255
* prediction[3].whenPeriod.start = "2020-04-01"
* prediction[3].whenPeriod.end = "2020-04-30"
* prediction[4].probabilityDecimal = 0.077255
* prediction[4].whenPeriod.start = "2020-05-01"
* prediction[4].whenPeriod.end = "2020-05-31"
* prediction[5].probabilityDecimal = 0.077255
* prediction[5].whenPeriod.start = "2020-06-01"
* prediction[5].whenPeriod.end = "2020-06-30"
* prediction[6].probabilityDecimal = 0.067255
* prediction[6].whenPeriod.start = "2020-07-01"
* prediction[6].whenPeriod.end = "2020-07-31"
* prediction[7].probabilityDecimal = 0.037255
* prediction[7].whenPeriod.start = "2020-08-01"
* prediction[7].whenPeriod.end = "2020-08-31"
* prediction[8].probabilityDecimal = 0.0555255
* prediction[8].whenPeriod.start = "2020-09-01"
* prediction[8].whenPeriod.end = "2020-09-30"
* prediction[9].probabilityDecimal = 0.067255
* prediction[9].whenPeriod.start = "2020-10-01"
* prediction[9].whenPeriod.end = "2020-10-31"
* prediction[10].probabilityDecimal = 0.074455
* prediction[10].whenPeriod.start = "2020-11-01"
* prediction[10].whenPeriod.end = "2020-11-30"
* prediction[11].probabilityDecimal = 0.0733255
* prediction[11].whenPeriod.start = "2020-12-01"
* prediction[11].whenPeriod.end = "2020-12-31"


Instance:       ExampleHccRiskScoreESRD
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-RiskAssessment
Title:          "Risk Score: CMS-HCC Risk Score for ESRD"
Usage:          #example
* id = "example-id-risk-score-esrd"
* status = #final
* code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/esrd-score"
* code.coding[0].code = #ESRD_SCORE
* code.coding[0].display = "CMS-HCC Risk Score for ESRD"
* basis[0].reference = "Observation/example-id-hcc-risk-flags"
* subject.reference = "Patient/example-id-patient"
* prediction[0].probabilityDecimal = 0.083132
* prediction[0].whenPeriod.start = "2020-01-01"
* prediction[0].whenPeriod.end = "2020-12-31"

Instance:       ExampleHccRiskScoreDisabled
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-RiskAssessment
Title:          "Risk Score: CMS-HCC Risk Score for Disabled"
Usage:          #example
* id = "example-id-risk-score-disabled"
* status = #final
* code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/disabled-score"
* code.coding[0].code = #DISABLED_SCORE
* code.coding[0].display = "CMS-HCC Risk Score for disabled"
* basis[0].reference = "Observation/example-id-hcc-risk-flags"
* subject.reference = "Patient/example-id-patient"
* prediction[0].probabilityDecimal = 0.083132
* prediction[0].whenPeriod.start = "2020-01-01"
* prediction[0].whenPeriod.end = "2020-12-31"

Instance:       ExampleHccRiskScoreAgedDual
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-RiskAssessment
Title:          "Risk Score: CMS-HCC Risk Score for Aged/Dual"
Usage:          #example
* id = "example-id-risk-score-aged-dual"
* status = #final
* code.coding[0].system = "https://bluebutton.cms.gov/resources/variables/alr/aged-dual-score"
* code.coding[0].code = #AGED_DUAL_SCORE
* code.coding[0].display = "CMS-HCC Risk Score for Aged/Dual"
* basis[0].reference = "Observation/example-id-hcc-risk-flags"
* subject.reference = "Patient/example-id-patient"
* prediction[0].probabilityDecimal = 0.083132
* prediction[0].whenPeriod.start = "2020-01-01"
* prediction[0].whenPeriod.end = "2020-12-31"
