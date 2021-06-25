---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_management_index"
sidebar_current: "docs-thunder-resource-snmp-server-management-index"
description: |-
	Provides details about thunder snmp server management index resource for A10
---

# thunder\_snmp\_server\_management\_index

`thunder_snmp_server_management_index` Provides details about thunder snmp server management index
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_snmp_server_management_index" "resourceSnmpServerManagementIndexTest" {
	mgmt_index = 0
uuid = "string"
 
}

```

## Argument Reference

* `mgmt_index` - index number
* `uuid` - uuid of the object

