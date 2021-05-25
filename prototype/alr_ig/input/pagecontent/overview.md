# Overview of the FHIR ALR Implementation

## Mappings

The current Assignment List Report has many elements, spanning across several files.  THe elements can be categorized
as follows

Category | Associated ALR Table(s) | | FHIR Resource
------ | ------|----------
Demographics      | Table 1-1     | | Patient
Assignment Flags       | Table 1-1     | | Group 
Exclusion Flags      | Table 1-1     | |  Group
HCC Risk Flags      | Table 1-1     | |  Observation
HCC Risk Scores      | Table 1-1     | |  RiskAssessment
Aggregates/Counts | Table 1-2, 1-3, 1-4 | |  Patient 
Covid Episode Info | Table 1-7 |  | EpisodeOfCare

To view details on element categories and FHIR mappings, see this spreadsheet.

![FHIR resources and relationships](diagram.png "FHIR resources and relationships")

