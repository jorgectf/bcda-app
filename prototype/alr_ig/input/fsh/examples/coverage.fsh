Instance:       ExampleCoverage
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-Coverage
Title:          "Coverage: General Example"
Usage:          #example
* id = "example-id-coverage"
* status = #active
* subscriberId = "-19999000381"
* beneficiary.reference = "Patient/example-id-patient"
* period.start = "2020-01-01"
* period.end = "2020-12-31"
* payor.reference = "Organization/123"
* extension[ext-enrollmentPeriod].extension[period][0].valuePeriod.start = "2020-01-01"
* extension[ext-enrollmentPeriod].extension[period][0].valuePeriod.end = "2020-06-30"
* extension[ext-enrollmentPeriod].extension[period][1].valuePeriod.start = "2020-10-01"
* extension[ext-enrollmentPeriod].extension[period][1].valuePeriod.end = "2020-12-31"


