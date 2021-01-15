---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor"
sidebar_current: "docs-thunder-resource-router-bgp-address-family-ipv6-neighbor-ipv6-neighbor"
description: |-
    Provides details about thunder router bgp address family ipv6 neighbor ipv6 neighbor resource for A10
---

# thunder\_router\_bgp\_address\_family\_ipv6\_neighbor\_ipv6\_neighbor

`thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor` Provides details about thunder router bgp address family ipv6 neighbor ipv6 neighbor
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_address_family_ipv6_neighbor_ipv6_neighbor" "resourceRouterBgpAddressFamilyIpv6NeighborIpv6NeighborTest" {
	neighbor_ipv6 = "string"
peer_group_name = "string"
activate = 0
allowas_in = 0
allowas_in_count = 0
prefix_list_direction = "string"
graceful_restart = 0
default_originate = 0
route_map = "string"
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
maximum_prefix = 0
maximum_prefix_thres = 0
next_hop_self = 0
neighbor-prefix-lists {   
	nbr_prefix_list =  "string" 
	nbr_prefix_list_direction =  "string" 
	}
remove_private_as = 0
neighbor-route-map-lists {   
	nbr_route_map =  "string" 
	nbr_rmap_direction =  "string" 
	}
send_community_val = "string"
inbound = 0
unsuppress_map = "string"
weight = 0
uuid = "string"
 
}

```

## Argument Reference

* `ipv6-neighbor` - Specify a peer-group neighbor router
* `neighbor-ipv6` - Neighbor IPv6 address
* `peer-group-name` - Configure peer-group (peer-group name)
* `activate` - Enable the Address Family for this Neighbor
* `allowas-in` - Accept as-path with my AS present in it
* `allowas-in-count` - Number of occurrences of AS number
* `prefix-list-direction` - 'both': both; 'receive': receive; 'send': send;
* `graceful-restart` - enable graceful-restart helper for this neighbor
* `default-originate` - Originate default route to this neighbor
* `route-map` - Route-map to specify criteria to originate default (route-map name)
* `distribute-list` - Filter updates to/from this neighbor (IP standard/extended/named access list)
* `distribute-list-direction` - 'in': in; 'out': out;
* `filter-list` - Establish BGP filters (AS path access-list name)
* `filter-list-direction` - 'in': in; 'out': out;
* `maximum-prefix` - Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
* `maximum-prefix-thres` - threshold-value, 1 to 100 percent
* `next-hop-self` - Disable the next hop calculation for this neighbor
* `nbr-prefix-list` - Filter updates to/from this neighbor (Name of a prefix list)
* `nbr-prefix-list-direction` - 'in': in; 'out': out;
* `remove-private-as` - Remove private AS number from outbound updates
* `nbr-route-map` - Apply route map to neighbor (Name of route map)
* `nbr-rmap-direction` - 'in': in; 'out': out;
* `send-community-val` - 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;
* `inbound` - Allow inbound soft reconfiguration for this neighbor
* `unsuppress-map` - Route-map to selectively unsuppress suppressed routes (Name of route map)
* `weight` - Set default weight for routes from this neighbor
* `uuid` - uuid of the object

