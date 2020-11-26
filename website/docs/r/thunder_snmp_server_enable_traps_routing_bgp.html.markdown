---
layout: "thunder"
page_title: "thunder: thunder_snmp_server_enable_traps_routing_bgp"
sidebar_current: "docs-thunder-resource-snmp-server-enable-traps-routing-bgp"
description: |-
	Provides details about thunder snmp server enable traps routing bgp resource for A10
---

# thunder\_snmp\_server\_enable\_traps\_routing\_bgp

`thunder_snmp_server_enable_traps_routing_bgp` Provides details about thunder snmp server enable traps routing bgp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_snmp_server_enable_traps_routing_bgp" "Snmp_Server_Enable_Traps_Routing_Bgp_Test" {

bgp_established_notification = 0
uuid = "string"
bgp_backward_trans_notification = 0
 
}
```

## Argument Reference

* `bgp_backward_trans_notification` - Enable bgpBackwardTransNotification traps
* `bgp_established_notification` - Enable bgpEstablishedNotification traps
* `uuid` - uuid of the object

