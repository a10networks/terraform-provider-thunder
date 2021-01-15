---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_address_family_ipv6_redistribute"
sidebar_current: "docs-thunder-resource-router-bgp-address-family-ipv6-redistribute"
description: |-
    Provides details about thunder router bgp address family ipv6 redistribute resource for A10
---

# thunder\_router\_bgp\_address\_family\_ipv6\_redistribute

`thunder_router_bgp_address_family_ipv6_redistribute` Provides details about thunder router bgp address family ipv6 redistribute
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_address_family_ipv6_redistribute" "resourceRouterBgpAddressFamilyIpv6RedistributeTest" {
	connected_cfg {  
 	connected =  0 
	route_map =  "string" 
	}
floating_ip_cfg {  
 	floating_ip =  0 
	route_map =  "string" 
	}
nat64_cfg {  
 	nat64 =  0 
	route_map =  "string" 
	}
nat_map_cfg {  
 	nat_map =  0 
	route_map =  "string" 
	}
lw4o6_cfg {  
 	lw4o6 =  0 
	route_map =  "string" 
	}
static_nat_cfg {  
 	static_nat =  0 
	route_map =  "string" 
	}
ip_nat_cfg {  
 	ip_nat =  0 
	route_map =  "string" 
	}
ip_nat_list_cfg {  
 	ip_nat_list =  0 
	route_map =  "string" 
	}
isis_cfg {  
 	isis =  0 
	route_map =  "string" 
	}
ospf_cfg {  
 	ospf =  0 
	route_map =  "string" 
	}
rip_cfg {  
 	rip =  0 
	route_map =  "string" 
	}
static_cfg {  
 	static =  0 
	route_map =  "string" 
	}
vip {  
 only_flagged_cfg {  
 	only_flagged =  0 
	route_map =  "string" 
	}
only_not_flagged_cfg {  
 	only_not_flagged =  0 
	route_map =  "string" 
	}
	}
uuid = "string"
 
}

```

## Argument Reference

* `redistribute` - Redistribute information from another routing protocol
* `connected` - Connected
* `route-map` - Route map reference (Pointer to route-map entries)
* `floating-ip` - Floating IP
* `nat64` - NAT64 Prefix
* `nat-map` - NAT MAP Prefix
* `lw4o6` - LW4O6 Prefix
* `static-nat` - Static NAT Prefix
* `ip-nat` - IP NAT
* `ip-nat-list` - IP NAT list
* `isis` - ISO IS-IS
* `ospf` - Open Shortest Path First (OSPF)
* `rip` - Routing Information Protocol (RIP)
* `static` - Static routes
* `only-flagged` - Selected Virtual IP (VIP)
* `only-not-flagged` - Only not flagged
* `uuid` - uuid of the object

