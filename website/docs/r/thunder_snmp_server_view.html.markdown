---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_view"
sidebar_current: "docs-thunder-resource-snmp-server-view"
description: |-
	Provides details about thunder snmp server view resource for A10
---

# thunder\_snmp\_server\_view

`thunder_snmp_server_view` Provides details about thunder snmp server view
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_snmp_server_view" "resourceSnmpServerViewTest" {
	type = "string"
oid = "string"
mask = "string"
uuid = "string"
viewname = "string"
 
}

```

## Argument Reference

* `mask` - Hex octets, separated by ‘.’, mask of oid
* `oid` - MIB view family name or oid
* `type` - ‘included’: MIB family is included in the view; ‘excluded’: MIB family is excluded from the view;
* `uuid` - uuid of the object
* `viewname` - Name of the view

