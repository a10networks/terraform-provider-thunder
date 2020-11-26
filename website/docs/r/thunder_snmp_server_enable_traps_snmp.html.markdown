---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_snmp"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-snmp"
description: |-
	Provides details about thunder snmp server enable traps snmp resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_snmp

`thunder_snmp_server_enable_traps_snmp` Provides details about thunder snmp server enable traps snmp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_snmp_server_enable_traps_snmp" "resourceSnmpServerEnableTrapsSnmpTest" {
	linkup = 0
all = 0
linkdown = 0
uuid = "string"
 
}

```

## Argument Reference

* `all` - Enable all SNMP group traps
* `linkdown` - Enable SNMP link-down trap
* `linkup` - Enable SNMP link-up trap
* `uuid` - uuid of the object

