Instance:       ExampleCoverage
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-Coverage
Title:          "Coverage: General Example"
Usage:          #example
* id = "example-id-coverage"
* status = #active
* text.div = "<div xmlns=\"http://www.w3.org/1999/xhtml\"><p>General example of a Coverage resource, with an extension for Enrollment periods (ALR enrollment flags)</p></div>"
* text.status = #additional
* subscriberId = "-19999000381"
* beneficiary.reference = "Patient/example-id-patient"
* period.start = "2020-01-01"
* period.end = "2020-12-31"
* payor.reference = "Organization/123"
* extension[ext-enrollmentPeriod].extension[period][0].valuePeriod.start = "2020-01-01"
* extension[ext-enrollmentPeriod].extension[period][0].valuePeriod.end = "2020-06-30"
* extension[ext-enrollmentPeriod].extension[period][1].valuePeriod.start = "2020-10-01"
* extension[ext-enrollmentPeriod].extension[period][1].valuePeriod.end = "2020-12-31"


