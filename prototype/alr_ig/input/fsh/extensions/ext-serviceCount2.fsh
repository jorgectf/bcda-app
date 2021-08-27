
Extension:      ServiceCount2
Id:             ext-serviceCount2
Title:          "ALR Service Count 2"
Description:    "ALR Service Count Description 2"

/*
* value[x] only string
*/

* extension contains
    participant 0..* and
    serviceCount 0..1
* extension[participant].value[x] only Identifier
* extension[serviceCount].value[x] only Quantity
