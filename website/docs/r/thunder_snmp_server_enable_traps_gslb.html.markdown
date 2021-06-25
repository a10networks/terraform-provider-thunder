---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_gslb"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-gslb"
description: |-
	Provides details about thunder snmp server enable traps gslb resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_gslb

`thunder_snmp_server_enable_traps_gslb` Provides details about thunder snmp server enable traps gslb
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_snmp_server_enable_traps_gslb" "Snmp_Server_Enable_Traps_Gslb_Test" {
all = 0
group = 0
uuid = "string"
zone = 0
site = 0
service_ip = 0
 
}

```

## Argument Reference

* `all` - Enable all GSLB traps
* `group` - Enable GSLB group related traps
* `service_ip` - Enable GSLB service-ip related traps
* `site` - Enable GSLB site related traps
* `uuid` - uuid of the object
* `zone` - Enable GSLB zone related traps

