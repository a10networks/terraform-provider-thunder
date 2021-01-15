---
layout: "thunder"
page_title: "thunder: thunder_router_bgp_neighbor_ipv4_neighbor"
sidebar_current: "docs-thunder-resource-router-bgp-neighbor-ipv4-neighbor"
description: |-
    Provides details about thunder router bgp neighbor ipv4 neighbor resource for A10
---

# thunder\_router\_bgp\_neighbor\_ipv4\_neighbor

`thunder_router_bgp_neighbor_ipv4_neighbor` Provides details about thunder router bgp neighbor ipv4 neighbor
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}


resource "thunder_router_bgp_neighbor_ipv4_neighbor" "resourceRouterBgpNeighborIpv4NeighborTest" {
	neighbor_ipv4 = "string"
nbr_remote_as = 0
peer_group_name = "string"
activate = 0
advertisement_interval = 0
allowas_in = 0
allowas_in_count = 0
as_origination_interval = 0
dynamic = 0
prefix_list_direction = "string"
route_refresh = 0
graceful_restart = 0
collide_established = 0
default_originate = 0
route_map = "string"
description = "string"
disallow_infinite_holdtime = 0
distribute-lists {   
	distribute_list =  "string" 
	distribute_list_direction =  "string" 
	}
acos_application_only = 0
dont_capability_negotiate = 0
ebgp_multihop = 0
ebgp_multihop_hop_count = 0
enforce_multihop = 0
bfd = 0
multihop = 0
key_id = 0
key_type = "string"
bfd_value = "string"
neighbor-filter-lists {   
	filter_list =  "string" 
	filter_list_direction =  "string" 
	}
maximum_prefix = 0
maximum_prefix_thres = 0
next_hop_self = 0
override_capability = 0
pass_value = "string"
passive = 0
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
shutdown = 0
strict_capability_match = 0
timers_keepalive = 0
timers_holdtime = 0
connect = 0
unsuppress_map = "string"
update_source_ip = "string"
update_source_ipv6 = "string"
ethernet = 0
loopback = 0
ve = 0
trunk = 0
lif = 0
tunnel = 0
weight = 0
uuid = "string"
 
}

```

## Argument Reference

* `ipv4-neighbor` - Specify a ipv4 neighbor router
* `neighbor-ipv4` - Neighbor address
* `nbr-remote-as` - Specify AS number of BGP neighbor
* `peer-group-name` - Configure peer-group (peer-group name)
* `activate` - Enable the Address Family for this Neighbor
* `advertisement-interval` - Minimum interval between sending BGP routing updates (time in seconds)
* `allowas-in` - Accept as-path with my AS present in it
* `allowas-in-count` - Number of occurrences of AS number
* `as-origination-interval` - Minimum interval between sending AS-origination routing updates (time in seconds)
* `dynamic` - Advertise dynamic capability to this neighbor
* `prefix-list-direction` - 'both': both; 'receive': receive; 'send': send;
* `route-refresh` - Advertise route-refresh capability to this neighbor
* `graceful-restart` - enable graceful-restart helper for this neighbor
* `collide-established` - Include Neighbor in Established State for Collision Detection
* `default-originate` - Originate default route to this neighbor
* `route-map` - Route-map to specify criteria to originate default (route-map name)
* `description` - Neighbor specific description (Up to 80 characters describing this neighbor)
* `disallow-infinite-holdtime` - BGP per neighbor disallow-infinite-holdtime
* `distribute-list` - Filter updates to/from this neighbor (IP standard/extended/named access list)
* `distribute-list-direction` - 'in': in; 'out': out;
* `acos-application-only` - Send BGP update to ACOS application
* `dont-capability-negotiate` - Do not perform capability negotiation
* `ebgp-multihop` - Allow EBGP neighbors not on directly connected networks
* `ebgp-multihop-hop-count` - maximum hop count
* `enforce-multihop` - Enforce EBGP neighbors to perform multihop
* `bfd` - Bidirectional Forwarding Detection (BFD)
* `multihop` - Enable multihop
* `key-id` - Key ID
* `key-type` - 'md5': md5; 'meticulous-md5': meticulous-md5; 'meticulous-sha1': meticulous-sha1; 'sha1': sha1; 'simple': simple;  (Keyed MD5/Meticulous Keyed MD5/Meticulous Keyed SHA1/Keyed SHA1/Simple Password)
* `bfd-value` - Key String
* `bfd-encrypted` - Do NOT use this option manually. (This is an A10 reserved keyword.) (The ENCRYPTED password string)
* `filter-list` - Establish BGP filters (AS path access-list name)
* `filter-list-direction` - 'in': in; 'out': out;
* `maximum-prefix` - Maximum number of prefix accept from this peer (maximum no. of prefix limit (various depends on model))
* `maximum-prefix-thres` - threshold-value, 1 to 100 percent
* `next-hop-self` - Disable the next hop calculation for this neighbor
* `override-capability` - Override capability negotiation result
* `pass-value` - Key String
* `passive` - Don't send open messages to this neighbor
* `nbr-prefix-list` - Filter updates to/from this neighbor (Name of a prefix list)
* `nbr-prefix-list-direction` - 'in': in; 'out': out;
* `remove-private-as` - Remove private AS number from outbound updates
* `nbr-route-map` - Apply route map to neighbor (Name of route map)
* `nbr-rmap-direction` - 'in': in; 'out': out;
* `send-community-val` - 'both': Send Standard and Extended Community attributes; 'none': Disable Sending Community attributes; 'standard': Send Standard Community attributes; 'extended': Send Extended Community attributes;
* `inbound` - Allow inbound soft reconfiguration for this neighbor
* `shutdown` - Administratively shut down this neighbor
* `strict-capability-match` - Strict capability negotiation match
* `timers-keepalive` - Keepalive interval
* `timers-holdtime` - Holdtime
* `connect` - BGP connect timer
* `unsuppress-map` - Route-map to selectively unsuppress suppressed routes (Name of route map)
* `update-source-ip` - IP address
* `update-source-ipv6` - IPv6 address
* `ethernet` - Ethernet interface (Port number)
* `loopback` - Loopback interface (Port number)
* `ve` - Virtual ethernet interface (Virtual ethernet interface number)
* `trunk` - Trunk interface (Trunk interface number)
* `lif` - Logical interface (Lif interface number)
* `tunnel` - Tunnel interface (Tunnel interface number)
* `weight` - Set default weight for routes from this neighbor
* `uuid` - uuid of the object

