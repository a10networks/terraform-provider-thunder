---
layout: "thunder"
page_title: "thunder: thunder_vrrp_vrid"
sidebar_current: "docs-thunder-resource-vrrp-vrid"
description: |-
	Provides details about thunder vrrp vrid resource for A10
---

# thunder\_vrrp\_vrid

`thunder_vrrp_vrid` Provides details about thunder vrrp vrid
## Example Usage


```hcl
provider "thunder" {
    address  = "${var.address}"
    username = "${var.username}"  
    password = "${var.password}"
}

resource "thunder_vrrp_vrid" "Vrrp_Vrid_Test" {
        pair_follow {  
        pair_follow =  0 
        vrid_lead =  "string" 
        }
blade_parameters {  
        priority =  0 
        fail_over_policy_template =  "string" 
        uuid =  "string" 
tracking_options {  
 vlan-cfg {   
        vlan =  0 
        timeout =  0 
        priority_cost =  0 
        }
        uuid =  "string" 
route {  
 ipv6-destination-cfg {   
        distance =  0 
        protocol =  "string" 
        priority_cost =  0 
        ipv6_destination =  "string" 
        gatewayv6 =  "string" 
        }
ip-destination-cfg {   
        distance =  0 
        protocol =  "string" 
        mask =  "string" 
        priority_cost =  0 
        ip_destination =  "string" 
        gateway =  "string" 
        }
        }
bgp {  
 bgp-ipv4-address-cfg {   
        bgp_ipv4_address =  "string" 
        priority_cost =  0 
        }
bgp-ipv6-address-cfg {   
        bgp_ipv6_address =  "string" 
        priority_cost =  0 
        }
        }
interface {   
        ethernet =  0 
        priority_cost =  0 
        }
gateway {  
 ipv4-gateway-list {   
        uuid =  "string" 
        ip_address =  "string" 
        priority_cost =  0 
        }
ipv6-gateway-list {   
        ipv6_address =  "string" 
        uuid =  "string" 
        priority_cost =  0 
        }
        }
trunk-cfg {   
        priority_cost =  0 
        trunk =  0 
        per_port_pri =  0 
        }
        }
        }
uuid = "string"
vrid_val = 0
user_tag = "string"
preempt_mode {  
        threshold =  0 
        disable =  0 
        }
sampling-enable {   
        counters1 =  "string" 
        }
floating_ip {  
 ipv6-address-part-cfg {   
        ethernet =  0 
        ipv6_address_partition =  "string" 
        ve =  0 
        trunk =  0 
        }
ip-address-cfg {   
        ip_address =  "string" 
        }
ip-address-part-cfg {   
        ip_address_partition =  "string" 
        }
ipv6-address-cfg {   
        ipv6_address =  "string" 
        ethernet =  0 
        ve =  0 
        trunk =  0 
        }
        }
follow {  
        vrid_lead =  "string" 
        }
 
}
```

## Argument Reference

* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `vrid_val` - Specify ha VRRP-A vrid
* `fail_over_policy_template` - Apply a fail over policy template (VRRP-A fail over policy template name)
* `priority` - VRRP-A priorty (Priority, default is 150)
* `priority_cost` - The amount the priority will decrease
* `vlan` - VLAN tracking (VLAN id)
* `distance` - Route’s administrative distance (default: match any)
* `gatewayv6` - Match the route’s gateway (next-hop) (default: match any)
* `ipv6_destination` - IPv6 Destination Prefix
* `protocol` - ‘any’: Match any routing protocol (default); ‘static’: Match only static routes (added by user); ‘dynamic’: Match routes added by dynamic routing protocols (e.g. OSPF);
* `gateway` - Match the route’s gateway (next-hop) (default: match any)
* `ip_destination` - Destination prefix
* `mask` - Destination prefix mask
* `bgp_ipv4_address` - bgp IP Address
* `bgp_ipv6_address` - IPV6 address
* `ethernet` - Ethernet (for link-local address only)
* `ip_address` - IP Address
* `ipv6_address` - IPV6 address
* `per_port_pri` - per port priority
* `trunk` - Trunk interface (for link-local address only)
* `ipv6_address_partition` - IPV6 address
* `ve` - VE interface (for link-local address only)
* `ip_address_partition` - IP Address
* `disable` - disable preemption
* `threshold` - preemption threshold (preemption threshhold (0-255), default 0)
* `vrid_lead` - Define a VRRP-A VRID leader
