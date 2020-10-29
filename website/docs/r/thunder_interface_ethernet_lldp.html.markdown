---
layout: "thunder"
page_title: "thunder: thunder_interface_ethernet_lldp"
sidebar_current: "docs-thunder-resource-interface-ethernet-lldp"
description: |-
	Provides details about thunder interface ethernet lldp resource for A10
---

# thunder\_interface\_ethernet\_lldp

`thunder_interface_ethernet_lldp` Provides details about thunder interface ethernet lldp
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_interface_ethernet_lldp" "Interface_Ethernet_Lldp_Test" {
        tx_dot1_cfg {  
        link_aggregation =  0 
        vlan =  0 
        tx_dot1_tlvs =  0 
        }
notification_cfg {  
        notification =  0 
        notif_enable =  0 
        }
enable_cfg {  
        rx =  0 
        tx =  0 
        rt_enable =  0 
        }
tx_tlvs_cfg {  
        system_capabilities =  0 
        system_description =  0 
        management_address =  0 
        tx_tlvs =  0 
        exclude =  0 
        port_description =  0 
        system_name =  0 
        }
uuid = "string"
 
}
```

## Argument Reference

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
