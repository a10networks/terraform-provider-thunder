---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_location"
sidebar_current: "docs-thunder-resource-snmp-server-location"
description: |-
	Provides details about thunder snmp server location resource for A10
---

# thunder\_snmp\_server\_location

`thunder_snmp_server_location` Provides details about thunder snmp server location
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_location" "resourceSnmpServerLocationTest" {
	loc = "string"
uuid = "string"
 
}

```

## Argument Reference

* `loc` - The physical location of this node
* `uuid` - uuid of the object

