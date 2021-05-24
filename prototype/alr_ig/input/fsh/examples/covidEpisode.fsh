Instance:       ExampleCovidEpisode
InstanceOf:     http://alr.cms.gov/ig/StructureDefinition/alr-covidEpisode
Title:          "Covid Episode: General Example"
Usage:          #example
* id = "example-id-episode"
* status = #active
* text.div = "<div xmlns=\"http://www.w3.org/1999/xhtml\"><p>General example of a Covid Episode </p></div>"
* text.status = #additional

* status = #finished
* patient.reference = "patient/example-id-patient"
* period.start = "2020-05-10"
* period.end = "2020-05-15"
* diagnosis.condition.identifier.system = "http://hl7.org/fhir/sid/icd-10"
* diagnosis.condition.identifier.value = "U071"
* extension[ext-covidEpisode].valueInteger = 1
* extension[ext-covidFlag][0].extension[period].valuePeriod.start = "2020-01-01"
* extension[ext-covidFlag][0].extension[period].valuePeriod.end = "2020-01-31"
* extension[ext-covidFlag][0].extension[flag].valueInteger = 0
* extension[ext-covidFlag][1].extension[period].valuePeriod.start = "2020-02-01"
* extension[ext-covidFlag][1].extension[period].valuePeriod.end = "2020-02-28"
* extension[ext-covidFlag][1].extension[flag].valueInteger = 0
* extension[ext-covidFlag][2].extension[period].valuePeriod.start = "2020-03-01"
* extension[ext-covidFlag][2].extension[period].valuePeriod.end = "2020-03-31"
* extension[ext-covidFlag][2].extension[flag].valueInteger = 0
* extension[ext-covidFlag][3].extension[period].valuePeriod.start = "2020-04-01"
* extension[ext-covidFlag][3].extension[period].valuePeriod.end = "2020-04-30"
* extension[ext-covidFlag][3].extension[flag].valueInteger = 0
* extension[ext-covidFlag][4].extension[period].valuePeriod.start = "2020-05-01"
* extension[ext-covidFlag][4].extension[period].valuePeriod.end = "2020-05-31"
* extension[ext-covidFlag][4].extension[flag].valueInteger = 1
* extension[ext-covidFlag][5].extension[period].valuePeriod.start = "2020-06-01"
* extension[ext-covidFlag][5].extension[period].valuePeriod.end = "2020-06-30"
* extension[ext-covidFlag][5].extension[flag].valueInteger = 0

