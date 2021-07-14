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
* extension[ext-covidFlag][0].extension[monthNum].valueInteger = 1
* extension[ext-covidFlag][0].extension[flag].valueInteger = 0
* extension[ext-covidFlag][1].extension[monthNum].valueInteger = 2
* extension[ext-covidFlag][1].extension[flag].valueInteger = 0
* extension[ext-covidFlag][2].extension[monthNum].valueInteger = 3
* extension[ext-covidFlag][2].extension[flag].valueInteger = 0
* extension[ext-covidFlag][3].extension[monthNum].valueInteger = 4
* extension[ext-covidFlag][3].extension[flag].valueInteger = 0
* extension[ext-covidFlag][4].extension[monthNum].valueInteger = 5
* extension[ext-covidFlag][4].extension[flag].valueInteger = 1
* extension[ext-covidFlag][5].extension[monthNum].valueInteger = 6
* extension[ext-covidFlag][5].extension[flag].valueInteger = 1

