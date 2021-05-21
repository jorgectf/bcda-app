
Extension:      ServiceCountTIN
Id:             ext-serviceCountTIN
Title:          "Service Count by ACO Participant TIN level"
Description:    "Assigned Beneficiaries and Number of Primary Care Services at the ACO Participant TIN Level "

/*
* value[x] only string
*/

* extension contains
    participant 0..1 and
    serviceCount 0..1
* extension[participant].value[x] only Identifier
* extension[serviceCount].value[x] only integer

