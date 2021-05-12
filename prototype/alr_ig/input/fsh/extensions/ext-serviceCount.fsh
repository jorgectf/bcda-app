
Extension:      ServiceCount
Id:             ext-serviceCount
Title:          "ALR Service Count"
Description:    "ALR Service Count Description"

/*
* value[x] only string
*/

* extension contains
    category 0..1 and
    participant 0..* and
    serviceCount 0..1
* extension[category].value[x] only Coding
* extension[participant].value[x] only Identifier
* extension[serviceCount].value[x] only integer

/*
* extension contains
    participant 0..* 
* extension[participant].value[x] only Identifier
*/

/*
* extension contains
    period 0..*
* extension[period].value[x] only Period

*/