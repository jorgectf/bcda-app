
Extension:      ServiceCountTINNPI
Id:             ext-serviceCountTIN-NPI
Title:          "Top ACO Participant TIN-Individual NPI"
Description:    "Top ACO Participant TIN-Individual NPI (by number of primary care services) for Assigned Beneficiaries"

/*
* value[x] only string
*/

* extension contains
    participant 0..* and
    serviceCount 0..1
* extension[participant].value[x] only Identifier
* extension[serviceCount].value[x] only integer

