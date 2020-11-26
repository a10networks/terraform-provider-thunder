---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_network"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-network"
description: |-
	Provides details about thunder snmp server enable traps network resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_network

`thunder_snmp_server_enable_traps_network` Provides details about thunder snmp server enable traps network
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_enable_traps_network" "Snmp_Server_Enable_Traps_Network_Test" {
        trunk_port_threshold = 0
uuid = "string"
 
}
```

## Argument Reference

* `trunk_port_threshold` - Enable network trunk-port-threshold trap
* `uuid` - uuid of the object

