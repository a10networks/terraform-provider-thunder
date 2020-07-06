---
layout: "vthunder"
page_title: "vthunder: vthunder_interface_ethernet_lldp"
sidebar_current: "docs-vthunder-resource-interface-ethernet-lldp"
description: |-
  Provides details about vthunder interface ethernet lldp resource for A10
---

# vthunder\_interface\_ethernet\_lldp

`vthunder_interface_ethernet_lldp` Provides details about vthunder interface ethernet lldp
## Example Usage


```hcl
provider "vthunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "vthunder_interface_ethernet_lldp" "ethernetlldp" {
  ifnum=1
  enable_cfg {
    rt_enable = 1
    rx = 1
    tx = 1
  }
}
```

## Argument Reference

* `ifnum` - Interface no.
* `uuid` - uuid of the object
* `link_aggregation` - Interface link aggregation information
* `tx_dot1_tlvs` - Interface lldp tx IEEE 802.1 Organizationally specific TLVs configuration
* `vlan` - Interface vlan information
* `notif_enable` - Interface lldp notification enable
* `notification` - Interface lldp notification configuration
* `exclude` - Configure which TLVs excluded. All basic TLVs will be included by default
* `management_address` - Interface lldp management address
* `port_description` - Interface lldp port description
* `system_capabilities` - Interface lldp system capabilities
* `system_description` - Interface lldp system description
* `system_name` - Interface lldp system name
* `tx_tlvs` - Interface lldp tx TLVs configuration
* `rt_enable` - Interface lldp enable/disable
* `rx` - Enable lldp rx
* `tx` - Enable lldp tx

