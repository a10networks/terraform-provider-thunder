---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_disable_traps"
sidebar_current: "docs-thunder-resource-snmp-server-disable-traps"
description: |-
	Provides details about thunder snmp server disable traps resource for A10
---

# thunder\_snmp\_server\_disable\_traps

`thunder_snmp_server_disable_traps` Provides details about thunder snmp server disable traps
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_disable_traps" "Snmp_Server_Disable_Traps_Test" {
        all = 0
slb_change = 0
uuid = "string"
vrrp_a = 0
snmp = 0
gslb = 0
slb = 0
 
}

```

## Argument Reference

* `all` - Disable all traps on this partition
* `gslb` - Disable all gslb traps on this partition
* `slb` - Disable all slb traps on this partition
* `slb_change` - Disable all slb-change traps on this partition
* `snmp` - Disable all snmp traps on this partition
* `uuid` - uuid of the object
* `vrrp_a` - Disable all vrrp-a on this partition

