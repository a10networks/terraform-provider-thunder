---
layout: "thunder"
page_title: "thunder: thunder_vrrp_a_vrid"
sidebar_current: "docs-thunder-resource-vrrp-a-vrid"
description: |-
    Provides details about thunder vrrp a vrid resource for A10
---

# thunder\_vrrp\_a\_vrid

`thunder_vrrp_a_vrid` Provides details about thunder vrrp a vrid
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_vrrp_a_vrid" "resourceVrrpAVridTest" {
	vrid_val = 0
floating_ip {  
 ip-address-cfg {   
	ip_address =  "string" 
	}
ip-address-part-cfg {   
	ip_address_partition =  "string" 
	}
ipv6-address-cfg {   
	ipv6_address =  "string" 
	ethernet =  0 
	trunk =  0 
	ve =  0 
	}
ipv6-address-part-cfg {   
	ipv6_address_partition =  "string" 
	ethernet =  0 
	trunk =  0 
	ve =  0 
	}
	}
preempt_mode {  
 	threshold =  0 
	disable =  0 
	}
follow {  
 	vrid_lead =  "string" 
	}
pair_follow {  
 	pair_follow =  0 
	vrid_lead =  "string" 
	}
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
blade_parameters {  
 	priority =  0 
	fail_over_policy_template =  "string" 
	uuid =  "string" 
tracking_options {  
 interface {   
	ethernet =  0 
	priority_cost =  0 
	}
route {  
 ip-destination-cfg {   
	ip_destination =  "string" 
	mask =  "string" 
	priority_cost =  0 
	gateway =  "string" 
	distance =  0 
	protocol =  "string" 
	}
ipv6-destination-cfg {   
	ipv6_destination =  "string" 
	priority_cost =  0 
	gatewayv6 =  "string" 
	distance =  0 
	protocol =  "string" 
	}
	}
trunk-cfg {   
	trunk =  0 
	priority_cost =  0 
	per_port_pri =  0 
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
vlan-cfg {   
	vlan =  0 
	timeout =  0 
	priority_cost =  0 
	}
	uuid =  "string" 
gateway {  
 ipv4-gateway-list {   
	ip_address =  "string" 
	priority_cost =  0 
	uuid =  "string" 
	}
ipv6-gateway-list {   
	ipv6_address =  "string" 
	priority_cost =  0 
	uuid =  "string" 
	}
	}
	}
	}
 
}

```

## Argument Reference

* `vrid` - Specify VRRP-A vrid
* `vrid-val` - Specify ha VRRP-A vrid
* `ip-address` - IP Address
* `ip-address-partition` - IP Address
* `ipv6-address` - IPV6 address
* `ethernet` - Ethernet Interface (Ethernet interface number)
* `trunk` - trunk tracking (Trunk Number)
* `ve` - VE interface (for link-local address only)
* `ipv6-address-partition` - IPV6 address
* `threshold` - preemption threshold (preemption threshhold (0-255), default 0)
* `disable` - disable preemption
* `vrid-lead` - Define a VRRP-A VRID leader
* `pair-follow` - Follow other VRRP-A vrid
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'associated_vip_count': Number of vips associated to vrid; 'associated_vport_count': Number of vports associated to vrid; 'associated_natpool_count': Number of nat pools associated to vrid;
* `priority` - VRRP-A priorty (Priority, default is 150)
* `fail-over-policy-template` - Apply a fail over policy template (VRRP-A fail over policy template name)
* `priority-cost` - The amount the priority will decrease
* `ip-destination` - Destination prefix
* `mask` - Destination prefix mask
* `gateway` - Match the route's gateway (next-hop) (default: match any)
* `distance` - Route's administrative distance (default: match any)
* `protocol` - 'any': Match any routing protocol (default); 'static': Match only static routes (added by user); 'dynamic': Match routes added by dynamic routing protocols (e.g. OSPF);
* `ipv6-destination` - IPv6 Destination Prefix
* `gatewayv6` - Match the route's gateway (next-hop) (default: match any)
* `per-port-pri` - per port priority
* `bgp-ipv4-address` - bgp IP Address
* `bgp-ipv6-address` - IPV6 address
* `vlan` - VLAN tracking (VLAN id)

