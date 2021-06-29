# Background

This overview provides information on the FHIR implementation of the Assignment List Report (ALR), produced by Center for Medicare and Medicaid Services (CMS).  The ALR  Implementation Guide is based on work done by the DaVinci project.  Work on the ALR project is done in collaboration and alignment with two specific DaVinci projects:

- [DaVinci Member Attribution List](https://confluence.hl7.org/pages/viewpage.action?pageId=65077395)
- [DaVinci Risk Adjustment](https://confluence.hl7.org/display/DVP/Risk+Adjustment)

Da Vinci is an HL7-sponsored project that brings together the U.S. payer ,providers, and technology suppliers (including EHR vendors) to help payers and providers to positively impact clinical, quality, cost, and care management outcomes using FHIR-related technologies. 
Additional information about Da Vinci, its members, the use cases and the implementation guides being developed and other materials can be found on the [Da Vinci Confluence page](https://confluence.hl7.org/display/DVP/Da+Vinci+Welcome).


## Underlying Technologies

ALR is part of the [Beneficiary Claims Data API (BCDA)](https://bcda.cms.gov/) family of products and subsequently is based on the same technologies, conventions, operations and mechanism used by BCDA, including [HL7 FHIR](http://www.hl7.org/fhir/index.html) and [Bulk Data IG](https://hl7.org/fhir/uv/bulkdata/index.html).  Implementers of this specification therefore need to understand basic information about BCDA along with the aforementioned specifications.

## FHIR Resources

This implementation guide uses terminology, notations and design principles that are specific to FHIR. Before reading this implementation guide, it’s important to be familiar with some of the basic principles of FHIR as well as general guidance on how to read FHIR specifications. Readers who are unfamiliar with FHIR are encouraged to read (or at least skim) the following prior to reading the rest of this implementation guide.

- FHIR overview
- Developer’s introduction (or Clinical introduction)
- FHIR data types
- Using codes
- References between resources
- How to read resource & profile definitions
- Base resource


## Use Cases and Overview

Summary, business context, goals, background, etc.

## Data Model

{{ image link here}}

## Mapping ALR to FHIR

The current Assignment List Report is delivered in seven separate CSV files.  The ALR data dictionary refers to each file as a "table", e.g. Table 1-1, Table 1-2…Table 1-7. To facilitate understanding of how the ALR data is mapped to FHIR, the data elements in each table of the ALR report have been grouped and categorized as shown in the following table, along with high level FHIR mapping information

|Table (File)|Description|Data Element Categories|FHIR Resource|
---|---|---|---|
|Table 1-1|Assignment Data|Demographics<br>Assignment Flags<br>Exclusion Flags<br>Enrollment Flags<br>Risk Flags<br>Risk Scores|Group, Patient<br>Group (Extension)<br>Group<br>Coverage<br>Observation<br>RiskAssessment|
|Table 1-2|Service Count by ACO TIN|Demographics<br>Service Count|Patient<br>Patient (Extension)|
|Table 1-3|Service Count by ACO NPI|Demographics<br>Service Count|Patient<br>Patient (Extension)|
|Table 1-4|Top N by ACO TIN+NPI|Demographics<br>Service Count|Patient<br>Patient (Extension)|
|Table 1-5|Turnover Analysis|Demographics<br>Turnover Flags|Group, Patient<br>Patient (Extension)|
|Table 1-6|Assignable Beneficiaries|Demographics<br>Assignment Flag|Group, Patient<br>Group (Extension)|
|Table 1-7|COVID Episode Information|Episode of Care|EpisodeOfCare
			

The Group resource is the focal point of the payload.  The Group.member list element contains a reference to each assigned patient, along with various assignment flags that apply to that group member.  The Patient resource contains the demographics found in all ALR tables, along with a few extensions to handle service counts (tables 1-2…1-4).  The HCC risk flags and scores are handled  using a combination of the Observation resource and the RiskAssessment resource.

The attached spreadsheet contains details on categories and mappings for each ALR data element

{{ LINK HERE }}

