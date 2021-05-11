Instance:       ExamplePatient
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-Patient
Title:          "Patient: General Example"
Usage:          #example
* id = "example-id-patient"
* text.status = #additional
* text.div = "<div xmlns=\"http://www.w3.org/1999/xhtml\"><p>General example of coverage with gaps in enrollment period</p></div>"
* identifier[0].system = "https://bluebutton.cms.gov/resources/variables/bene_id"
* identifier[0].value = "example-ccw-id"
* identifier[1].system = "https://bluebutton.cms.gov/resources/identifier/mbi-hash"
* identifier[1].value = "412be343c343343324f661527671abb5307d5e89031db9a9f"
* identifier[2].system = "http://hl7.org/fhir/sid/us-mbi"
* identifier[2].value = "example-mbi-patient"
// TODO: there is a BB extension that goes with us-mbi that dictates whether mbi is current or historical
// but I"m basing this Patient off of CARIN, not BB so....
* name.given[0] = "John"
* name.family = "Doe"
* gender = #male
* birthDate = "1957-01-17"
* deceasedDateTime = "2020-03-18"
* address.line[0] = "1000 Main Street"
* address.line[1] = "Apt #2"
* address.city = "Anytown"
* address.state = "UT"
* address.district = "SL County"

