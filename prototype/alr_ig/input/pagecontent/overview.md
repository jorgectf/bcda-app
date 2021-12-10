### Background

This page provides information on the FHIR implementation of the Assignment List Report (ALR), produced by Center for Medicare and Medicaid Services (CMS).  The ALR  Implementation Guide is based on work done by the DaVinci project.  Work on the ALR project is done in collaboration and alignment with the following DaVinci projects:

- [DaVinci Member Attribution List](https://confluence.hl7.org/pages/viewpage.action?pageId=65077395)
- [DaVinci Risk Adjustment](https://confluence.hl7.org/display/DVP/Risk+Adjustment)

Da Vinci is an HL7-sponsored project that brings together the U.S. payer ,providers, and technology suppliers (including EHR vendors) to help payers and providers to positively impact clinical, quality, cost, and care management outcomes using FHIR-related technologies. 
Additional information about Da Vinci, its members, the use cases and the implementation guides being developed and other materials can be found on the [Da Vinci Confluence page](https://confluence.hl7.org/display/DVP/Da+Vinci+Welcome).


### Overview and Use Cases

As part of participating in CMS’ alternative payment models, Accountable Care Organizations (ACOs) are assigned beneficiaries. CMS provides ACOs with information on their assigned population and financial performance at the start of the agreement period and routinely during the performance year. The Assignment List Report (ALRs) is the mechanism that provides this information to assist ACOs with healthcare operations and reporting / analytics. 



### Underlying Technologies

ALR is part of the [Beneficiary Claims Data API (BCDA)](https://bcda.cms.gov/) family of products and subsequently is based on the same technologies, conventions, operations and mechanisms used by BCDA, including [the HL7 FHIR specification](http://www.hl7.org/fhir/index.html) and the [Bulk FHIR specification](https://hl7.org/fhir/uv/bulkdata/index.html).  Implementers of this ALR specification therefore need to understand basic information about BCDA along with the aforementioned specifications.

### FHIR Resources

This implementation guide uses terminology, notations and design principles that are specific to FHIR. Before reading this implementation guide, it’s important to be familiar with some of the basic principles of FHIR as well as general guidance on how to read FHIR specifications. Readers who are unfamiliar with FHIR are encouraged to read (or at least skim) the following prior to reading the rest of this implementation guide.


- [FHIR overview](http://hl7.org/fhir/R4/overview.html)
- [Developer’s introduction](http://hl7.org/fhir/R4/overview-dev.html)
- [FHIR data types](http://hl7.org/fhir/R4/datatypes.html)
- [Using codes](http://hl7.org/fhir/R4/terminologies.html)
- [References between resources](http://hl7.org/fhir/R4/references.html)
- [How to read resource & profile definitions](http://hl7.org/fhir/us/davinci-atr/background.html)
- [Base resource](http://hl7.org/fhir/R4/resource.html)




### Understanding the Data Model

The Group resource serves as the container for assigned beneficiaries and is the focal point of the payload.  The member list element in the Group resource contains a reference to each assigned patient, along with various assignment and exclusion flags that apply to that member.  The Patient resource contains the demographics for each beneficiary, along with a few extensions to handle things such as aggregate service counts.  Turnover beneficiaries (e.g. beneficiaries who were assigned in a previous quarter but are no longer assigned) are not members of the Group, but are included in the Patient extract.  The HCC risk flags and scores are handled using a combination of the Observation resource and the RiskAssessment resource.  HCC Risk Flags for a given patient are delivered in a single Observation resource, with all flags contained in the Observation.component[] list.  Enrollment flags are supplied in the Coverage resource, and COVID episode information is populated in the EpisodeOfCare resource.  


![data model diagram](diagram.png)



### Mapping ALR to FHIR

This section describes how fields in the current CSV based ALR report are mapped to FHIR elements.  The current Assignment List Report is delivered in seven separate CSV files.  The ALR data dictionary refers to each file as a "table", e.g. Table 1-1, Table 1-2…Table 1-7. To facilitate understanding of how the ALR data is mapped to FHIR, the data elements in each table of the ALR report have been grouped and categorized as shown in the following table.  Each data element category also shows the associated FHIR resource.


<table border="1">
<tr>
    <th>Table (File)</th>
    <th>Description</th>
    <th>Data Element Category</th>
    <th>FHIR Resource</th>
    <th>Notes</th>
</tr>
<tr>
    <td>Table 1-1</td>
    <td>Assignment Data</td>
    <td>
        <ul>
        <li>Demographics</li>
        <li>Assignment Flags</li>
        <li>Exclusion Flags</li>
        <li>Enrollment Flags</li>
        <li>Risk Flags</li>
        <li>Risk Scores</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Group, Patient</li>
        <li>Group (Extensions)</li>
        <li>Group (Extensions)</li>
        <li>Coverage (Extensions)</li>
        <li>Observation</li>
        <li>Risk Asessment</li>
        </ul>
    </td>
    <td>
        Assigned patients from ALR table 1-1 appear in FHIR as members of the Group resource (e.g. Group.member).  Each Group.member is a reference to the corresponding Patient resource along with some relevant extensions to clarify their membership.  Patients in ALR table 1-1 can have a status of assigned but excluded (see EXCLUDED in the ALR data dictionary, table 1-1).  In such cases, the changeType extension will indicate "excluded" and the changeReason extension will indicate the reason for exclusion (e.g. deceased, missing id, ghp).  Similarly, patients who are no longer assigned (table 1-5, turnover) are also members of the group, but have a changeType of "dropped" and the changeReason extension will indicate the reason (e.g. plurality, AB, HMO).  See Table 1-5 in the ALR data dictionary.
    </td>
</tr>
<tr>
    <td>Table 1-2</td>
    <td>Service Count by ACO TIN</td>
    <td>
        <ul>
        <li>Demographics</li>
        <li>Service Count</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Patient</li>
        <li>Patient (Extensions)</li>
        </ul>
    </td>
    <td>
        This count appears as an extension in the corresponding patient resource
    </td>
</tr>
<tr>
    <td>Table 1-3</td>
    <td>Service Count by ACO NPI</td>
    <td>
        <ul>
        <li>Demographics</li>
        <li>Service Count</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Patient</li>
        <li>Patient (Extensions)</li>
        </ul>
    </td>
    <td>
        This count appears as an extension in the corresponding patient resource
    </td>
</tr>
<tr>
    <td>Table 1-4</td>
    <td>Top N by ACO TIN+NPI</td>
    <td>
        <ul>
        <li>Demographics</li>
        <li>Service Count</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Patient</li>
        <li>Patient (Extensions)</li>
        </ul>
    </td>
    <td>
        This count appears as an extension in the corresponding patient resource
    </td>
</tr>
<tr>
    <td>Table 1-5</td>
    <td>Turnover Analysis</td>
    <td>
        <ul>
        <li>Demographics</li>
        <li>Turnover Flags</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Patient</li>
        <li>Patient (Extensions)</li>
        </ul>
    </td>
    <td>
        Patients no longer assigned (ALR table 1-5) will appear in FHIR as members of the Group, similar to Table 1-1, but with a change type of "dropped".  The changeReason extension provides the reason they are no longer assigned.
    </td>
</tr>
<tr>
    <td>Table 1-6</td>
    <td>Assignable Beneficiaries</td>
    <td>
        <ul>
        <li>Demographics</li>
        <li>Assignment Flag</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Patient</li>
        </ul>
    </td>
    <td>
        Assignable (but not assigned) beneficiaries appear as Patient resources in the payload but are not members of the Group resource.
    </td>
</tr>
<tr>
    <td>Table 1-7</td>
    <td>COVID Episode Information</td>
    <td>
        <ul>
        <li>Episode of Care</li>
        </ul>
    </td>
    <td>
        <ul>
        <li>Episode of Care</li>
        </ul>
    </td>
</tr>

</table>





