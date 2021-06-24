---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_vrrp_a"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-vrrp-a"
description: |-
	Provides details about thunder snmp server enable traps vrrp a resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_vrrp\_a

`thunder_snmp_server_enable_traps_vrrp_a` Provides details about thunder snmp server enable traps vrrp a
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}



resource "thunder_snmp_server_enable_traps_vrrp_a" "resourceSnmpServerEnableTrapsVrrpATest" {
	active = 0
standby = 0
all = 0
uuid = "string"
 
}

```

## Argument Reference

* `active` - Enable VRRP-A active trap
* `all` - Enable all VRRP-A group traps
* `standby` - Enable VRRP-A standby trap
* `uuid` - uuid of the object

