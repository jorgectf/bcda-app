
Extension:      ServiceCountCCN
Id:             ext-serviceCountCCN
Title:          "Service Count by ACO Participant CCN level"
Description:    "Assigned Beneficiaries and Number of Primary Care Services at the ACO Participant CCN Level"

/*
* value[x] only string
*/

* extension contains
    participant 0..1 and
    serviceCount 0..1
* extension[participant].value[x] only Identifier
* extension[serviceCount].value[x] only integer

